package variable_test

import (
	"strings"
	"testing"

	"github.com/remieven/ysgo/internal/tree"
	"github.com/remieven/ysgo/variable"
)

func TestComplexityScore(t *testing.T) {
	tests := map[string]int{
		"1 < 2":                           1,
		"true":                            1,
		"$knows_player && $money < 3":     2,
		"!$knows_player && $money < -(3)": 3,
		"f(3)":                            1,
		"g(!$survived, !$knows_player && $money < -(3))": 4,
	}

	for expression, expectedValue := range tests {
		t.Run(expression, func(t *testing.T) {
			if actual := parseExpression(expression).ComplexityScore(); actual != expectedValue {
				t.Errorf("unexpected result: got [%v], wanted [%v]", actual, expectedValue)
				return
			}
		})
	}

	t.Run("nil expression", func(t *testing.T) {
		var e *variable.Expression
		if actual := e.ComplexityScore(); actual != 0 {
			t.Errorf("unexpected result: got [%v], wanted [%v]", actual, 0)
		}
	})
}

func TestIsConstant(t *testing.T) {
	tests := map[string]bool{
		"!(-1 < 2)":                       true,
		"true":                            true,
		"$knows_player && $money < 3":     false,
		"!$knows_player && $money < -(3)": false,
		"f(3)":                            false,
		"g(!$survived, !$knows_player && $money < -(3))": false,
	}

	for expression, expectedValue := range tests {
		t.Run(expression, func(t *testing.T) {
			if actual := parseExpression(expression).IsConstant(); actual != expectedValue {
				t.Errorf("unexpected result: got [%v], wanted [%v]", actual, expectedValue)
				return
			}
		})
	}

	t.Run("nil expression", func(t *testing.T) {
		var e *variable.Expression
		if actual := e.IsConstant(); actual {
			t.Errorf("unexpected result: got [%v], wanted [%v]", actual, false)
		}
	})
}

func parseExpression(expression string) *variable.Expression {
	script := `
title: start
---
{` + expression + `}
===
`
	dialogue, err := tree.FromReader(strings.NewReader(script))
	if err != nil {
		panic("failed to parse expression: " + err.Error())
	}
	return dialogue.Nodes[0].Statements[0].LineStatement.Text.Elements[0].Expression
}
