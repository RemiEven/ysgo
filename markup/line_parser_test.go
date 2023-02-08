package markup

import (
	"testing"

	"github.com/remieven/ysgo/internal/testutils"
)

type lineParserTest struct {
	input          string
	expectedErr    error
	expectedResult *ParseResult
}

func TestLineParser(t *testing.T) {
	tests := map[string]lineParserTest{
		"simple attribute": {
			input: `Oh, [wave]hello[/wave] there!`,
			expectedResult: &ParseResult{
				Text: `Oh, hello there!`,
				Attributes: []Attribute{
					{
						Name:           "wave",
						Position:       4,
						SourcePosition: 4,
						Length:         5,
						Properties:     map[string]Value{},
					},
				},
			},
		},
		"overlapping attributes": {
			input: `Oh, [wave]hello [bounce]there![/bounce][/wave]`,
			expectedResult: &ParseResult{
				Text: `Oh, hello there!`,
				Attributes: []Attribute{
					{
						Name:           "wave",
						Position:       4,
						SourcePosition: 4,
						Length:         12,
						Properties:     map[string]Value{},
					},
					{
						Name:           "bounce",
						Position:       10,
						SourcePosition: 16,
						Length:         6,
						Properties:     map[string]Value{},
					},
				},
			},
		},
		"close-all marker": {
			input: `[wave][bounce]Hello![/]`,
			expectedResult: &ParseResult{
				Text: `Hello!`,
				Attributes: []Attribute{
					{
						Name:           "wave",
						Position:       0,
						SourcePosition: 0,
						Length:         6,
						Properties:     map[string]Value{},
					},
					{
						Name:           "bounce",
						Position:       0,
						SourcePosition: 6,
						Length:         6,
						Properties:     map[string]Value{},
					},
				},
			},
		},
		"property": {
			input: `[wave size=2]Wavy![/wave]`,
			expectedResult: &ParseResult{
				Text: `Wavy!`,
				Attributes: []Attribute{
					{
						Name:           "wave",
						Position:       0,
						SourcePosition: 0,
						Length:         5,
						Properties: map[string]Value{
							"size": {IntegerValue: 2, ValueType: ValueTypeInteger},
						},
					},
				},
			},
		},
		"short-hand property": {
			input: `[wave=2]Wavy![/wave]`,
			expectedResult: &ParseResult{
				Text: `Wavy!`,
				Attributes: []Attribute{
					{
						Name:           "wave",
						Position:       0,
						SourcePosition: 0,
						Length:         5,
						Properties: map[string]Value{
							"wave": {IntegerValue: 2, ValueType: ValueTypeInteger},
						},
					},
				},
			},
		},
		"float property": {
			input: `[wave size=2.4]Wavy![/wave]`,
			expectedResult: &ParseResult{
				Text: `Wavy!`,
				Attributes: []Attribute{
					{
						Name:           "wave",
						Position:       0,
						SourcePosition: 0,
						Length:         5,
						Properties: map[string]Value{
							"size": {FloatValue: 2.4, ValueType: ValueTypeFloat},
						},
					},
				},
			},
		},
		"boolean true property": {
			input: `[wave small=true]Wavy![/wave]`,
			expectedResult: &ParseResult{
				Text: `Wavy!`,
				Attributes: []Attribute{
					{
						Name:           "wave",
						Position:       0,
						SourcePosition: 0,
						Length:         5,
						Properties: map[string]Value{
							"small": {BoolValue: true, ValueType: ValueTypeBool},
						},
					},
				},
			},
		},
		"boolean false property": {
			input: `[wave small=false]Wavy![/wave]`,
			expectedResult: &ParseResult{
				Text: `Wavy!`,
				Attributes: []Attribute{
					{
						Name:           "wave",
						Position:       0,
						SourcePosition: 0,
						Length:         5,
						Properties: map[string]Value{
							"small": {BoolValue: false, ValueType: ValueTypeBool},
						},
					},
				},
			},
		},
		"unquoted string property": {
			input: `[mood=angry]Grr![/mood]`,
			expectedResult: &ParseResult{
				Text: `Grr!`,
				Attributes: []Attribute{
					{
						Name:           "mood",
						Position:       0,
						SourcePosition: 0,
						Length:         4,
						Properties: map[string]Value{
							"mood": {StringValue: "angry", ValueType: ValueTypeString},
						},
					},
				},
			},
		},
		"quoted string property": {
			input: `[mood="angry"]Grr![/mood]`,
			expectedResult: &ParseResult{
				Text: `Grr!`,
				Attributes: []Attribute{
					{
						Name:           "mood",
						Position:       0,
						SourcePosition: 0,
						Length:         4,
						Properties: map[string]Value{
							"mood": {StringValue: "angry", ValueType: ValueTypeString},
						},
					},
				},
			},
		},
		"trim whitespace after self-closing, nominal case": {
			input: `A [wave/] B`,
			expectedResult: &ParseResult{
				Text: `A B`,
				Attributes: []Attribute{
					{
						Name:           "wave",
						Position:       2,
						SourcePosition: 2,
						Length:         0,
						Properties:     map[string]Value{},
					},
				},
			},
		},
		"trim whitespace after self-closing, beginning of line": {
			input: `[wave/] A B`,
			expectedResult: &ParseResult{
				Text: `A B`,
				Attributes: []Attribute{
					{
						Name:           "wave",
						Position:       0,
						SourcePosition: 0,
						Length:         0,
						Properties:     map[string]Value{},
					},
				},
			},
		},
		"trimwhitespace property set to false": {
			input: `A [wave trimwhitespace=false/] B`,
			expectedResult: &ParseResult{
				Text: `A  B`,
				Attributes: []Attribute{
					{
						Name:           "wave",
						Position:       2,
						SourcePosition: 2,
						Length:         0,
						Properties: map[string]Value{
							"trimwhitespace": {BoolValue: false, ValueType: ValueTypeBool},
						},
					},
				},
			},
		},
		"text extraction": {
			input: `A [b]B [c]C[/c][/b]`,
			expectedResult: &ParseResult{
				Text: `A B C`,
				Attributes: []Attribute{
					{
						Name:           "b",
						Position:       2,
						SourcePosition: 2,
						Length:         3,
						Properties:     map[string]Value{},
					},
					{
						Name:           "c",
						Position:       4,
						SourcePosition: 7,
						Length:         1,
						Properties:     map[string]Value{},
					},
				},
			},
		},
		`multibyte character parsing "á [á]S[/á]"`: multibyteCharacterParsing(`á [á]S[/á]`, `á S`, `á`),
		`multibyte character parsing "á [a]á[/a]"`: multibyteCharacterParsing(`á [a]á[/a]`, `á á`, `a`),
		`multibyte character parsing "á [a]S[/a]"`: multibyteCharacterParsing(`á [a]S[/a]`, `á S`, `a`),
		`multibyte character parsing "S [á]S[/á]"`: multibyteCharacterParsing(`S [á]S[/á]`, `S S`, `á`),
		`multibyte character parsing "S [a]á[/a]"`: multibyteCharacterParsing(`S [a]á[/a]`, `S á`, `a`),
		`multibyte character parsing "S [a]S[/a]"`: multibyteCharacterParsing(`S [a]S[/a]`, `S S`, `a`),
		"markup escaping": {
			input: `A [b]B [c]C[/c][/b]`,
			expectedResult: &ParseResult{
				Text: `A B C`,
				Attributes: []Attribute{
					{
						Name:           "b",
						Position:       2,
						SourcePosition: 2,
						Length:         3,
						Properties:     map[string]Value{},
					},
					{
						Name:           "c",
						Position:       4,
						SourcePosition: 7,
						Length:         1,
						Properties:     map[string]Value{},
					},
				},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			lineParser := LineParser{input: test.input}
			parseResult, err := lineParser.parseMarkup()
			if !testutils.ErrorEqual(err, test.expectedErr) {
				t.Errorf("unexpected error: wanted %v, got %v", test.expectedErr, err)
				return
			}
			if diff := testutils.DeepEqual(parseResult, test.expectedResult); diff != "" {
				t.Errorf("unexpected parse result: " + diff)
				return
			}
		})
	}
}

func multibyteCharacterParsing(input, expectedText, expectedAttributeName string) lineParserTest {
	return lineParserTest{
		input: input,
		expectedResult: &ParseResult{
			Text: expectedText,
			Attributes: []Attribute{
				{
					Name:           expectedAttributeName,
					Position:       2,
					SourcePosition: 2,
					Length:         1,
					Properties:     map[string]Value{},
				},
			},
		},
	}
}
