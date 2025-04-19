package ysgo

import (
	"fmt"
	"math"

	"github.com/remieven/ysgo/variable"
)

func evaluateExpression(e *variable.Expression, retriever variable.Retriever, caller functionCaller, smartVariables map[string]*variable.Expression) (*variable.Value, error) {
	switch {
	case e.VariableID != nil:
		if expression, ok := smartVariables[*e.VariableID]; ok {
			return evaluateExpression(expression, retriever, caller, smartVariables)
		}
		value, ok := retriever.GetValue(*e.VariableID)
		if !ok {
			return nil, fmt.Errorf("variable [%v] not found in storage", *e.VariableID)
		}
		return value, nil
	case e.FunctionCall != nil:
		return evaluateFunctionCall(e.FunctionCall, retriever, caller, smartVariables)
	case e.Value != nil:
		return e.Value, nil
	case e.NegativeExpression != nil:
		negativeExpressionValue, err := evaluateExpression(e.NegativeExpression, retriever, caller, smartVariables)
		if err != nil {
			return nil, fmt.Errorf("failed to evaluate negative expression: %w", err)
		} else if negativeExpressionValue.Number == nil {
			return nil, fmt.Errorf("cannot compute the negative value of something that is not a number")
		}
		return variable.NewNumber(-*negativeExpressionValue.Number), nil
	case e.NotExpression != nil:
		notExpressionValue, err := evaluateExpression(e.NotExpression, retriever, caller, smartVariables)
		if err != nil {
			return nil, fmt.Errorf("failed to evaluate not-expression: %w", err)
		} else if notExpressionValue.Boolean == nil {
			return nil, fmt.Errorf("cannot compute the negation of something that is not a boolean")
		}
		return variable.NewBoolean(!*notExpressionValue.Boolean), nil
	case e.Operator != nil:
		return evaluateBinaryOperation(*e.Operator, e.LeftOperand, e.RightOperand, retriever, caller, smartVariables)
	}
	return nil, nil
}

func evaluateBinaryOperation(operator int, leftOperand, rightOperand *variable.Expression, retriever variable.Retriever, caller functionCaller, smartVariables map[string]*variable.Expression) (*variable.Value, error) {
	leftOperandValue, err := evaluateExpression(leftOperand, retriever, caller, smartVariables)
	if err != nil {
		return nil, fmt.Errorf("failed to evaluate left operand of expression: %w", err)
	}

	// lazy-evaluation if possible
	if operator == variable.AndBinaryOperator {
		if leftOperandValue.Boolean == nil {
			return nil, fmt.Errorf("cannot evaluate AND expression if left operand is not a boolean")
		}
		if !*leftOperandValue.Boolean {
			return leftOperandValue, nil
		}
	}
	if operator == variable.OrBinaryOperator {
		if leftOperandValue.Boolean == nil {
			return nil, fmt.Errorf("cannot evaluate OR expression if left operand is not a boolean")
		}
		if *leftOperandValue.Boolean {
			return leftOperandValue, nil
		}
	}

	rightOperandValue, err := evaluateExpression(rightOperand, retriever, caller, smartVariables)
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
	case variable.MultiplicationBinaryOperator:
		if bothOperandsAreNumbers {
			return variable.NewNumber((*leftOperandValue.Number) * (*rightOperandValue.Number)), nil
		}
		return nil, fmt.Errorf("cannot multiply two values that are not both numbers")
	case variable.DivisionBinaryOperator:
		if bothOperandsAreNumbers {
			return variable.NewNumber((*leftOperandValue.Number) / (*rightOperandValue.Number)), nil
		}
		return nil, fmt.Errorf("cannot divide two values that are not both numbers")
	case variable.ModuloBinaryOperator:
		if bothOperandsAreNumbers {
			return variable.NewNumber(math.Mod(*leftOperandValue.Number, *rightOperandValue.Number)), nil
		}
		return nil, fmt.Errorf("cannot divide two values that are not both numbers")
	case variable.AdditionBinaryOperator:
		if bothOperandsAreNumbers {
			return variable.NewNumber((*leftOperandValue.Number) + (*rightOperandValue.Number)), nil
		} else if bothOperandsAreStrings {
			return variable.NewString((*leftOperandValue.String) + (*rightOperandValue.String)), nil
		}
		return nil, fmt.Errorf("cannot add two values that are not either both numbers or both strings")
	case variable.SubtractionBinaryOperator:
		if bothOperandsAreNumbers {
			return variable.NewNumber((*leftOperandValue.Number) - (*rightOperandValue.Number)), nil
		}
		return nil, fmt.Errorf("cannot subtract two values that are not both numbers")
	case variable.LessThanEqualsBinaryOperator:
		if bothOperandsAreNumbers {
			return variable.NewBoolean((*leftOperandValue.Number) <= (*rightOperandValue.Number)), nil
		}
		return nil, fmt.Errorf("cannot compare two values with LTE that are not both numbers")
	case variable.GreaterThanEqualsBinaryOperator:
		if bothOperandsAreNumbers {
			return variable.NewBoolean((*leftOperandValue.Number) >= (*rightOperandValue.Number)), nil
		}
		return nil, fmt.Errorf("cannot compare two values with GTE that are not both numbers")
	case variable.LessBinaryOperator:
		if bothOperandsAreNumbers {
			return variable.NewBoolean((*leftOperandValue.Number) < (*rightOperandValue.Number)), nil
		}
		return nil, fmt.Errorf("cannot compare two values with LT that are not both numbers")
	case variable.GreaterBinaryOperator:
		if bothOperandsAreNumbers {
			return variable.NewBoolean((*leftOperandValue.Number) > (*rightOperandValue.Number)), nil
		}
		return nil, fmt.Errorf("cannot compare two values with GT that are not both numbers")
	case variable.EqualsBinaryOperator:
		if bothOperandsAreNumbers {
			return variable.NewBoolean((*leftOperandValue.Number) == (*rightOperandValue.Number)), nil
		} else if bothOperandsAreBooleans {
			return variable.NewBoolean((*leftOperandValue.Boolean) == (*rightOperandValue.Boolean)), nil
		} else if bothOperandsAreStrings {
			return variable.NewBoolean((*leftOperandValue.String) == (*rightOperandValue.String)), nil
		}
	case variable.NotEqualsBinaryOperator:
		if bothOperandsAreNumbers {
			return variable.NewBoolean((*leftOperandValue.Number) != (*rightOperandValue.Number)), nil
		} else if bothOperandsAreBooleans {
			return variable.NewBoolean((*leftOperandValue.Boolean) != (*rightOperandValue.Boolean)), nil
		} else if bothOperandsAreStrings {
			return variable.NewBoolean((*leftOperandValue.String) != (*rightOperandValue.String)), nil
		}
	case variable.AndBinaryOperator:
		if bothOperandsAreBooleans {
			return rightOperandValue, nil // leftOperandValue has already been accounted for when checking if lazy-evaluation was possible
		}
		return nil, fmt.Errorf("cannot AND two values that are not both booleans")
	case variable.OrBinaryOperator:
		if bothOperandsAreBooleans {
			return rightOperandValue, nil // leftOperandValue has already been accounted for when checking if lazy-evaluation was possible
		}
		return nil, fmt.Errorf("cannot OR two values that are not both booleans")
	case variable.XorBinaryOperator:
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

func evaluateFunctionCall(call *variable.FunctionCall, retriever variable.Retriever, caller functionCaller, smartVariables map[string]*variable.Expression) (*variable.Value, error) {
	evaluatedArgs := make([]*variable.Value, 0, len(call.Arguments))
	for i := range call.Arguments {
		evaluatedArg, err := evaluateExpression(call.Arguments[i], retriever, caller, smartVariables)
		if err != nil {
			return nil, fmt.Errorf("failed to evaluate argument number %d: %w", i, err)
		}
		evaluatedArgs = append(evaluatedArgs, evaluatedArg)
	}

	result, err := caller.call(call.FunctionID, evaluatedArgs)
	if err != nil {
		return nil, fmt.Errorf("call to function %s failed: %w", call.FunctionID, err)
	}
	return result, nil
}
