package variable

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

// Expression represents something that can be evaluated to a YarnSpinner value when a Dialogue is run.
type Expression struct {
	Value              *Value        `json:"val,omitempty"`
	VariableID         *string       `json:"var,omitempty"`
	FunctionCall       *FunctionCall `json:"fc,omitempty"`
	NegativeExpression *Expression   `json:"ne,omitempty"`
	NotExpression      *Expression   `json:"not-e,omitempty"`
	LeftOperand        *Expression   `json:"lo,omitempty"`
	RightOperand       *Expression   `json:"ro,omitempty"`
	Operator           *int          `json:"o,omitempty"`
}

// IsConstant returns whether a given expression is a constant or not.
func (e *Expression) IsConstant() bool {
	switch {
	case e == nil:
		return false
	case e.Value != nil:
		return true
	case e.NegativeExpression != nil:
		return e.NegativeExpression.IsConstant()
	case e.NotExpression != nil:
		return e.NotExpression.IsConstant()
	case e.Operator != nil:
		return e.LeftOperand.IsConstant() && e.RightOperand.IsConstant()
	}
	return false
}

// ComplexityScore returns a score representing how complex an expression is.
// This value can be useful for saliency strategies.
func (e *Expression) ComplexityScore() int {
	if e == nil {
		return 0
	}
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
		Value: NewString(str),
	}
}

// FunctionCall is used to represent the call of a function from a YarnSpinner script.
type FunctionCall struct {
	FunctionID string
	Arguments  []*Expression
}
