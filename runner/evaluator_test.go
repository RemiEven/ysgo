package runner

import (
	"testing"

	"github.com/RemiEven/ysgo/internal/testutils"
	"github.com/RemiEven/ysgo/internal/tree"
	"github.com/RemiEven/ysgo/variable"
)

func TestEvaluateExpression(t *testing.T) {
	tests := map[string]struct {
		expression    *tree.Expression
		expectedValue *variable.Value
		expectedError error
	}{
		// TODO: complete this
		"simple number": {
			expression:    simpleNumberExpression(12.4),
			expectedValue: variable.NewNumber(12.4),
		},
		"negative number": {
			expression: &tree.Expression{
				NegativeExpression: simpleNumberExpression(56.7),
			},
			expectedValue: variable.NewNumber(-56.7),
		},
		"simple boolean": {
			expression:    simpleBooleanExpression(true),
			expectedValue: variable.NewBoolean(true),
		},
		"logic not boolean": {
			expression: &tree.Expression{
				NotExpression: simpleBooleanExpression(false),
			},
			expectedValue: variable.NewBoolean(true),
		},
		"number LTE comparison": {
			expression: &tree.Expression{
				Operator:     toPointer(tree.LessThanEqualsBinaryOperator),
				LeftOperand:  simpleNumberExpression(2),
				RightOperand: simpleNumberExpression(4),
			},
			expectedValue: variable.NewBoolean(true),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actualValue, err := evaluateExpression(test.expression, nil, nil) // TODO: use an in memory variable storer here when testing for variable evaluation
			if !testutils.ErrorEqual(err, test.expectedError) {
				t.Errorf("unexpected err: wanted [%v], got [%v]", test.expectedError, err)
			}
			if diff := testutils.DeepEqual(actualValue, test.expectedValue); diff != "" {
				t.Errorf("unexpected value: " + diff)
			}
		})
	}
}

func simpleNumberExpression(value float64) *tree.Expression {
	return &tree.Expression{
		Value: variable.NewNumber(value),
	}
}

func simpleBooleanExpression(value bool) *tree.Expression {
	return &tree.Expression{
		Value: variable.NewBoolean(value),
	}
}

func toPointer[T any](value T) *T {
	return &value
}
