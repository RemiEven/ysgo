package runner

import (
	"errors"
	"fmt"
	"math"
	"strings"

	"github.com/RemiEven/ysgo/container"
	"github.com/RemiEven/ysgo/markup"
	"github.com/RemiEven/ysgo/runner/rng"
	"github.com/RemiEven/ysgo/tree"
)

type DialogueRunner struct {
	dialogue        *tree.Dialogue
	statementsToRun container.Stack[*StatementQueue] // TODO: would the name "next steps" be better here?
	lastStatement   *tree.Statement
	variableStorer  VariableStorer
	functionStorer  *FunctionStorer
	lineParser      markup.LineParser
}

type DialogueElement struct { // dialogueStep? dialogueElement?
	Line    *markup.ParseResult
	Options []DialogueOption
}

type DialogueOption struct {
	Line     *markup.ParseResult
	Disabled bool
}

func (dr *DialogueRunner) textElementsToMarkup(elements []*tree.LineFormattedTextElement) (*markup.ParseResult, error) {
	var builder strings.Builder
	for i := range elements {
		if elements[i].Text != "" {
			builder.WriteString(elements[i].Text)
		} else if expression := elements[i].Expression; expression != nil {
			value, err := evaluateExpression(expression, dr.variableStorer, dr.functionStorer)
			if err != nil {
				return nil, fmt.Errorf("failed to evaluate expression: %w", err)
			}
			builder.WriteString(value.ToString())
		}
	}
	markupResult, err := dr.lineParser.ParseMarkup(builder.String())
	if err != nil {
		return nil, fmt.Errorf("failed to parse markup: %w", err)
	}

	return markupResult, nil
}

func (dr *DialogueRunner) isWaitingForChoice() bool {
	return dr.lastStatement != nil && dr.lastStatement.ShortcutOptionStatement != nil
}

func (dr *DialogueRunner) Next(choice int) (*DialogueElement, bool, error) { // TODO: have nicer arguments there? like an input struct maybe?
	if dr.isWaitingForChoice() {
		dr.statementsToRun.Push(&StatementQueue{
			statements: dr.lastStatement.ShortcutOptionStatement.Options[choice].Statements,
		})
	}

	if dr.statementsToRun.Size() == 0 {
		return nil, false, nil
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
		markupResult, err := dr.textElementsToMarkup(nextStatement.LineStatement.Text.Elements)
		if err != nil {
			return nil, false, fmt.Errorf("failed to prepare line: %w", err)
		}
		return &DialogueElement{
			Line: markupResult,
		}, true, nil
	case nextStatement.ShortcutOptionStatement != nil:
		options := make([]DialogueOption, 0, len(nextStatement.ShortcutOptionStatement.Options))
		for i, option := range nextStatement.ShortcutOptionStatement.Options {
			markupResult, err := dr.textElementsToMarkup(option.LineStatement.Text.Elements)
			if err != nil {
				return nil, false, fmt.Errorf("failed to prepare option %v: %w", i, err)
			}
			disabled := false
			if option.LineStatement.Condition != nil {
				enabled, err := evaluateExpression(option.LineStatement.Condition, dr.variableStorer, dr.functionStorer)
				if err != nil {
					return nil, false, fmt.Errorf("failed to evaluate line condition: %w", err)
				} else if enabled.Boolean == nil {
					return nil, false, fmt.Errorf("encountered non boolean line condition")
				}
				disabled = !*enabled.Boolean
			}
			options = append(options, DialogueOption{
				Line:     markupResult,
				Disabled: disabled,
			})
		}
		return &DialogueElement{
			Options: options,
		}, true, nil
	case nextStatement.SetStatement != nil:
		if err := dr.executeSetStatement(nextStatement.SetStatement); err != nil {
			return nil, false, fmt.Errorf("failed to execute set statement: %w", err)
		}
		return dr.Next(choice)
	case nextStatement.JumpStatement != nil:
		if err := dr.executeJumpStatement(nextStatement.JumpStatement); err != nil {
			return nil, false, fmt.Errorf("failed to execute jump statement: %w", err)
		}
		return dr.Next(choice)
	case nextStatement.IfStatement != nil:
		if err := dr.executeIfStatement(nextStatement.IfStatement); err != nil {
			return nil, false, fmt.Errorf("failed to execute if statement: %w", err)
		}
		return dr.Next(choice)
	}

	return nil, false, errors.New("encountered an unsupported type of statement")
}

func NewDialogueRunner(dialogue *tree.Dialogue, storer VariableStorer, rngSeed string) (*DialogueRunner, error) {
	statementsToRun := container.Stack[*StatementQueue]{}
	statementsToRun.Push(&StatementQueue{statements: dialogue.Nodes[0].Statements})

	if storer == nil {
		storer = NewInMemoryVariableStorer()
	}

	rng, err := rng.NewRNG(rngSeed)
	if err != nil {
		return nil, fmt.Errorf("failed to create rng: %w", err)
	}

	return &DialogueRunner{
		dialogue:        dialogue,
		statementsToRun: statementsToRun,
		variableStorer:  storer,
		functionStorer:  newFunctionStorer(rng),
	}, nil
}

func (dr *DialogueRunner) executeSetStatement(statement *tree.SetStatement) error {
	value, err := evaluateExpression(statement.Expression, dr.variableStorer, dr.functionStorer)
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

func (dr *DialogueRunner) executeJumpStatement(statement *tree.JumpStatement) error {
	value, err := evaluateExpression(statement.Expression, dr.variableStorer, dr.functionStorer)
	if err != nil {
		return fmt.Errorf("failed to evaluate expression: %w", err)
	} else if value.String == nil {
		return fmt.Errorf("destination must be a string")
	} else if node, ok := dr.dialogue.FindNode(*value.String); !ok {
		return fmt.Errorf("node [%s] not found in dialogue", *value.String)
	} else {
		dr.statementsToRun.Push(&StatementQueue{statements: node.Statements})
	}
	return nil
}

func (dr *DialogueRunner) executeIfStatement(statement *tree.IfStatement) error {
	for _, clause := range statement.Clauses {
		condition, err := evaluateExpression(clause.Condition, dr.variableStorer, dr.functionStorer)
		if err != nil {
			return fmt.Errorf("failed to evaluate condition: %w", err)
		} else if !condition.IsBoolean() {
			return fmt.Errorf("condition must be a boolean")
		}
		if *condition.Boolean {
			dr.statementsToRun.Push(&StatementQueue{statements: clause.Statements})
			return nil
		}
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
