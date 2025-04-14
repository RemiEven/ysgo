package ysgo_test

import (
	"os"
	"strings"
	"testing"

	"github.com/remieven/ysgo"
	"github.com/remieven/ysgo/internal/testutils"
	"github.com/remieven/ysgo/variable"
)

func TestRunnerPlan(t *testing.T) {
	tests := []string{
		"Commands",
		"DecimalNumbers",
		"Detours",
		"Escaping",
		"Expressions",
		"FormatFunctions",
		"Functions",
		"Identifiers",
		"IfStatements",
		"Indentation",
		"InlineExpressions",
		"Jumps",
		"LineGroups",
		"Lines",
		"NodeHeaders",
		"ShortcutOptions",
		"Smileys",
		"Types",
		"VariableStorage",
		"VisitCount",
		"Visited",
		"VisitTracking",
	}

	for _, test := range tests {
		t.Run(test, func(t *testing.T) {
			plan, err := parseTestPlan("testdata/" + test + ".testplan")
			if err != nil {
				t.Errorf("failed to parse test plan: %v", err)
				return
			}

			reader, err := os.Open("testdata/" + test + ".yarn")
			if err != nil {
				t.Errorf("failed to open dialogue file: %v", err)
				return
			}
			defer reader.Close()

			dr, err := ysgo.NewDialogueRunner(nil, "", reader)
			if err != nil {
				t.Errorf("failed to create dialogue runner: %v", err)
				return
			}

			dr.ConvertAndAddFunction("assert", func(b bool) {
				if !b {
					t.Errorf("assertion failed")
				}
			})
			dr.ConvertAndAddFunction("add_three_operands", func(a, b, c float64) float64 {
				return a + b + c
			})

			commandOutputs := []string{}
			addTestCommand := func(commandName string) {
				dr.AddCommand(commandName, func(args []*variable.Value) <-chan error {
					output := commandName
					for _, arg := range args {
						output += " " + arg.ToString()
					}
					commandOutputs = append(commandOutputs, output)
					ch := make(chan error, 1)
					ch <- nil
					return ch
				})
			}

			for _, commandName := range []string{"number", "expression", "string", "bool", "variable", "flip", "toggle", "settings", "iffy", "nulled", "orion", "andorian", "note", "isActive", "p", "hide", "show"} {
				addTestCommand(commandName)
			}

			expectedCommandOutputs := make([]string, 0, len(commandOutputs))

			testPlanIndex := 0
			choice := 0
		planStepLoop:
			for testPlanIndex < len(plan) {
				testPlanStep := plan[testPlanIndex]

				switch testPlanStep.stepType {
				case stepTypeCommand:
					expectedCommandOutputs = append(expectedCommandOutputs, testPlanStep.stringValue)
				case stepTypeLine:
					element, err := dr.Next(choice)
					if err != nil {
						t.Errorf("got an error while running dialog: %v", err)
						return
					}
					choice = 0
					if element == nil {
						t.Errorf("expected a line at test plan step [%v] but got none", testPlanIndex)
						return
					} else if line := element.Line; line == nil {
						t.Errorf("expected a line at test plan step [%v] but got none", testPlanIndex)
						return
					} else if actual, expected := line.Text, testPlanStep.stringValue; actual != expected {
						t.Errorf("unexpected text for line at step [%v]: got [%s], wanted [%s]", testPlanIndex, actual, expected)
						return
					}
				case stepTypeSelect:
					choice = testPlanStep.intValue - 1
				case stepTypeOption:
					endOptionIndex := testPlanIndex
					for endOptionIndex < len(plan) && plan[endOptionIndex].stepType == stepTypeOption {
						endOptionIndex++
					}
					element, err := dr.Next(choice)
					if err != nil {
						t.Errorf("got an error while running dialog: %v", err)
						return
					}
					choice = 0
					if element.Options == nil {
						t.Errorf("expected options at test plan step [%v] but got none", testPlanIndex)
						return
					}
					if actualLen, expectedLen := len(element.Options), (endOptionIndex - testPlanIndex); actualLen != expectedLen {
						t.Errorf("got an unexpected number of options: got [%v], wanted [%v]", actualLen, expectedLen)
						return
					}
					for i := range element.Options {
						if actualLine, expectedLine := element.Options[i].Line.Text, plan[testPlanIndex+i].stringValue; actualLine != expectedLine {
							t.Errorf("got an unexpected text for option [%v]: got [%v], wanted [%v]", i, actualLine, expectedLine)
							return
						}
						if actualDisabled, expectedDisabled := element.Options[i].Disabled, !plan[testPlanIndex+i].expectOptionEnabled; actualDisabled != expectedDisabled {
							t.Errorf("got an unexpected disabled boolean value for option [%v]: got [%v], wanted [%v]", i, actualDisabled, expectedDisabled)
							return
						}
					}
					testPlanIndex = endOptionIndex - 1
				case stepTypeStop:
					break planStepLoop
				}

				testPlanIndex++
			}

			if _, err := dr.Next(choice); err != nil {
				t.Errorf("got an unexpected error at the last step of test plan: %v", err)
				return
			}

			if diff := testutils.DeepEqual(commandOutputs, expectedCommandOutputs); diff != "" {
				t.Error("unexpected command outputs: " + diff)
			}
		})
	}
}

func TestRunnerRestoreAt(t *testing.T) {
	script := `
title: node_1
---
Do you like chocolate? Here, get some icecream!
<<declare $flavor to "chocolate">>
Thanks!
<<jump node_2>>
===
title: node_2
---
You like {$flavor}, right? Here's more icecream!
Actually, I'd like raspberries this time!
<<declare $flavor to "raspberries">>
Oh ok, {$flavor} it is then!
Thanks!
===`

	dr, err := ysgo.NewDialogueRunner(nil, "", strings.NewReader(script))
	if err != nil {
		t.Errorf("failed to create dialogue runner: %v", err)
		return
	}

	assertNextLine := func(expectedLine string) {
		dialogueElement, err := dr.Next(0)
		if err != nil {
			t.Errorf("failed to get to dialogue element: %v", err)
			return
		}
		if expected, actual := expectedLine, dialogueElement.Line.Text; expected != actual {
			t.Errorf("unexpected line text: wanted [%s], got [%s]", expected, actual)
			return
		}
	}

	assertNextLine("Do you like chocolate? Here, get some icecream!")
	assertNextLine("Thanks!")
	assertNextLine("You like chocolate, right? Here's more icecream!")
	assertNextLine("Actually, I'd like raspberries this time!")
	assertNextLine("Oh ok, raspberries it is then!")

	snapshot := dr.Snapshot()

	if err := dr.RestoreAt(snapshot); err != nil {
		t.Errorf("unexpected error when restoring snapshot: %v", err)
	}

	assertNextLine("You like chocolate, right? Here's more icecream!")
}
