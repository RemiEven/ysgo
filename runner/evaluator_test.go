package runner

import (
	"testing"

	"github.com/RemiEven/ysgo/testutils"
	"github.com/RemiEven/ysgo/tree"
)

func TestEvaluateExpression(t *testing.T) {
	tests := map[string]struct {
		expression    *tree.Expression
		expectedValue *tree.Value
		expectedError error
	}{
		// TODO: complete this
		"simple number": {
			expression:    simpleNumberExpression(12.4),
			expectedValue: tree.NewNumberValue(12.4),
		},
		"negative number": {
			expression: &tree.Expression{
				NegativeExpression: simpleNumberExpression(56.7),
			},
			expectedValue: tree.NewNumberValue(-56.7),
		},
		"simple boolean": {
			expression:    simpleBooleanExpression(true),
			expectedValue: tree.NewBooleanValue(true),
		},
		"logic not boolean": {
			expression: &tree.Expression{
				NotExpression: simpleBooleanExpression(false),
			},
			expectedValue: tree.NewBooleanValue(true),
		},
		"number LTE comparison": {
			expression: &tree.Expression{
				Operator:     toPointer(tree.LessThanEqualsBinaryOperator),
				LeftOperand:  simpleNumberExpression(2),
				RightOperand: simpleNumberExpression(4),
			},
			expectedValue: tree.NewBooleanValue(true),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actualValue, err := evaluateExpression(test.expression)
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
		Value: tree.NewNumberValue(value),
	}
}

func simpleBooleanExpression(value bool) *tree.Expression {
	return &tree.Expression{
		Value: tree.NewBooleanValue(value),
	}
}
