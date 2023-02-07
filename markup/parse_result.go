package markup

import (
	"fmt"
	"strconv"
)

// ValueType represents what type the value of an attribute property holds.
type ValueType int8

const (
	// All the types that an attribute property can hold.
	ValueTypeInteger = ValueType(iota)
	ValueTypeFloat
	ValueTypeString
	ValueTypeBool
)

// Value holds the value of an attribute property.
type Value struct {
	IntegerValue int
	FloatValue   float64
	StringValue  string
	BoolValue    bool
	ValueType    ValueType
}

func (v *Value) toString() string {
	switch v.ValueType {
	case ValueTypeInteger:
		return strconv.Itoa(v.IntegerValue)
	case ValueTypeFloat:
		if v.FloatValue == float64(int(v.FloatValue)) {
			return strconv.Itoa(int(v.FloatValue))
		}
		return fmt.Sprint(v.FloatValue)
	case ValueTypeString:
		return v.StringValue
	case ValueTypeBool:
		if v.BoolValue {
			return "True"
		}
		return "False"
	}

	return ""
}

type property struct {
	name  string
	value Value
}

func toPropertyMap(properties []property) map[string]Value {
	result := make(map[string]Value, len(properties))
	for _, property := range properties {
		result[property.name] = property.value
	}

	return result
}

type attributeMarker struct {
	name           string
	position       int
	sourcePosition int
	properties     []property
	tagType        tagType
}

// GetProperty returns the value associated to the given property name in the given attribute marker, if there's one.
func (am *attributeMarker) GetProperty(name string) (Value, bool) {
	for _, property := range am.properties {
		if property.name == name {
			return property.value, true
		}
	}
	return Value{}, false
}

// Attribute holds data about a section of text between two markup tags.
type Attribute struct {
	Position       int
	Length         int
	Name           string
	Properties     map[string]Value
	SourcePosition int
}

// ParseResult is the result of parsing a line of text into interpolated text and a set of markup attributes.
type ParseResult struct {
	Text       string
	Attributes []Attribute
}

// Attribute returns the value of the attribute with the given name, if there is one.
func (parseResult *ParseResult) Attribute(name string) (Attribute, bool) {
	for _, attribute := range parseResult.Attributes {
		if attribute.Name == name {
			return attribute, true
		}
	}
	return Attribute{}, false
}

// TextForAttribute returns the part of the parsed line of text that is covered by the given attribute.
func (parseResult *ParseResult) TextForAttribute(attribute Attribute) string {
	if attribute.Length == 0 {
		return ""
	}
	if len(parseResult.Text) < attribute.Position+attribute.Length {
		panic("attribute represents a range not representable by this text")
	}

	return string([]rune(parseResult.Text)[attribute.Position : attribute.Position+attribute.Length])
}
