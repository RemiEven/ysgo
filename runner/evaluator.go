package runner

import (
	"fmt"
	"math"

	"github.com/RemiEven/ysgo/tree"
)

func evaluateExpression(e *tree.Expression, retriever variableRetriever) (*tree.Value, error) {
	switch {
	case e.Value != nil && e.Value.VariableID != nil:
		value, ok := retriever.GetValue(*e.Value.VariableID)
		if !ok {
			return nil, fmt.Errorf("variable [%v] not found in storage", *e.Value.VariableID)
		}
		return value, nil
	case e.Value != nil:
		return e.Value, nil
	case e.NegativeExpression != nil:
		negativeExpressionValue, err := evaluateExpression(e.NegativeExpression, retriever)
		if err != nil {
			return nil, fmt.Errorf("failed to evaluate negative expression: %w", err)
		} else if !negativeExpressionValue.IsNumber() {
			return nil, fmt.Errorf("cannot compute the negative value of something that is not a number")
		}
		return tree.NewNumberValue(-*negativeExpressionValue.Number), nil
	case e.NotExpression != nil:
		notExpressionValue, err := evaluateExpression(e.NotExpression, retriever)
		if err != nil {
			return nil, fmt.Errorf("failed to evaluate not-expression: %w", err)
		} else if !notExpressionValue.IsBoolean() {
			return nil, fmt.Errorf("cannot compute the negation of something that is not a boolean")
		}
		return tree.NewBooleanValue(!*notExpressionValue.Boolean), nil
	case e.Operator != nil:
		return evaluateBinaryOperation(*e.Operator, e.LeftOperand, e.RightOperand, retriever)
	}
	return nil, nil
}

func evaluateBinaryOperation(operator int, leftOperand, rightOperand *tree.Expression, retriever variableRetriever) (*tree.Value, error) {
	leftOperandValue, err := evaluateExpression(leftOperand, retriever)
	if err != nil {
		return nil, fmt.Errorf("failed to evaluate left operand of expression: %w", err)
	}

	// lazy-evaluation if possible
	if operator == tree.AndBinaryOperator {
		if !leftOperandValue.IsBoolean() {
			return nil, fmt.Errorf("cannot evaluate AND expression if left operand is not a boolean")
		}
		if !*leftOperandValue.Boolean {
			return leftOperandValue, nil
		}
	}
	if operator == tree.OrBinaryOperator {
		if !leftOperandValue.IsBoolean() {
			return nil, fmt.Errorf("cannot evaluate OR expression if left operand is not a boolean")
		}
		if *leftOperandValue.Boolean {
			return leftOperandValue, nil
		}
	}

	rightOperandValue, err := evaluateExpression(rightOperand, retriever)
	if err != nil {
		return nil, fmt.Errorf("failed to evaluate right operand of expression: %w", err)
	}

	bothOperandsAreNumbers := leftOperandValue.IsNumber() && rightOperandValue.IsNumber()
	bothOperandsAreBooleans := leftOperandValue.IsBoolean() && rightOperandValue.IsBoolean()
	bothOperandsAreStrings := leftOperandValue.IsString() && rightOperandValue.IsString()

	if !(bothOperandsAreNumbers || bothOperandsAreBooleans || bothOperandsAreStrings) {
		return nil, fmt.Errorf("cannot evaluate a binary expression if operands have different types")
	}

	switch operator {
	case tree.MultiplicationBinaryOperator:
		if bothOperandsAreNumbers {
			return tree.NewNumberValue((*leftOperandValue.Number) * (*rightOperandValue.Number)), nil
		}
		return nil, fmt.Errorf("cannot multiply two values that are not both numbers")
	case tree.DivisionBinaryOperator:
		if bothOperandsAreNumbers {
			return tree.NewNumberValue((*leftOperandValue.Number) / (*rightOperandValue.Number)), nil
		}
		return nil, fmt.Errorf("cannot divide two values that are not both numbers")
	case tree.ModuloBinaryOperator:
		if bothOperandsAreNumbers {
			return tree.NewNumberValue(math.Mod(*leftOperandValue.Number, *rightOperandValue.Number)), nil
		}
		return nil, fmt.Errorf("cannot divide two values that are not both numbers")
	case tree.AdditionBinaryOperator:
		if bothOperandsAreNumbers {
			return tree.NewNumberValue((*leftOperandValue.Number) + (*rightOperandValue.Number)), nil
		} else if bothOperandsAreStrings {
			return tree.NewStringValue((*leftOperandValue.String) + (*rightOperandValue.String)), nil
		}
		return nil, fmt.Errorf("cannot add two values that are not either both numbers or both strings")
	case tree.SubtractionBinaryOperator:
		if bothOperandsAreNumbers {
			return tree.NewNumberValue((*leftOperandValue.Number) - (*rightOperandValue.Number)), nil
		}
		return nil, fmt.Errorf("cannot substract two values that are not both numbers")
	case tree.LessThanEqualsBinaryOperator:
		if bothOperandsAreNumbers {
			return tree.NewBooleanValue((*leftOperandValue.Number) <= (*rightOperandValue.Number)), nil
		}
		return nil, fmt.Errorf("cannot compare two values with LTE that are not both numbers")
	case tree.GreaterThanEqualsBinaryOperator:
		if bothOperandsAreNumbers {
			return tree.NewBooleanValue((*leftOperandValue.Number) >= (*rightOperandValue.Number)), nil
		}
		return nil, fmt.Errorf("cannot compare two values with GTE that are not both numbers")
	case tree.LessBinaryOperator:
		if bothOperandsAreNumbers {
			return tree.NewBooleanValue((*leftOperandValue.Number) < (*rightOperandValue.Number)), nil
		}
		return nil, fmt.Errorf("cannot compare two values with LT that are not both numbers")
	case tree.GreaterBinaryOperator:
		if bothOperandsAreNumbers {
			return tree.NewBooleanValue((*leftOperandValue.Number) > (*rightOperandValue.Number)), nil
		}
		return nil, fmt.Errorf("cannot compare two values with GT that are not both numbers")
	case tree.EqualsBinaryOperator:
		if bothOperandsAreNumbers {
			return tree.NewBooleanValue((*leftOperandValue.Number) == (*rightOperandValue.Number)), nil
		} else if bothOperandsAreBooleans {
			return tree.NewBooleanValue((*leftOperandValue.Boolean) == (*rightOperandValue.Boolean)), nil
		} else if bothOperandsAreStrings {
			return tree.NewBooleanValue((*leftOperandValue.String) == (*rightOperandValue.String)), nil
		}
	case tree.NotEqualsBinaryOperator:
		if bothOperandsAreNumbers {
			return tree.NewBooleanValue((*leftOperandValue.Number) != (*rightOperandValue.Number)), nil
		} else if bothOperandsAreBooleans {
			return tree.NewBooleanValue((*leftOperandValue.Boolean) != (*rightOperandValue.Boolean)), nil
		} else if bothOperandsAreStrings {
			return tree.NewBooleanValue((*leftOperandValue.String) != (*rightOperandValue.String)), nil
		}
	case tree.AndBinaryOperator:
		if bothOperandsAreBooleans {
			return rightOperandValue, nil // leftOperandValue has already been accounted for when checking if lazy-evaluation was possible
		}
		return nil, fmt.Errorf("cannot AND two values that are not both booleans")
	case tree.OrBinaryOperator:
		if bothOperandsAreBooleans {
			return rightOperandValue, nil // leftOperandValue has already been accounted for when checking if lazy-evaluation was possible
		}
		return nil, fmt.Errorf("cannot OR two values that are not both booleans")
	case tree.XorBinaryOperator:
		if bothOperandsAreBooleans {
			return tree.NewBooleanValue(xor(*leftOperandValue.Boolean, *rightOperandValue.Boolean)), nil
		}
		return nil, fmt.Errorf("cannot XOR two values that are not both booleans")
	}

	return nil, fmt.Errorf("unknown operator")
}

func toPointer[T any](value T) *T {
	return &value
}

func xor(a, b bool) bool {
	return (a && !b) || (!a && b)
}
