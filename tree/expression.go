package tree

import (
	"fmt"
	"strconv"

	"github.com/RemiEven/ysgo/parser"
)

const (
	MultiplicationBinaryOperator = iota
	DivisionBinaryOperator
	ModuloBinaryOperator
	AdditionBinaryOperator
	SubtractionBinaryOperator
	LessThanEqualsBinaryOperator
	GreaterThanEqualsBinaryOperator
	LessBinaryOperator
	GreaterBinaryOperator
	EqualsBinaryOperator
	NotEqualsBinaryOperator
	AndBinaryOperator
	OrBinaryOperator
	XorBinaryOperator

	AssignmentInPlaceOperator = iota
	MultiplicationInPlaceOperator
	DivisionInPlaceOperator
	ModuloInPlaceOperator
	AdditionInPlaceOperator
	SubtractionInPlaceOperator
)

func tokenToBinaryOperator(token int) (*int, bool) {
	operator, ok := map[int]int{
		parser.YarnSpinnerLexerOPERATOR_LOGICAL_LESS_THAN_EQUALS:    LessThanEqualsBinaryOperator,
		parser.YarnSpinnerLexerOPERATOR_LOGICAL_GREATER_THAN_EQUALS: GreaterThanEqualsBinaryOperator,
		parser.YarnSpinnerLexerOPERATOR_LOGICAL_EQUALS:              EqualsBinaryOperator,
		parser.YarnSpinnerLexerOPERATOR_LOGICAL_LESS:                LessBinaryOperator,
		parser.YarnSpinnerLexerOPERATOR_LOGICAL_GREATER:             GreaterBinaryOperator,
		parser.YarnSpinnerLexerOPERATOR_LOGICAL_NOT_EQUALS:          NotEqualsBinaryOperator,
		parser.YarnSpinnerLexerOPERATOR_LOGICAL_AND:                 AndBinaryOperator,
		parser.YarnSpinnerLexerOPERATOR_LOGICAL_OR:                  OrBinaryOperator,
		parser.YarnSpinnerLexerOPERATOR_LOGICAL_XOR:                 XorBinaryOperator,
		parser.YarnSpinnerLexerOPERATOR_MATHS_ADDITION:              AdditionBinaryOperator,
		parser.YarnSpinnerLexerOPERATOR_MATHS_SUBTRACTION:           SubtractionBinaryOperator,
		parser.YarnSpinnerLexerOPERATOR_MATHS_MULTIPLICATION:        MultiplicationBinaryOperator,
		parser.YarnSpinnerLexerOPERATOR_MATHS_DIVISION:              DivisionBinaryOperator,
		parser.YarnSpinnerLexerOPERATOR_MATHS_MODULUS:               ModuloBinaryOperator,
	}[token]

	return &operator, ok
}

func tokenToInplaceOperator(token int) (*int, bool) {
	operator, ok := map[int]int{
		parser.YarnSpinnerLexerOPERATOR_ASSIGNMENT:                  AssignmentInPlaceOperator,
		parser.YarnSpinnerLexerOPERATOR_MATHS_MULTIPLICATION_EQUALS: MultiplicationInPlaceOperator,
		parser.YarnSpinnerLexerOPERATOR_MATHS_DIVISION_EQUALS:       DivisionInPlaceOperator,
		parser.YarnSpinnerLexerOPERATOR_MATHS_MODULUS_EQUALS:        ModuloInPlaceOperator,
		parser.YarnSpinnerLexerOPERATOR_MATHS_ADDITION_EQUALS:       AdditionInPlaceOperator,
		parser.YarnSpinnerLexerOPERATOR_MATHS_SUBTRACTION_EQUALS:    SubtractionInPlaceOperator,
	}[token]

	return &operator, ok
}

type Expression struct {
	Value                     *Value
	NegativeExpression        *Expression
	NotExpression             *Expression
	LeftOperand, RightOperand *Expression
	Operator                  *int
}

type Value struct {
	Number       *float64
	Boolean      *bool
	String       *string
	VariableID   *string
	FunctionCall *FunctionCall
}

func NewNumberValue(number float64) *Value {
	return &Value{
		Number: &number,
	}
}

func NewBooleanValue(boolean bool) *Value {
	return &Value{
		Boolean: &boolean,
	}
}

func NewStringValue(str string) *Value {
	return &Value{
		String: &str,
	}
}

func (v *Value) IsNumber() bool {
	return v.Number != nil || v.FunctionCall != nil
}

func (v *Value) IsBoolean() bool {
	return v.Boolean != nil || v.FunctionCall != nil
}

func (v *Value) IsString() bool {
	return v.String != nil || v.FunctionCall != nil
}

func (v *Value) ToString() string {
	switch {
	case v.IsNumber():
		n := *v.Number
		if n == float64(int(n)) {
			return strconv.Itoa(int(n))
		}
		return fmt.Sprintf("%f", n)
	case v.IsBoolean():
		return strconv.FormatBool(*v.Boolean)
	case v.IsString():
		return *v.String
	case v.VariableID != nil:
		return "$" + *v.VariableID
	case v.FunctionCall != nil:
		return v.FunctionCall.FunctionID + "(...)"
	}
	return "empty value"
}

type FunctionCall struct {
	FunctionID string
	Arguments  []*Expression
}
