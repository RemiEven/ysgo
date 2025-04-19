package tree

import (
	"github.com/remieven/ysgo/internal/parser"
	"github.com/remieven/ysgo/variable"
)

func tokenToBinaryOperator(token int) (*int, bool) {
	operator, ok := map[int]int{
		parser.YarnSpinnerLexerOPERATOR_LOGICAL_LESS_THAN_EQUALS:    variable.LessThanEqualsBinaryOperator,
		parser.YarnSpinnerLexerOPERATOR_LOGICAL_GREATER_THAN_EQUALS: variable.GreaterThanEqualsBinaryOperator,
		parser.YarnSpinnerLexerOPERATOR_LOGICAL_EQUALS:              variable.EqualsBinaryOperator,
		parser.YarnSpinnerLexerOPERATOR_LOGICAL_LESS:                variable.LessBinaryOperator,
		parser.YarnSpinnerLexerOPERATOR_LOGICAL_GREATER:             variable.GreaterBinaryOperator,
		parser.YarnSpinnerLexerOPERATOR_LOGICAL_NOT_EQUALS:          variable.NotEqualsBinaryOperator,
		parser.YarnSpinnerLexerOPERATOR_LOGICAL_AND:                 variable.AndBinaryOperator,
		parser.YarnSpinnerLexerOPERATOR_LOGICAL_OR:                  variable.OrBinaryOperator,
		parser.YarnSpinnerLexerOPERATOR_LOGICAL_XOR:                 variable.XorBinaryOperator,
		parser.YarnSpinnerLexerOPERATOR_MATHS_ADDITION:              variable.AdditionBinaryOperator,
		parser.YarnSpinnerLexerOPERATOR_MATHS_SUBTRACTION:           variable.SubtractionBinaryOperator,
		parser.YarnSpinnerLexerOPERATOR_MATHS_MULTIPLICATION:        variable.MultiplicationBinaryOperator,
		parser.YarnSpinnerLexerOPERATOR_MATHS_DIVISION:              variable.DivisionBinaryOperator,
		parser.YarnSpinnerLexerOPERATOR_MATHS_MODULUS:               variable.ModuloBinaryOperator,
	}[token]

	return &operator, ok
}

func tokenToInplaceOperator(token int) (*int, bool) {
	operator, ok := map[int]int{
		parser.YarnSpinnerLexerOPERATOR_ASSIGNMENT:                  variable.AssignmentInPlaceOperator,
		parser.YarnSpinnerLexerOPERATOR_MATHS_MULTIPLICATION_EQUALS: variable.MultiplicationInPlaceOperator,
		parser.YarnSpinnerLexerOPERATOR_MATHS_DIVISION_EQUALS:       variable.DivisionInPlaceOperator,
		parser.YarnSpinnerLexerOPERATOR_MATHS_MODULUS_EQUALS:        variable.ModuloInPlaceOperator,
		parser.YarnSpinnerLexerOPERATOR_MATHS_ADDITION_EQUALS:       variable.AdditionInPlaceOperator,
		parser.YarnSpinnerLexerOPERATOR_MATHS_SUBTRACTION_EQUALS:    variable.SubtractionInPlaceOperator,
	}[token]

	return &operator, ok
}
