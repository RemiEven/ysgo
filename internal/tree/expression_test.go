package tree

import (
	"strings"
	"testing"
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

	for name, expectedValue := range tests {
		t.Run(name, func(t *testing.T) {
			script := `
title: start
---
{` + name + `}
===
`
			dialogue, err := FromReader(strings.NewReader(script))
			if err != nil {
				t.Errorf("failed to parse expression: %v", err)
				return
			}
			expression := dialogue.Nodes[0].Statements[0].LineStatement.Text.Elements[0].Expression
			if actual := expression.ComplexityScore(); actual != expectedValue {
				t.Errorf("unexpected result: got [%v], wanted [%v]", actual, expectedValue)
				return
			}
		})
	}
}
