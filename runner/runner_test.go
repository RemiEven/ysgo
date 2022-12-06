package runner

import (
	"strconv"
	"testing"

	"github.com/RemiEven/ysgo/markup"
	"github.com/RemiEven/ysgo/testutils"
	"github.com/RemiEven/ysgo/tree"
)

func TestRunner(t *testing.T) {
	tests := map[string]struct {
		script                   string
		inputs                   []int
		expectedDialogueElements []DialogueElement
		expectedLastOK           bool
		expectedLastErr          error
	}{
		"simple": {
			script: "simple",
			inputs: []int{0, 0, 0, 0},
			expectedDialogueElements: []DialogueElement{
				simpleTextDialogueElement("Hi there! What do you feel like doing today?"),
				simpleTextDialogueElement("I want to go swimming."),
				simpleTextDialogueElement("Sounds good!"),
			},
		},
		"condition statement": {
			script: "condition_statement",
			inputs: []int{0, 0, 0},
			expectedDialogueElements: []DialogueElement{
				simpleTextDialogueElementFromCharacter("Player", "I'd like to buy a pie!"),
				simpleTextDialogueElementFromCharacter("Baker", "Well, you can't afford one!"),
			},
		},
		"visited and visited_count": {
			script: "visited",
			inputs: []int{0, 1, 0, 0, 0, 0, 0, 0, 3, 0},
			expectedDialogueElements: []DialogueElement{
				simpleOptionDialogueElement("I want to go to the lake.", "I want to go to the forest.", "I want to go to the hill.", "I've seen enough."),
				simpleTextDialogueElement("This is a nice forest!"),
				simpleOptionDialogueElement("I want to go to the lake.", "I want to go to the forest.", "I want to go to the hill.", "I've seen enough."),
				simpleTextDialogueElement("This is a nice lake!"),
				simpleOptionDialogueElement("I want to go to the lake.", "I want to go to the forest.", "I want to go to the hill.", "I've seen enough."),
				simpleTextDialogueElement("This is a nice lake!"),
				simpleTextDialogueElement("We've been here before!"),
				simpleOptionDialogueElement("I want to go to the lake.", "I want to go to the forest.", "I want to go to the hill.", "I've seen enough."),
				simpleTextDialogueElement("We've been 2 times to the lake, 1 times to the forest, and 0 times to the hill."),
			},
		},
		"options without statements": {
			script: "options_without_statements",
			inputs: []int{0, 0, 1},
			expectedDialogueElements: []DialogueElement{
				simpleTextDialogueElement("Hi there! What do you feel like doing today?"),
				simpleOptionDialogueElement("I want to go swimming.", "I want to go hiking."),
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			dialogueTree, err := tree.FromFile("testdata/" + test.script + ".yarn")
			if err != nil {
				t.Errorf("failed to read testdata file: %v", err)
				return
			}

			runner, err := NewDialogueRunner(dialogueTree, nil, "abcd")
			if err != nil {
				t.Errorf("failed to create runner: %v", err)
				return
			}

			for i, input := range test.inputs {
				stepPrefix := "step " + strconv.Itoa(i) + ": "
				isLastStep := i == len(test.inputs)-1
				dialogueElement, err := runner.Next(input)
				ok := dialogueElement != nil && err == nil
				switch {
				case err != nil && isLastStep:
					if !testutils.ErrorEqual(err, test.expectedLastErr) {
						t.Errorf(stepPrefix+"unexpected error at the last step: got [%v], wanted [%v]", err, test.expectedLastErr)
						return
					}
				case err != nil:
					t.Errorf(stepPrefix+"unexpected error before the last step: [%v]", err)
					return
				case !ok && !isLastStep:
					t.Errorf(stepPrefix + "received a nok before reaching the last expected step")
					return
				case isLastStep && ok != test.expectedLastOK:
					t.Errorf(stepPrefix+"unexpected ok at the last step: got [%v], wanted [%v]", ok, test.expectedLastOK)
					return
				case !isLastStep:
					if diff := testutils.DeepEqual(dialogueElement, &test.expectedDialogueElements[i]); diff != "" {
						t.Errorf(stepPrefix + "unexpected dialogueElement: " + diff)
						return
					}
				}
			}
		})
	}
}

func simpleTextDialogueElement(text string) DialogueElement {
	return DialogueElement{
		Line: &markup.ParseResult{Text: text, Attributes: []markup.Attribute{}},
	}
}

func simpleTextDialogueElementFromCharacter(character, text string) DialogueElement {
	return DialogueElement{
		Line: &markup.ParseResult{Text: character + ": " + text, Attributes: []markup.Attribute{
			{
				Position: 0,
				Length:   len(character) + 2,
				Name:     "character",
				Properties: map[string]markup.Value{
					"name": {
						StringValue: character,
						ValueType:   markup.ValueTypeString,
					},
				},
			},
		}},
	}
}

func simpleOptionDialogueElement(options ...string) DialogueElement {
	element := DialogueElement{
		Options: []DialogueOption{},
	}

	for _, option := range options {
		element.Options = append(element.Options, DialogueOption{
			Line: &markup.ParseResult{Text: option, Attributes: []markup.Attribute{}},
		})
	}

	return element
}
