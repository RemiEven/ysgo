package runner

import (
	"errors"
	"fmt"
	"math"
	"strings"

	"github.com/RemiEven/ysgo/internal/container"
	"github.com/RemiEven/ysgo/internal/tree"
	"github.com/RemiEven/ysgo/markup"
	"github.com/RemiEven/ysgo/runner/rng"
	"github.com/RemiEven/ysgo/variable"
)

type DialogueRunner struct {
	dialogue        *tree.Dialogue
	statementsToRun container.Stack[*statementQueue]
	lastStatement   *tree.Statement
	variableStorer  variable.Storer
	functionStorer  *functionStorer
	commandStorer   *commandStorer
	commandErrChan  <-chan error
	lineParser      markup.LineParser
	currentNode     string
	visitedNodes    map[string]int
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

func (dr *DialogueRunner) Next(choice int) (*DialogueElement, error) {
	if dr.commandErrChan != nil {
		select {
		case err := <-dr.commandErrChan:
			dr.commandErrChan = nil
			if err != nil {
				return nil, fmt.Errorf("failed to execute ongoing command: %w", err)
			}
		default:
			return nil, ErrWaitingForCommandCompletion
		}
	}

	if dr.isWaitingForChoice() {
		if statements := dr.lastStatement.ShortcutOptionStatement.Options[choice].Statements; len(statements) != 0 {
			dr.statementsToRun.Push(&statementQueue{
				statements: statements,
			})
		}
	}

	if dr.statementsToRun.Size() == 0 {
		return nil, nil
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
			return nil, fmt.Errorf("failed to prepare line: %w", err)
		}
		return &DialogueElement{
			Line: markupResult,
		}, nil
	case nextStatement.ShortcutOptionStatement != nil:
		options := make([]DialogueOption, 0, len(nextStatement.ShortcutOptionStatement.Options))
		for i, option := range nextStatement.ShortcutOptionStatement.Options {
			markupResult, err := dr.textElementsToMarkup(option.LineStatement.Text.Elements)
			if err != nil {
				return nil, fmt.Errorf("failed to prepare option %v: %w", i, err)
			}
			disabled := false
			if option.LineStatement.Condition != nil {
				enabled, err := evaluateExpression(option.LineStatement.Condition, dr.variableStorer, dr.functionStorer)
				if err != nil {
					return nil, fmt.Errorf("failed to evaluate line condition: %w", err)
				} else if enabled.Boolean == nil {
					return nil, fmt.Errorf("encountered non boolean line condition")
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
		}, nil
	case nextStatement.SetStatement != nil:
		if err := dr.executeSetStatement(nextStatement.SetStatement); err != nil {
			return nil, fmt.Errorf("failed to execute set statement: %w", err)
		}
		return dr.Next(choice)
	case nextStatement.JumpStatement != nil:
		if err := dr.executeJumpStatement(nextStatement.JumpStatement); err != nil {
			return nil, fmt.Errorf("failed to execute jump statement: %w", err)
		}
		return dr.Next(choice)
	case nextStatement.IfStatement != nil:
		if err := dr.executeIfStatement(nextStatement.IfStatement); err != nil {
			return nil, fmt.Errorf("failed to execute if statement: %w", err)
		}
		return dr.Next(choice)
	case nextStatement.CommandStatement != nil:
		if stop, err := dr.executeCommandStatement(nextStatement.CommandStatement); err != nil {
			return nil, fmt.Errorf("failed to execute command statement: %w", err)
		} else if stop {
			return nil, nil
		} else if dr.commandErrChan != nil {
			return nil, ErrWaitingForCommandCompletion
		}
		return dr.Next(choice)
	case nextStatement.CallStatement != nil:
		if err := dr.executeCallStatement(nextStatement.CallStatement); err != nil {
			return nil, fmt.Errorf("failed to execute call statement: %w", err)
		}
		return dr.Next(choice)
	case nextStatement.DeclareStatement != nil:
		if err := dr.executeDeclareStatement(nextStatement.DeclareStatement); err != nil {
			return nil, fmt.Errorf("failed to execute declare statement: %w", err)
		}
		return dr.Next(choice)
	}

	return nil, errors.New("encountered an unsupported type of statement")
}

func NewDialogueRunner(dialogue *tree.Dialogue, storer variable.Storer, rngSeed string) (*DialogueRunner, error) {
	statementsToRun := container.Stack[*statementQueue]{}
	firstNode := dialogue.Nodes[0]
	statementsToRun.Push(&statementQueue{statements: firstNode.Statements})

	if storer == nil {
		storer = variable.NewInMemoryStorer()
	}

	rng, err := rng.NewRNG(rngSeed)
	if err != nil {
		return nil, fmt.Errorf("failed to create rng: %w", err)
	}

	runner := &DialogueRunner{
		dialogue:        dialogue,
		statementsToRun: statementsToRun,
		variableStorer:  storer,
		commandStorer:   newCommandStorer(),
		visitedNodes:    map[string]int{},
		currentNode:     firstNode.Title(),
	}

	functionStorer := newFunctionStorer(rng)
	functionStorer.ConvertAndAddFunction("visited", func(node string) bool {
		_, ok := runner.visitedNodes[node]
		return ok
	})
	functionStorer.ConvertAndAddFunction("visited_count", func(node string) int {
		count := runner.visitedNodes[node]
		return count
	})

	runner.functionStorer = functionStorer

	return runner, nil
}

func (dr *DialogueRunner) executeSetStatement(statement *tree.SetStatement) error {
	value, err := evaluateExpression(statement.Expression, dr.variableStorer, dr.functionStorer)
	if err != nil {
		return fmt.Errorf("failed to evaluate expression: %w", err)
	}
	previousValue, ok := dr.variableStorer.GetValue(statement.VariableID)

	if ok {
		bothValuesAreNumbers := value.Number != nil && previousValue.Number != nil
		bothValuesAreBooleans := value.Boolean != nil && previousValue.Boolean != nil
		bothValuesAreStrings := value.String != nil && previousValue.String != nil
		if !(bothValuesAreNumbers || bothValuesAreBooleans || bothValuesAreStrings) {
			return fmt.Errorf("variable [%s] type cannot be changed", statement.VariableID)
		}
	} else if statement.InPlaceOperator != tree.AssignmentInPlaceOperator {
		return fmt.Errorf("variable [%s] not found in storage, only assignment is allowed", statement.VariableID)
	}

	switch {
	case value.Number != nil:
		var newNumberValue float64
		numberValue := *value.Number
		switch statement.InPlaceOperator {
		case tree.AssignmentInPlaceOperator:
			newNumberValue = numberValue
		case tree.MultiplicationInPlaceOperator:
			newNumberValue = (*previousValue.Number) * (numberValue)
		case tree.DivisionInPlaceOperator:
			newNumberValue = (*previousValue.Number) / (numberValue)
		case tree.ModuloInPlaceOperator:
			newNumberValue = math.Mod(*previousValue.Number, numberValue)
		case tree.AdditionInPlaceOperator:
			newNumberValue = (*previousValue.Number) + (numberValue)
		case tree.SubtractionInPlaceOperator:
			newNumberValue = (*previousValue.Number) - (numberValue)
		default:
			return fmt.Errorf("unknown assignment operator encountered")
		}
		dr.variableStorer.SetNumberValue(statement.VariableID, newNumberValue)
	case value.Boolean != nil:
		if statement.InPlaceOperator != tree.AssignmentInPlaceOperator {
			return fmt.Errorf("unsupported assignment operator for boolean variable encountered")
		}
		dr.variableStorer.SetBooleanValue(statement.VariableID, *value.Boolean)
	case value.String != nil:
		var newStringValue string
		stringValue := *value.String
		switch statement.InPlaceOperator {
		case tree.AssignmentInPlaceOperator:
			newStringValue = stringValue
		case tree.AdditionInPlaceOperator:
			newStringValue = (stringValue) + (*previousValue.String)
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
		dr.incrementNodeTrackingIfAllowed()
		dr.statementsToRun.Clear()
		dr.statementsToRun.Push(&statementQueue{statements: node.Statements})
		dr.currentNode = node.Title()
	}
	return nil
}

func (dr *DialogueRunner) incrementNodeTrackingIfAllowed() {
	node, ok := dr.dialogue.FindNode(dr.currentNode)
	if ok && node.Headers != nil && node.Headers["tracking"] != "never" {
		dr.visitedNodes[dr.currentNode]++
	}
}

func (dr *DialogueRunner) executeIfStatement(statement *tree.IfStatement) error {
	for _, clause := range statement.Clauses {
		condition, err := evaluateExpression(clause.Condition, dr.variableStorer, dr.functionStorer)
		if err != nil {
			return fmt.Errorf("failed to evaluate condition: %w", err)
		} else if condition.Boolean == nil {
			return fmt.Errorf("condition must be a boolean")
		}
		if *condition.Boolean {
			dr.statementsToRun.Push(&statementQueue{statements: clause.Statements})
			return nil
		}
	}
	return nil
}

func (dr *DialogueRunner) executeCommandStatement(statement *tree.CommandStatement) (stop bool, err error) {
	if len(statement.Elements) == 0 {
		return false, fmt.Errorf("missing command name")
	}
	values := make([]*variable.Value, 0, len(statement.Elements))
	for i := range statement.Elements {
		value, err := evaluateExpression(statement.Elements[i].Expression, dr.variableStorer, dr.functionStorer)
		if err != nil {
			return false, fmt.Errorf("failed to evaluate element [%v] of command: %w", i, err)
		}
		values = append(values, value)
	}

	if values[0].String == nil {
		return false, fmt.Errorf("first element of command must be a string")
	}

	commandName := *values[0].String

	if commandName == "stop" {
		return true, nil
	}

	errChan := dr.commandStorer.Call(commandName, values[1:])
	select {
	case err := <-errChan:
		if err != nil {
			return false, fmt.Errorf("failed to execute command [%s]: %w", commandName, err)
		}
		return false, nil
	default:
		dr.commandErrChan = errChan
		return false, nil
	}
}

func (dr *DialogueRunner) executeCallStatement(statement *tree.CallStatement) error {
	values := make([]*variable.Value, 0, len(statement.Arguments))
	for i := range statement.Arguments {
		value, err := evaluateExpression(statement.Arguments[i], dr.variableStorer, dr.functionStorer)
		if err != nil {
			return fmt.Errorf("failed to evaluate argument [%v] of call statement: %w", i, err)
		}
		values = append(values, value)
	}

	if _, err := dr.functionStorer.Call(statement.FunctionID, values); err != nil {
		return fmt.Errorf("failed to execute call statement for function [%s]: %w", statement.FunctionID, err)
	}

	return nil
}

func (dr *DialogueRunner) executeDeclareStatement(statement *tree.DeclareStatement) error {
	return dr.executeSetStatement(&tree.SetStatement{
		VariableID:      statement.VariableID,
		InPlaceOperator: tree.AssignmentInPlaceOperator,
		Expression:      statement.Value,
	})
}

func (dr *DialogueRunner) AddFunction(functionID string, function YarnSpinnerFunction) {
	dr.functionStorer.AddFunction(functionID, function)
}

func (dr *DialogueRunner) ConvertAndAddFunction(functionID string, function any) error {
	return dr.functionStorer.ConvertAndAddFunction(functionID, function)
}

func (dr *DialogueRunner) AddCommand(commandID string, command YarnSpinnerCommand) {
	dr.commandStorer.AddCommand(commandID, command)
}

func (dr *DialogueRunner) ConvertAndAddCommand(commandID string, command any) error {
	return dr.commandStorer.ConvertAndAddCommand(commandID, command)
}

type statementQueue struct {
	statements []*tree.Statement
	pointer    int
}

func (sq *statementQueue) NextStatement() (*tree.Statement, bool) {
	if sq.pointer >= len(sq.statements) {
		return nil, false
	}

	next := sq.statements[sq.pointer]
	sq.pointer++
	return next, true
}

type errWaitingForCommandCompletion string

func (e errWaitingForCommandCompletion) Error() string {
	return string(e)
}

const ErrWaitingForCommandCompletion errWaitingForCommandCompletion = "waiting for command completion"
