package runner_test

import (
	"errors"
	"testing"

	"github.com/RemiEven/ysgo/runner"
	"github.com/RemiEven/ysgo/tree"
)

func TestRunnerPlan(t *testing.T) {
	tests := []string{
		"Escaping",
		"Functions",
		"IfStatements",
		"Lines",
		"VisitCount",
		"Visited",
	}

	for _, test := range tests {
		t.Run(test, func(t *testing.T) {
			plan, err := parseTestPlan("testdata/plans/" + test + ".testplan")
			if err != nil {
				t.Errorf("failed to parse test plan: %v", err)
				return
			}

			di, err := tree.FromFile("testdata/plans/" + test + ".yarn")
			if err != nil {
				t.Errorf("failed to parse dialogue from file: %v", err)
				return
			}

			dr, err := runner.NewDialogueRunner(di, nil, "")
			if err != nil {
				t.Errorf("failed to create dialogue runner: %v", err)
				return
			}

			dr.ConvertAndAddCommand("assert", func(b bool) error {
				if !b {
					return errors.New("assertion failed")
				}
				return nil
			})
			dr.ConvertAndAddFunction("add_three_operands", func(a, b, c float64) float64 {
				return a + b + c
			})

			testPlanIndex := 0
			for testPlanIndex < len(plan) {
				testPlanStep := plan[testPlanIndex]
				choice := 0
				if testPlanIndex != 0 && testPlanStep.intValue != 0 {
					choice = testPlanStep.intValue - 1
				}
				element, err := dr.Next(choice)
				if err != nil {
					t.Errorf("got an error while running dialog: %v", err)
					return
				}

				switch testPlanStep.stepType {
				case stepTypeLine:
					if line := element.Line; line == nil {
						t.Errorf("expected a line at test plan step [%v] but got none", testPlanIndex)
						return
					} else if actual, expected := line.Text, testPlanStep.stringValue; actual != expected {
						t.Errorf("unexpected text for line at step [%v]: got [%s], wanted [%s]", testPlanIndex, actual, expected)
						return
					}
				}

				testPlanIndex++
			}
		})
	}

}
