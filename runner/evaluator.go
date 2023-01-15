package runner

import (
	"fmt"
	"math"

	"github.com/RemiEven/ysgo/internal/tree"
	"github.com/RemiEven/ysgo/variable"
)

func evaluateExpression(e *tree.Expression, retriever variable.Retriever, caller functionCaller) (*variable.Value, error) {
	switch {
	case e.VariableID != nil:
		value, ok := retriever.GetValue(*e.VariableID)
		if !ok {
			return nil, fmt.Errorf("variable [%v] not found in storage", *e.VariableID)
		}
		return value, nil
	case e.FunctionCall != nil:
		return evaluateFunctionCall(e.FunctionCall, retriever, caller)
	case e.Value != nil:
		return e.Value, nil
	case e.NegativeExpression != nil:
		negativeExpressionValue, err := evaluateExpression(e.NegativeExpression, retriever, caller)
		if err != nil {
			return nil, fmt.Errorf("failed to evaluate negative expression: %w", err)
		} else if negativeExpressionValue.Number == nil {
			return nil, fmt.Errorf("cannot compute the negative value of something that is not a number")
		}
		return variable.NewNumber(-*negativeExpressionValue.Number), nil
	case e.NotExpression != nil:
		notExpressionValue, err := evaluateExpression(e.NotExpression, retriever, caller)
		if err != nil {
			return nil, fmt.Errorf("failed to evaluate not-expression: %w", err)
		} else if notExpressionValue.Boolean == nil {
			return nil, fmt.Errorf("cannot compute the negation of something that is not a boolean")
		}
		return variable.NewBoolean(!*notExpressionValue.Boolean), nil
	case e.Operator != nil:
		return evaluateBinaryOperation(*e.Operator, e.LeftOperand, e.RightOperand, retriever, caller)
	}
	return nil, nil
}

func evaluateBinaryOperation(operator int, leftOperand, rightOperand *tree.Expression, retriever variable.Retriever, caller functionCaller) (*variable.Value, error) {
	leftOperandValue, err := evaluateExpression(leftOperand, retriever, caller)
	if err != nil {
		return nil, fmt.Errorf("failed to evaluate left operand of expression: %w", err)
	}

	// lazy-evaluation if possible
	if operator == tree.AndBinaryOperator {
		if leftOperandValue.Boolean == nil {
			return nil, fmt.Errorf("cannot evaluate AND expression if left operand is not a boolean")
		}
		if !*leftOperandValue.Boolean {
			return leftOperandValue, nil
		}
	}
	if operator == tree.OrBinaryOperator {
		if leftOperandValue.Boolean == nil {
			return nil, fmt.Errorf("cannot evaluate OR expression if left operand is not a boolean")
		}
		if *leftOperandValue.Boolean {
			return leftOperandValue, nil
		}
	}

	rightOperandValue, err := evaluateExpression(rightOperand, retriever, caller)
	if err != nil {
		return nil, fmt.Errorf("failed to evaluate right operand of expression: %w", err)
	}

	bothOperandsAreNumbers := leftOperandValue.Number != nil && rightOperandValue.Number != nil
	bothOperandsAreBooleans := leftOperandValue.Boolean != nil && rightOperandValue.Boolean != nil
	bothOperandsAreStrings := leftOperandValue.String != nil && rightOperandValue.String != nil

	if !(bothOperandsAreNumbers || bothOperandsAreBooleans || bothOperandsAreStrings) {
		return nil, fmt.Errorf("cannot evaluate a binary expression if operands have different types")
	}

	switch operator {
	case tree.MultiplicationBinaryOperator:
		if bothOperandsAreNumbers {
			return variable.NewNumber((*leftOperandValue.Number) * (*rightOperandValue.Number)), nil
		}
		return nil, fmt.Errorf("cannot multiply two values that are not both numbers")
	case tree.DivisionBinaryOperator:
		if bothOperandsAreNumbers {
			return variable.NewNumber((*leftOperandValue.Number) / (*rightOperandValue.Number)), nil
		}
		return nil, fmt.Errorf("cannot divide two values that are not both numbers")
	case tree.ModuloBinaryOperator:
		if bothOperandsAreNumbers {
			return variable.NewNumber(math.Mod(*leftOperandValue.Number, *rightOperandValue.Number)), nil
		}
		return nil, fmt.Errorf("cannot divide two values that are not both numbers")
	case tree.AdditionBinaryOperator:
		if bothOperandsAreNumbers {
			return variable.NewNumber((*leftOperandValue.Number) + (*rightOperandValue.Number)), nil
		} else if bothOperandsAreStrings {
			return variable.NewString((*leftOperandValue.String) + (*rightOperandValue.String)), nil
		}
		return nil, fmt.Errorf("cannot add two values that are not either both numbers or both strings")
	case tree.SubtractionBinaryOperator:
		if bothOperandsAreNumbers {
			return variable.NewNumber((*leftOperandValue.Number) - (*rightOperandValue.Number)), nil
		}
		return nil, fmt.Errorf("cannot substract two values that are not both numbers")
	case tree.LessThanEqualsBinaryOperator:
		if bothOperandsAreNumbers {
			return variable.NewBoolean((*leftOperandValue.Number) <= (*rightOperandValue.Number)), nil
		}
		return nil, fmt.Errorf("cannot compare two values with LTE that are not both numbers")
	case tree.GreaterThanEqualsBinaryOperator:
		if bothOperandsAreNumbers {
			return variable.NewBoolean((*leftOperandValue.Number) >= (*rightOperandValue.Number)), nil
		}
		return nil, fmt.Errorf("cannot compare two values with GTE that are not both numbers")
	case tree.LessBinaryOperator:
		if bothOperandsAreNumbers {
			return variable.NewBoolean((*leftOperandValue.Number) < (*rightOperandValue.Number)), nil
		}
		return nil, fmt.Errorf("cannot compare two values with LT that are not both numbers")
	case tree.GreaterBinaryOperator:
		if bothOperandsAreNumbers {
			return variable.NewBoolean((*leftOperandValue.Number) > (*rightOperandValue.Number)), nil
		}
		return nil, fmt.Errorf("cannot compare two values with GT that are not both numbers")
	case tree.EqualsBinaryOperator:
		if bothOperandsAreNumbers {
			return variable.NewBoolean((*leftOperandValue.Number) == (*rightOperandValue.Number)), nil
		} else if bothOperandsAreBooleans {
			return variable.NewBoolean((*leftOperandValue.Boolean) == (*rightOperandValue.Boolean)), nil
		} else if bothOperandsAreStrings {
			return variable.NewBoolean((*leftOperandValue.String) == (*rightOperandValue.String)), nil
		}
	case tree.NotEqualsBinaryOperator:
		if bothOperandsAreNumbers {
			return variable.NewBoolean((*leftOperandValue.Number) != (*rightOperandValue.Number)), nil
		} else if bothOperandsAreBooleans {
			return variable.NewBoolean((*leftOperandValue.Boolean) != (*rightOperandValue.Boolean)), nil
		} else if bothOperandsAreStrings {
			return variable.NewBoolean((*leftOperandValue.String) != (*rightOperandValue.String)), nil
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
			return variable.NewBoolean(xor(*leftOperandValue.Boolean, *rightOperandValue.Boolean)), nil
		}
		return nil, fmt.Errorf("cannot XOR two values that are not both booleans")
	}

	return nil, fmt.Errorf("unknown operator")
}

func xor(a, b bool) bool {
	return (a && !b) || (!a && b)
}

func evaluateFunctionCall(call *tree.FunctionCall, retriever variable.Retriever, caller functionCaller) (*variable.Value, error) {
	evaluatedArgs := make([]*variable.Value, 0, len(call.Arguments))
	for i := range call.Arguments {
		evaluatedArg, err := evaluateExpression(call.Arguments[i], retriever, caller)
		if err != nil {
			return nil, fmt.Errorf("failed to evaluate argument number %d: %w", i, err)
		}
		evaluatedArgs = append(evaluatedArgs, evaluatedArg)
	}

	result, err := caller.Call(call.FunctionID, evaluatedArgs)
	if err != nil {
		return nil, fmt.Errorf("call to function %s failed: %w", call.FunctionID, err)
	}
	return result, nil
}
