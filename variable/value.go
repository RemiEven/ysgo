package variable

import (
	"fmt"
	"strconv"
)

// Value holds an evaluated value of a YarnSpinner script.
type Value struct {
	Number  *float64 `json:"n,omitempty"`
	Boolean *bool    `json:"b,omitempty"`
	String  *string  `json:"s,omitempty"`
}

// NewNumber creates a new YarnSpinner dialogue value from the given number.
func NewNumber(number float64) *Value {
	return &Value{
		Number: &number,
	}
}

// NewBoolean creates a new YarnSpinner dialogue value from the given boolean.
func NewBoolean(boolean bool) *Value {
	return &Value{
		Boolean: &boolean,
	}
}

// NewString creates a new YarnSpinner dialogue value from the given string.
func NewString(str string) *Value {
	return &Value{
		String: &str,
	}
}

// ToString serializes a value into a string.
func (v *Value) ToString() string {
	switch {
	case v.Number != nil:
		n := *v.Number
		if n == float64(int(n)) {
			return strconv.Itoa(int(n))
		}
		return fmt.Sprint(n)
	case v.Boolean != nil:
		if *v.Boolean {
			return "True"
		}
		return "False"
	case v.String != nil:
		return *v.String
	}
	return ""
}
