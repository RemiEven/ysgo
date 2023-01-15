package variable

import (
	"fmt"
	"strconv"
)

type Value struct {
	Number  *float64
	Boolean *bool
	String  *string
}

func NewNumber(number float64) *Value {
	return &Value{
		Number: &number,
	}
}

func NewBoolean(boolean bool) *Value {
	return &Value{
		Boolean: &boolean,
	}
}

func NewString(str string) *Value {
	return &Value{
		String: &str,
	}
}

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
