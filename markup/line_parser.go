package markup

import (
	"fmt"
	"io"
	"math"
	"regexp"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf16"

	"golang.org/x/exp/slices"
)

const (
	replacementMarkerContents      = "contents"
	trimWhitespaceProperty         = "trimwhitespace"
	characterAttribute             = "character"
	characterAttributeNameProperty = "name"
)

var endOfCharacterMarker = regexp.MustCompile(`:\s*`)

type LineParser struct {
	input            string
	reader           *strings.Reader
	sourcePosition   int
	position         int
	markerProcessors map[string]func(*AttributeMarker) (string, error)
}

func (lineParser *LineParser) ParseMarkup(input string) (*ParseResult, error) {
	lineParser.input = input
	return lineParser.parseMarkup()
}

func (lineParser *LineParser) parseMarkup() (*ParseResult, error) {
	lineParser.reader = strings.NewReader(lineParser.input)
	builder := strings.Builder{}
	markers := []AttributeMarker{}
	var lastRune rune

	for {
		nextRune, _, err := lineParser.reader.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to read next rune while parsing markup: %w", err)
		}

		if nextRune == '\\' {
			nextRune, err := peekRune(lineParser.reader)
			if err != nil {
				return nil, fmt.Errorf("failed to peek next rune after '\\': %w", err)
			}
			if nextRune == '[' || nextRune == ']' {
				nextRune, _, err := lineParser.reader.ReadRune()
				if err != nil {
					return nil, fmt.Errorf("failed to read next rune after '\\': %w", err)
				}
				builder.WriteRune(nextRune)
				lineParser.sourcePosition++
				continue
			}
		}
		if nextRune == '[' {
			lineParser.position = len([]rune(builder.String()))

			marker, err := lineParser.parseAttributeMarker()
			if err != nil {
				return nil, fmt.Errorf("failed to parse attribute marker: %w", err)
			}
			markers = append(markers, marker)

			hadPrecedingWhitespaceOrLineStart := lineParser.position == 0 || unicode.IsSpace(lastRune)

			wasReplacementMarker := false

			if _, ok := lineParser.markerProcessors[marker.name]; ok {
				wasReplacementMarker = true

				replacementText, err := lineParser.processReplacementMarker(marker)
				if err != nil {
					return nil, fmt.Errorf("failed to process replacement marker: %w", err)
				}

				builder.WriteString(replacementText)
			}

			trimWhitespaceIfAble := false

			if hadPrecedingWhitespaceOrLineStart {
				if marker.tagType == TagTypeSelfClosing {
					trimWhitespaceIfAble = !wasReplacementMarker
				}

				if property, ok := marker.GetProperty(trimWhitespaceProperty); ok {
					if property.ValueType != ValueTypeBool {
						return nil, fmt.Errorf("failed to parse line %q: attribute %q at position %d has a %v property %q - this property is required to be a boolean value", lineParser.input, marker.name, lineParser.position, property.ValueType, trimWhitespaceProperty)
					}

					trimWhitespaceIfAble = property.BoolValue
				}
			}

			if trimWhitespaceIfAble {
				if match, err := lineParser.peekWhitespace(); err != nil {
					return nil, fmt.Errorf("failed to peek whitespace: %w", err)
				} else if match {
					if _, _, err := lineParser.reader.ReadRune(); err != nil {
						return nil, fmt.Errorf("failed to read whitespace: %w", err)
					}
					lineParser.sourcePosition++
				}
			}
		} else {
			builder.WriteRune(nextRune)
			lineParser.sourcePosition++
		}

		lastRune = nextRune
	}

	attributes, err := lineParser.buildAttributesFromMarkers(markers)
	if err != nil {
		return nil, fmt.Errorf("failed to build attributes from markers: %w", err)
	}

	characterAttributeIsPresent := false
	for _, attribute := range attributes {
		if attribute.Name == characterAttribute {
			characterAttributeIsPresent = true
			break
		}
	}

	if !characterAttributeIsPresent {
		match := endOfCharacterMarker.FindStringIndex(lineParser.input)
		if match != nil {
			characterName := lineParser.input[:match[0]]
			nameValue := Value{
				StringValue: characterName,
				ValueType:   ValueTypeString,
			}
			characterAttribute := Attribute{
				Name:           characterAttribute,
				Position:       0,
				SourcePosition: 0,
				Length:         match[1],
				Properties: map[string]Value{
					characterAttributeNameProperty: nameValue,
				},
			}

			attributes = append(attributes, characterAttribute)
		}
	}

	return &ParseResult{
		Text:       builder.String(),
		Attributes: attributes,
	}, nil
}

func (lineParser *LineParser) buildAttributesFromMarkers(markers []AttributeMarker) ([]Attribute, error) {
	unclosedMarkers := []AttributeMarker{}

	attributes := make([]Attribute, 0, len(markers))

	for _, marker := range markers {
		switch marker.tagType {
		case TagTypeOpen:
			unclosedMarkers = append(unclosedMarkers, marker)
		case TagTypeClose:
			matchedOpenMarkerIndex := -1
			for i := range unclosedMarkers {
				if unclosedMarkers[i].name == marker.name {
					matchedOpenMarkerIndex = i
					break
				}
			}

			if matchedOpenMarkerIndex == -1 {
				return nil, fmt.Errorf("unexpected close marker %q at position %d in line %q", marker.name, marker.position, lineParser.input)
			}

			matchedOpenMarker := unclosedMarkers[matchedOpenMarkerIndex]
			unclosedMarkers = slices.Delete(unclosedMarkers, matchedOpenMarkerIndex, matchedOpenMarkerIndex)

			length := marker.position - matchedOpenMarker.position
			attribute := Attribute{
				Name:           marker.name,
				Properties:     toPropertyMap(matchedOpenMarker.properties),
				Position:       matchedOpenMarker.position,
				Length:         length,
				SourcePosition: matchedOpenMarker.sourcePosition,
			}

			attributes = append(attributes, attribute)
		case TagTypeSelfClosing:
			attribute := Attribute{
				Name:           marker.name,
				Properties:     toPropertyMap(marker.properties),
				Position:       marker.position,
				Length:         0,
				SourcePosition: marker.sourcePosition,
			}
			attributes = append(attributes, attribute)
		case TagTypeCloseAll:
			for _, openMarker := range unclosedMarkers {
				length := marker.position - openMarker.position
				attribute := Attribute{
					Name:           openMarker.name,
					Properties:     toPropertyMap(openMarker.properties),
					Position:       openMarker.position,
					Length:         length,
					SourcePosition: openMarker.sourcePosition,
				}
				attributes = append(attributes, attribute)
			}
			unclosedMarkers = unclosedMarkers[:0]
		}
	}

	slices.SortStableFunc(attributes, func(a, b Attribute) bool { return a.Position < b.Position })

	return attributes, nil
}

func (lineParser *LineParser) processReplacementMarker(marker AttributeMarker) (string, error) {
	if marker.tagType != TagTypeOpen && marker.tagType != TagTypeSelfClosing {
		return "", nil
	}

	if marker.tagType == TagTypeOpen {
		markerContents, err := lineParser.parseRawTextUpToAttributeClose(marker.name)
		if err != nil {
			return "", fmt.Errorf("failed to parse text up to attribute close: %w", err)
		}

		marker.properties = append(marker.properties, Property{
			name: replacementMarkerContents,
			value: Value{
				StringValue: markerContents,
				ValueType:   ValueTypeString,
			},
		})
	}

	replacementText, err := lineParser.markerProcessors[marker.name](&marker)
	if err != nil {
		return "", fmt.Errorf("failed to process marker %q: %w", marker.name, err)
	}
	return replacementText, nil
}

func (lineParser *LineParser) parseRawTextUpToAttributeClose(markerName string) (string, error) {
	rawRemainderOfLine, err := io.ReadAll(lineParser.reader)
	if err != nil {
		return "", fmt.Errorf("failed to read remainder of line: %w", err)
	}
	remainderOfLine := string(rawRemainderOfLine)

	closeTagMarkerRegexp, err := regexp.Compile(`\[\s*\/\s*({` + markerName + `})?\s*\]`)
	if err != nil {
		return "", fmt.Errorf("failed to compile close tag marker regexp: %w", err)
	}

	matchIndices := closeTagMarkerRegexp.FindStringIndex(string(remainderOfLine))

	if matchIndices == nil {
		return "", fmt.Errorf("unterminated marker %q in line %q at position %d", markerName, lineParser.input, lineParser.position)
	}

	closeMarkerPosition := matchIndices[0]

	rawTextSubstring := remainderOfLine[:closeMarkerPosition]
	lineAfterRawText := remainderOfLine[closeMarkerPosition:]

	lineParser.reader = strings.NewReader(lineAfterRawText)
	return rawTextSubstring, nil
}

func (lineParser *LineParser) parseAttributeMarker() (AttributeMarker, error) {
	sourcePositionAtMarkerStart := lineParser.sourcePosition
	lineParser.sourcePosition++

	if match, err := lineParser.expectPeek('/'); err != nil {
		return AttributeMarker{}, err
	} else if match {
		if err := lineParser.parseRune('/'); err != nil {
			return AttributeMarker{}, err
		}

		if match, err := lineParser.expectPeek(']'); err != nil {
			return AttributeMarker{}, err
		} else if match {
			if err := lineParser.parseRune(']'); err != nil {
				return AttributeMarker{}, err
			}
			return AttributeMarker{
				position:       lineParser.position,
				sourcePosition: sourcePositionAtMarkerStart,
				tagType:        TagTypeCloseAll,
			}, nil
		} else {
			tagName, err := lineParser.parseID()
			if err != nil {
				return AttributeMarker{}, fmt.Errorf("failed to parse ID: %w", err)
			}
			if err := lineParser.parseRune(']'); err != nil {
				return AttributeMarker{}, err
			}
			return AttributeMarker{
				name:           tagName,
				position:       lineParser.position,
				sourcePosition: sourcePositionAtMarkerStart,
				tagType:        TagTypeClose,
			}, nil
		}
	}

	attributeName, err := lineParser.parseID()
	if err != nil {
		return AttributeMarker{}, fmt.Errorf("failed to parse attribute name: %w", err)
	}
	properties := []Property{}

	if match, err := lineParser.expectPeek('='); err != nil {
		return AttributeMarker{}, err
	} else if match {
		if err := lineParser.parseRune('='); err != nil {
			return AttributeMarker{}, err
		}
		value, err := lineParser.parseValue()
		if err != nil {
			return AttributeMarker{}, fmt.Errorf("failed to parse value: %w", err)
		}
		properties = append(properties, Property{attributeName, value})
	}

	for {
		if err := lineParser.consumeWhitespace(); err != nil {
			return AttributeMarker{}, fmt.Errorf("failed to consume whitespace: %w", err)
		}

		nextRune, err := peekRune(lineParser.reader)
		if err == io.EOF {
			return AttributeMarker{}, fmt.Errorf("unexpected end of line inside markup in line %q", lineParser.input)
		} else if err != nil {
			return AttributeMarker{}, err
		}

		if nextRune == ']' {
			if err := lineParser.parseRune(']'); err != nil {
				return AttributeMarker{}, err
			}
			return AttributeMarker{
				name:           attributeName,
				position:       lineParser.position,
				sourcePosition: sourcePositionAtMarkerStart,
				properties:     properties,
				tagType:        TagTypeOpen,
			}, nil
		}

		if nextRune == '/' {
			if err := lineParser.parseRune('/'); err != nil {
				return AttributeMarker{}, err
			}
			if err := lineParser.parseRune(']'); err != nil {
				return AttributeMarker{}, err
			}
			return AttributeMarker{
				name:           attributeName,
				position:       lineParser.position,
				sourcePosition: sourcePositionAtMarkerStart,
				properties:     properties,
				tagType:        TagTypeSelfClosing,
			}, nil
		}

		propertyName, err := lineParser.parseID()
		if err != nil {
			return AttributeMarker{}, fmt.Errorf("failed to parse property name: %w", err)
		}
		if err := lineParser.parseRune('='); err != nil {
			return AttributeMarker{}, err
		}
		propertyValue, err := lineParser.parseValue()
		if err != nil {
			return AttributeMarker{}, fmt.Errorf("failed to parse property value: %w", err)
		}
		properties = append(properties, Property{propertyName, propertyValue})
	}
}

func (lineParser *LineParser) parseRune(r rune) error {
	if err := lineParser.consumeWhitespace(); err != nil {
		return err
	}

	nextRune, _, err := lineParser.reader.ReadRune()
	if err == io.EOF {
		return fmt.Errorf("unexpected end of line inside markup in line %q", lineParser.input)
	}
	if nextRune != r {
		return fmt.Errorf("expected a %v inside markup in line %q", r, lineParser.input)
	}

	lineParser.sourcePosition++
	return nil
}

func (lineParser *LineParser) parseID() (string, error) {
	if err := lineParser.consumeWhitespace(); err != nil {
		return "", fmt.Errorf("failed to consume whitespace before parsing ID: %w", err)
	}

	builder := strings.Builder{}

	nextRune, _, err := lineParser.reader.ReadRune()
	if err == io.EOF {
		return "", fmt.Errorf("unexpected end of line inside markup in line %q", lineParser.input)
	} else if err != nil {
		return "", fmt.Errorf("failed to read first rune of ID: %w", err)
	}
	lineParser.sourcePosition++

	e := string([]rune{nextRune})
	_ = e

	if utf16.IsSurrogate(nextRune) {
		nextNextRune, _, err := lineParser.reader.ReadRune()
		if err == io.EOF {
			return "", fmt.Errorf("unexpected end of line inside markup in line %q", lineParser.input)
		} else if err != nil {
			return "", fmt.Errorf("failed to read next rune after surrogate: %w", err)
		}
		lineParser.sourcePosition++

		builder.WriteRune(nextRune)
		builder.WriteRune(nextNextRune)
	} else if unicode.IsLetter(nextRune) || unicode.IsDigit(nextRune) || nextRune == '_' {
		builder.WriteRune(nextRune)
	} else {
		return "", fmt.Errorf("expected an identifier inside markup in line %q", lineParser.input)
	}

	for {
		nextRune, err := peekRune(lineParser.reader)
		if err == io.EOF {
			break
		} else if err != nil {
			return "", fmt.Errorf("failed to peek next character in ID: %w", err)
		}

		if utf16.IsSurrogate(nextRune) {
			if _, _, err := lineParser.reader.ReadRune(); err != nil {
				return "", fmt.Errorf("failed to read next character in ID: %w", err)
			}
			lineParser.sourcePosition++

			nextNextRune, _, err := lineParser.reader.ReadRune()
			if err == io.EOF {
				return "", fmt.Errorf("unexpected end of line inside markup in line %q", lineParser.input)
			} else if err != nil {
				return "", fmt.Errorf("failed to read next character after surrogate: %w", err)
			}
			lineParser.sourcePosition++
			builder.WriteRune(nextRune)
			builder.WriteRune(nextNextRune)
		} else if unicode.IsLetter(nextRune) || unicode.IsDigit(nextRune) || nextRune == '_' {
			builder.WriteRune(nextRune)
			lineParser.sourcePosition++
			if _, _, err := lineParser.reader.ReadRune(); err != nil {
				return "", fmt.Errorf("failed to read next character in ID: %w", err)
			}
		} else {
			break
		}
	}

	return builder.String(), nil
}

func (lineParser *LineParser) parseValue() (Value, error) {
	if match, err := lineParser.peekNumeric(); err != nil {
		return Value{}, err
	} else if match {
		i, err := lineParser.parseInteger()
		if err != nil {
			return Value{}, fmt.Errorf("failed to parse integer: %w", err)
		}
		if match, err := lineParser.expectPeek('.'); err != nil {
			return Value{}, err
		} else if match {
			if err := lineParser.parseRune('.'); err != nil {
				return Value{}, err
			}

			fraction, err := lineParser.parseInteger()
			if err != nil {
				return Value{}, fmt.Errorf("failed to parse fraction: %w", err)
			}
			f := float64(i) + float64(fraction)*float64(math.Pow10(-len(strconv.Itoa(fraction))))

			return Value{FloatValue: f, ValueType: ValueTypeFloat}, nil
		} else {
			return Value{IntegerValue: i, ValueType: ValueTypeInteger}, nil
		}
	}

	if match, err := lineParser.expectPeek('"'); err != nil {
		return Value{}, err
	} else if match {
		str, err := lineParser.parseString()
		if err != nil {
			return Value{}, fmt.Errorf("failed to parse string: %w", err)
		}
		return Value{StringValue: str, ValueType: ValueTypeString}, nil
	}

	word, err := lineParser.parseID()
	if err != nil {
		return Value{}, fmt.Errorf("failed to parse word: %w", err)
	}
	switch strings.ToLower(word) {
	case "true":
		return Value{BoolValue: true, ValueType: ValueTypeBool}, nil
	case "false":
		return Value{BoolValue: false, ValueType: ValueTypeBool}, nil
	default:
		return Value{StringValue: word, ValueType: ValueTypeString}, nil
	}
}

func (lineParser *LineParser) parseInteger() (int, error) {
	if err := lineParser.consumeWhitespace(); err != nil {
		return 0, fmt.Errorf("failed to consume whitespace: %w", err)
	}

	builder := strings.Builder{}

	for {
		nextRune, err := peekRune(lineParser.reader)
		if err == io.EOF {
			return 0, fmt.Errorf("unexpected end of line inside markup in line %q", lineParser.input)
		} else if err != nil {
			return 0, fmt.Errorf("failed to peek next rune: %w", err)
		}
		if unicode.IsDigit(nextRune) {
			if _, _, err := lineParser.reader.ReadRune(); err != nil {
				return 0, fmt.Errorf("failed to read next rune: %w", err)
			}
			lineParser.sourcePosition++
			builder.WriteRune(nextRune)
		} else {
			i, err := strconv.Atoi(builder.String())
			if err != nil {
				return 0, fmt.Errorf("failed to parse integer: %w", err)
			}
			return i, nil
		}
	}
}

func (lineParser *LineParser) parseString() (string, error) {
	if err := lineParser.consumeWhitespace(); err != nil {
		return "", fmt.Errorf("failed to consume whitespace: %w", err)
	}

	nextRune, _, err := lineParser.reader.ReadRune()
	if err == io.EOF {
		return "", fmt.Errorf("unexpected end of line inside markup in line %q", lineParser.input)
	} else if err != nil {
		return "", fmt.Errorf("failed to read first rune of string: %w", err)
	}
	if nextRune != '"' {
		return "", fmt.Errorf("expected a string inside markup in line %q", lineParser.input)
	}
	lineParser.sourcePosition++

	builder := strings.Builder{}

	for {
		nextRune, _, err := lineParser.reader.ReadRune()
		if err == io.EOF {
			return "", fmt.Errorf("unexpected end of line inside markup in line %q", lineParser.input)
		} else if err != nil {
			return "", fmt.Errorf("failed to read rune of string: %w", err)
		}
		lineParser.sourcePosition++

		if nextRune == '"' {
			break
		}
		if nextRune == '\\' {
			nextNextRune, _, err := lineParser.reader.ReadRune()
			if err == io.EOF {
				return "", fmt.Errorf("unexpected end of line inside markup in line %q", lineParser.input)
			} else if err != nil {
				return "", fmt.Errorf("failed to read escaped rune of string: %w", err)
			}

			lineParser.sourcePosition++
			builder.WriteRune(nextNextRune)
		} else {
			builder.WriteRune(nextRune)
		}
	}

	return builder.String(), nil
}

func (lineParser *LineParser) consumeWhitespace() error {
	allowEndOfLine := false
	for {
		nextRune, err := peekRune(lineParser.reader)
		if err == io.EOF && !allowEndOfLine {
			return fmt.Errorf("unexpected end of line inside markup in line %q", lineParser.input)
		}
		if err != io.EOF && err != nil {
			return fmt.Errorf("failed to peek rune to determine if it is whitespace: %w", err)
		}

		if !unicode.IsSpace(nextRune) {
			return nil
		}

		if _, _, err := lineParser.reader.ReadRune(); err != nil {
			return fmt.Errorf("failed to read whitespace rune: %w", err)
		}
		lineParser.sourcePosition++
	}
}

func (lineParser *LineParser) expectPeek(expected rune) (bool, error) {
	if err := lineParser.consumeWhitespace(); err != nil {
		return false, fmt.Errorf("failed to consume whitespace before peeking expected rune: %w", err)
	}
	nextRune, err := peekRune(lineParser.reader)
	if err == io.EOF {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("failed to peek rune: %w", err)
	}

	return nextRune == expected, nil
}

func (lineParser *LineParser) peekNumeric() (bool, error) {
	if err := lineParser.consumeWhitespace(); err != nil {
		return false, fmt.Errorf("failed to consume whitespace: %w", err)
	}

	nextRune, err := peekRune(lineParser.reader)
	switch {
	case err == io.EOF:
		return false, nil
	case err != nil:
		return false, err
	default:
		return unicode.IsDigit(nextRune), nil
	}
}

func (lineParser *LineParser) peekWhitespace() (bool, error) {
	nextRune, err := peekRune(lineParser.reader)
	if err == io.EOF {
		return false, nil
	} else if err != nil {
		return false, fmt.Errorf("failed to peek next rune: %w", err)
	}

	return unicode.IsSpace(nextRune), nil
}

func peekRune(reader io.RuneScanner) (rune, error) {
	nextRune, _, err := reader.ReadRune()
	switch {
	case err == io.EOF:
		return 0, nil
	case err != nil:
		return 0, fmt.Errorf("failed to read rune: %w", err)
	}
	if err := reader.UnreadRune(); err != nil {
		return 0, fmt.Errorf("failed to unread rune: %w", err)
	}
	return nextRune, nil
}
