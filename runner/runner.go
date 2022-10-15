package runner

import (
	"strings"

	"github.com/RemiEven/ysgo/container"
	"github.com/RemiEven/ysgo/tree"
)

type DialogueRunner struct {
	dialogue        *tree.Dialogue
	statementsToRun container.Stack[*StatementQueue] // TODO: would the name "next steps" be better here?
	lastStatement   *tree.Statement
	variableStorer  VariableStorer
}

type DialogueElement struct { // dialogueStep? dialogueElement?
	Line    string
	Options []string
}

func (dr *DialogueRunner) textElementsToString(elements []*tree.LineFormattedTextElement) string {
	var builder strings.Builder
	for i := range elements {
		if elements[i].Text != "" {
			builder.WriteString(elements[i].Text)
		} else if expression := elements[i].Expression; expression != nil {
			value, err := evaluateExpression(expression, dr.variableStorer)
			if err != nil {
				panic(err) // TODO: actually handle error here
			}
			builder.WriteString(value.ToString())
		}
	}
	return builder.String()
}

func (dr *DialogueRunner) isWaitingForChoice() bool {
	return dr.lastStatement != nil && dr.lastStatement.ShortcutOptionStatement != nil
}

func (dr *DialogueRunner) Next(choice int) (*DialogueElement, bool) { // TODO: have nicer arguments there? like an input struct maybe?
	if dr.isWaitingForChoice() {
		dr.statementsToRun.Push(&StatementQueue{
			statements: dr.lastStatement.ShortcutOptionStatement.Options[choice].Statements,
		})
	}

	if dr.statementsToRun.Size() == 0 {
		return nil, false
	}

	statementsToRun := dr.statementsToRun.Peek()

	nextStatement, ok := statementsToRun.NextStatement()
	if !ok {
		dr.statementsToRun.Pop()
		return dr.Next(choice)
	}

	dr.lastStatement = nextStatement
	switch {
	case nextStatement.LineStatement != nil:
		return &DialogueElement{
			Line: dr.textElementsToString(nextStatement.LineStatement.Text.Elements),
		}, true
	case nextStatement.ShortcutOptionStatement != nil:
		options := make([]string, 0, len(nextStatement.ShortcutOptionStatement.Options))
		for _, option := range nextStatement.ShortcutOptionStatement.Options {
			options = append(options, dr.textElementsToString(option.LineStatement.Text.Elements))
		}
		return &DialogueElement{
			Options: options,
		}, true
	case nextStatement.SetStatement != nil:
		value, err := evaluateExpression(nextStatement.SetStatement.Expression, dr.variableStorer)
		if err != nil {
			panic(err) // FIXME: better handle errors here
		}
		switch { // FIXME: take operator into account to support += and the likes
		case value.IsNumber():
			dr.variableStorer.SetNumberValue(nextStatement.SetStatement.VariableID, *value.Number)
		case value.IsBoolean():
			dr.variableStorer.SetBooleanValue(nextStatement.SetStatement.VariableID, *value.Boolean)
		case value.IsString():
			dr.variableStorer.SetStringValue(nextStatement.SetStatement.VariableID, *value.String)
		}
		return dr.Next(choice)
	}

	return nil, false // TODO: we should never get there
}

func NewDialogueRunner(dialogue *tree.Dialogue, storer VariableStorer) *DialogueRunner {
	statementsToRun := container.Stack[*StatementQueue]{}
	statementsToRun.Push(&StatementQueue{statements: dialogue.Nodes[0].Statements})

	if storer == nil {
		storer = NewInMemoryVariableStorer()
	}

	return &DialogueRunner{
		dialogue:        dialogue,
		statementsToRun: statementsToRun,
		variableStorer:  storer,
	}
}

type StatementQueue struct {
	statements []*tree.Statement
	pointer    int
}

func (sq *StatementQueue) NextStatement() (*tree.Statement, bool) {
	if sq.pointer >= len(sq.statements) {
		return nil, false
	}

	next := sq.statements[sq.pointer]
	sq.pointer++
	return next, true
}
