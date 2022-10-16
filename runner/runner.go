package runner

import (
	"fmt"
	"math"
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
		if err := dr.executeSetStatement(nextStatement.SetStatement); err != nil {
			panic(err) // FIXME: better handle errors here
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

func (dr *DialogueRunner) executeSetStatement(statement *tree.SetStatement) error {
	value, err := evaluateExpression(statement.Expression, dr.variableStorer)
	if err != nil {
		return fmt.Errorf("failed to evaluate expression: %w", err)
	}
	previousValue, ok := dr.variableStorer.GetValue(statement.VariableID)

	if ok {
		bothValuesAreNumbers := value.IsNumber() && previousValue.IsNumber()
		bothValuesAreBooleans := value.IsBoolean() && previousValue.IsBoolean()
		bothValuesAreStrings := value.IsString() && previousValue.IsString()
		if !(bothValuesAreNumbers || bothValuesAreBooleans || bothValuesAreStrings) {
			return fmt.Errorf("variable [%s] type cannot be changed", statement.VariableID)
		}
	} else if statement.InPlaceOperator != tree.AssignmentInPlaceOperator {
		return fmt.Errorf("variable [%s] not found in storage, only assignment is allowed", statement.VariableID)
	}

	switch {
	case value.IsNumber():
		var newNumberValue float64
		switch statement.InPlaceOperator {
		case tree.AssignmentInPlaceOperator:
			newNumberValue = *value.Number
		case tree.MultiplicationInPlaceOperator:
			newNumberValue = (*value.Number) * (*previousValue.Number)
		case tree.DivisionInPlaceOperator:
			newNumberValue = (*value.Number) / (*previousValue.Number)
		case tree.ModuloInPlaceOperator:
			newNumberValue = math.Mod(*value.Number, *previousValue.Number)
		case tree.AdditionInPlaceOperator:
			newNumberValue = (*value.Number) + (*previousValue.Number)
		case tree.SubtractionInPlaceOperator:
			newNumberValue = (*value.Number) - (*previousValue.Number)
		default:
			return fmt.Errorf("unknown assignment operator encountered")
		}
		dr.variableStorer.SetNumberValue(statement.VariableID, newNumberValue)
	case value.IsBoolean():
		if statement.InPlaceOperator != tree.AssignmentInPlaceOperator {
			return fmt.Errorf("unsupported assignment operator for boolean variable encountered")
		}
		dr.variableStorer.SetBooleanValue(statement.VariableID, *value.Boolean)
	case value.IsString():
		var newStringValue string
		switch statement.InPlaceOperator {
		case tree.AssignmentInPlaceOperator:
			newStringValue = *value.String
		case tree.AdditionInPlaceOperator:
			newStringValue = (*value.String) + (*previousValue.String)
		default:
			return fmt.Errorf("unsupported assignment operator for string variable encountered")
		}
		dr.variableStorer.SetStringValue(statement.VariableID, newStringValue)
	}

	return nil
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
