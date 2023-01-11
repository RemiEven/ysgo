package markup

import (
	"fmt"
	"strconv"
)

type ValueType int

const (
	ValueTypeInteger = ValueType(iota)
	ValueTypeFloat
	ValueTypeString
	ValueTypeBool
)

func (valueType ValueType) String() string {
	return map[ValueType]string{
		ValueTypeInteger: "integer",
		ValueTypeFloat:   "float",
		ValueTypeString:  "string",
		ValueTypeBool:    "boolean",
	}[valueType]
}

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

type Property struct {
	name  string
	value Value
}

func toPropertyMap(properties []Property) map[string]Value {
	result := make(map[string]Value, len(properties))
	for _, property := range properties {
		result[property.name] = property.value
	}

	return result
}

type AttributeMarker struct {
	name           string
	position       int
	sourcePosition int
	properties     []Property
	tagType        TagType
}

func (attributeMarker *AttributeMarker) GetProperty(name string) (Value, bool) {
	for _, property := range attributeMarker.properties {
		if property.name == name {
			return property.value, true
		}
	}
	return Value{}, false
}

type Attribute struct {
	Position       int
	Length         int
	Name           string
	Properties     map[string]Value
	SourcePosition int
}

type ParseResult struct {
	Text       string
	Attributes []Attribute
}

func (parseResult *ParseResult) GetAttributeWithName(name string) (Attribute, bool) {
	for _, attribute := range parseResult.Attributes {
		if attribute.Name == name {
			return attribute, true
		}
	}
	return Attribute{}, false
}

func (parseResult *ParseResult) TextForAttribute(attribute Attribute) string {
	if attribute.Length == 0 {
		return ""
	}
	if len(parseResult.Text) < attribute.Position+attribute.Length {
		panic("attribute represents a range not representable by this text")
	}

	return string([]rune(parseResult.Text)[attribute.Position : attribute.Position+attribute.Length])
}
