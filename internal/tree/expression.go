package tree

import (
	"github.com/remieven/ysgo/internal/parser"
	"github.com/remieven/ysgo/variable"
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

// Expression represents something that can be evaluated to a YarnSpinner value when a Dialogue is run.
type Expression struct {
	Value                     *variable.Value
	VariableID                *string
	FunctionCall              *FunctionCall
	NegativeExpression        *Expression
	NotExpression             *Expression
	LeftOperand, RightOperand *Expression
	Operator                  *int
}

// ComplexityScore returns a score representing how complex an expression is.
// This value can be useful for saliency strategies.
func (e *Expression) ComplexityScore() int {
	return 1 + e.booleanOperatorCount()
}

// booleanOperatorCount returns the total number of boolean operations - ands, ors, nots, and
// xors - present in an expression and its sub-expressions.
func (e *Expression) booleanOperatorCount() int {
	switch {
	case e.FunctionCall != nil:
		count := 0
		for _, argument := range e.FunctionCall.Arguments {
			count += argument.booleanOperatorCount()
		}
		return count
	case e.NegativeExpression != nil:
		return e.NegativeExpression.booleanOperatorCount()
	case e.NotExpression != nil:
		return 1 + e.NotExpression.booleanOperatorCount()
	case e.Operator != nil:
		count := 0
		if *e.Operator == AndBinaryOperator || *e.Operator == OrBinaryOperator || *e.Operator == XorBinaryOperator {
			count++
		}
		return count + e.LeftOperand.booleanOperatorCount() + e.RightOperand.booleanOperatorCount()
	}
	return 0
}

// NewStringExpression creates a new expression holding a constant string.
func NewStringExpression(str string) *Expression {
	return &Expression{
		Value: variable.NewString(str),
	}
}

// FunctionCall is used to represent the call of a function from a YarnSpinner script.
type FunctionCall struct {
	FunctionID string
	Arguments  []*Expression
}
