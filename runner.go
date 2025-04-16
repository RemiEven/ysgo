// Package ysgo defines a DialogueRunner type able to execute dialogues.
package ysgo

import (
	"errors"
	"fmt"
	"io"
	"maps"
	"math"
	"slices"
	"strings"

	"github.com/remieven/ysgo/internal/container"
	"github.com/remieven/ysgo/internal/rng"
	"github.com/remieven/ysgo/internal/tree"
	"github.com/remieven/ysgo/markup"
	"github.com/remieven/ysgo/variable"
)

// DialogueRunner is able to execute a YarnSpinner dialogue.
// It keeps track of the current state of the dialogue (variables, current and visited steps).
// It orchestrates the call of Commands and Functions.
type DialogueRunner struct {
	dialogue        *tree.Dialogue
	statementsToRun container.Stack[*container.Stack[*statementQueue]]
	lastStatement   *tree.Statement
	variableStorer  variable.Storer
	functionStorer  *functionStorer
	commandStorer   *commandStorer
	commandErrChan  <-chan error
	lineParser      markup.LineParser

	saliencyStrategies      map[string]ContentSaliencyStrategy
	currentSaliencyStrategy ContentSaliencyStrategy

	currentNodes         container.Stack[string]
	visitedNodes         map[string]int
	variableSnapshot     map[string]variable.Value
	visitedNodesSnapshot map[string]int
}

// DialogueElement represents a step of a dialogue as it is presented in a game.
// Either Line or Options holds a value.
// Node holds the name of the dialogue node that contains the element.
type DialogueElement struct {
	Node    string
	Line    *Line
	Options []DialogueOption
}

// DialogueOption holds the data about one possible choice the player is presented with.
type DialogueOption struct {
	Line     *Line
	Disabled bool
}

// Line contains an interpolated text line along with its tags, if any.
type Line struct {
	markup.ParseResult
	Tags []string
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

// Next advances the dialogue to the next step.
// If the previous step was a choice, then the selected option index should be given as an argument,
// else that argument is ignored.
// If an ongoing command (eg. <<wait 3>>) from the dialogue is still running, ErrWaitingForCommandCompletion
// is returned.
// Else, if no other error is encountered, the next DialogueElement to display is returned.
// If the Dialogue has ended, then both return values will be nil.
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
			statementsToRun := dr.statementsToRun.Peek()
			statementsToRun.Push(&statementQueue{
				statements: statements,
			})
		}
	}

	if dr.statementsToRun.Size() > 0 && dr.statementsToRun.Peek().Size() == 0 {
		dr.incrementNodeTrackingIfAllowed(dr.currentNodes.Peek())
		dr.currentNodes.Pop()
		dr.statementsToRun.Pop()
	}

	if dr.statementsToRun.Size() == 0 {
		return nil, nil
	}

	statementsToRun := dr.statementsToRun.Peek()

	nextStatement, ok := statementsToRun.Peek().nextStatement()
	if !ok {
		statementsToRun.Pop()
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
			Node: dr.currentNodes.Peek(),
			Line: &Line{
				ParseResult: *markupResult,
				Tags:        nextStatement.LineStatement.Tags,
			},
		}, nil
	case nextStatement.ShortcutOptionStatement != nil:
		options := make([]DialogueOption, 0, len(nextStatement.ShortcutOptionStatement.Options))
		for i, option := range nextStatement.ShortcutOptionStatement.Options {
			markupResult, err := dr.textElementsToMarkup(option.LineStatement.Text.Elements)
			if err != nil {
				return nil, fmt.Errorf("failed to prepare option %v: %w", i, err)
			}
			disabled := false
			if condition := option.LineStatement.Condition; condition != nil {
				enabled, err := evaluateExpression(option.LineStatement.Condition, dr.variableStorer, dr.functionStorer)
				if err != nil {
					return nil, fmt.Errorf("failed to evaluate line condition: %w", err)
				} else if enabled.Boolean == nil {
					return nil, fmt.Errorf("encountered non boolean line condition")
				}
				disabled = !*enabled.Boolean
			}
			options = append(options, DialogueOption{
				Line: &Line{
					ParseResult: *markupResult,
					Tags:        option.LineStatement.Tags,
				},
				Disabled: disabled,
			})
		}
		return &DialogueElement{
			Node:    dr.currentNodes.Peek(),
			Options: options,
		}, nil
	case nextStatement.LineGroupStatement != nil:
		if err := dr.executeLineGroupStatement(nextStatement.LineGroupStatement); err != nil {
			return nil, fmt.Errorf("failed to execute line group statement: %w", err)
		}
		return dr.Next(choice)
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
	case nextStatement.ReturnStatement != nil:
		dr.incrementNodeTrackingIfAllowed(dr.currentNodes.Peek())
		dr.currentNodes.Pop()
		dr.statementsToRun.Pop()
		return dr.Next(choice)
	}

	return nil, errors.New("encountered an unsupported type of statement")
}

// NewDialogueRunner creates a new runner working on the dialogue that can be
// parsed from the given readers.
// The storer argument is optional and if provided, allows to store variables
// elsewhere than in the default in-memory store.
// The given rngSeed serves to create deterministic random values. It will be
// used when eg. the dice(6) function is called in a dialogue.
func NewDialogueRunner(storer variable.Storer, rngSeed string, readers ...io.Reader) (*DialogueRunner, error) {
	dialogue, err := tree.FromReader(io.MultiReader(readers...))
	if err != nil {
		return nil, fmt.Errorf("failed to create dialogue: %w", err)
	}

	return newDialogueRunner(storer, rngSeed, dialogue)
}

// NewDialogueRunnerFromMarshaled creates a new runner working on the previously
// parsed and marshaled dialogue in the given reader.
// The storer argument is optional and if provided, allows to store variables
// elsewhere than in the default in-memory store.
// The given rngSeed serves to create deterministic random values. It will be
// used when eg. the dice(6) function is called in a dialogue.
func NewDialogueRunnerFromMarshaled(storer variable.Storer, rngSeed string, reader io.Reader) (*DialogueRunner, error) {
	dialogue, err := tree.FromMarshaledReader(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to create dialogue: %w", err)
	}

	return newDialogueRunner(storer, rngSeed, dialogue)
}

func newDialogueRunner(storer variable.Storer, rngSeed string, dialogue *tree.Dialogue) (*DialogueRunner, error) {
	statementsToRun := &container.Stack[*statementQueue]{}
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
		statementsToRun: container.Stack[*container.Stack[*statementQueue]]{statementsToRun},
		variableStorer:  storer,
		commandStorer:   newCommandStorer(),
		visitedNodes:    map[string]int{},
		currentNodes:    container.Stack[string]{firstNode.Title()},

		saliencyStrategies: map[string]ContentSaliencyStrategy{
			FirstSaliencyStrategyName:                         &FirstSaliencyStrategy{},
			BestSaliencyStrategyName:                          &BestSaliencyStrategy{},
			BestLeastRecentlyViewedSaliencyStrategyName:       &BestLeastRecentlyViewedSaliencyStrategy{},
			RandomBestLeastRecentlyViewedSaliencyStrategyName: &RandomBestLeastRecentlyViewedSaliencyStrategy{},
		},
	}
	runner.currentSaliencyStrategy = runner.saliencyStrategies[FirstSaliencyStrategyName]

	functionStorer := newFunctionStorer(rng)
	functionStorer.convertAndAddFunction("visited", func(node string) bool {
		_, ok := runner.visitedNodes[node]
		return ok
	})
	functionStorer.convertAndAddFunction("visited_count", func(node string) int {
		count := runner.visitedNodes[node]
		return count
	})

	runner.functionStorer = functionStorer

	return runner, nil
}

func (dr *DialogueRunner) executeLineGroupStatement(statement *tree.LineGroupStatement) error {
	// find items that can be selected
	items := make([]*tree.LineGroupItem, 0, len(statement.Items))
	for _, item := range statement.Items {
		if condition := item.LineStatement.Condition; condition != nil {
			possible, err := evaluateExpression(condition, dr.variableStorer, dr.functionStorer)
			if err != nil {
				return fmt.Errorf("failed to evaluate item condition: %w", err)
			} else if possible.Boolean == nil {
				return fmt.Errorf("encountered non boolean item condition")
			} else if !*possible.Boolean {
				continue // this item cannot be run now
			}
		}
		items = append(items, item)
	}

	// turn those items into options for the saliency strategy
	itemsByOption := make(map[*ContentSaliencyOption]*tree.LineGroupItem, len(items))
	for _, item := range items {
		contentID := item.LineStatement.LineId()
		if contentID == "" {
			return fmt.Errorf("encountered a line without an ID in a line group")
		}

		viewCount := 0
		viewCountVariableName := viewCountVariableNameForContent(contentID)
		if viewCountVariable, ok := dr.variableStorer.GetValue(viewCountVariableName); ok {
			if viewCountVariable.Number == nil {
				return fmt.Errorf("variable [%s] should be a number", viewCountVariableName)
			}
			viewCount = (int)(*viewCountVariable.Number)
		}

		itemsByOption[&ContentSaliencyOption{
			ContentID:  contentID,
			Complexity: item.LineStatement.Condition.ComplexityScore(),
			ViewCount:  viewCount,
		}] = item
	}
	options := make([]*ContentSaliencyOption, 0, len(itemsByOption))
	options = slices.AppendSeq(options, maps.Keys(itemsByOption))

	// select the most salient option, and update state accordingly
	selectedOption, ok := dr.currentSaliencyStrategy.QueryBestContent(options)
	if !ok {
		// no items can be run, continue the dialogue
		return nil
	}
	dr.currentSaliencyStrategy.ContentWasSelected(selectedOption)
	dr.variableStorer.SetNumberValue(viewCountVariableNameForContent(selectedOption.ContentID), float64(selectedOption.ViewCount+1))

	// stack the statements of the selected item on top of the queue
	selectedItem := itemsByOption[selectedOption]
	statementsToRun := dr.statementsToRun.Peek()
	statementsToAdd := make([]*tree.Statement, 1, 1+len(selectedItem.Statements))
	statementsToAdd[0] = &tree.Statement{
		LineStatement: selectedItem.LineStatement,
	}
	statementsToAdd = append(statementsToAdd, selectedItem.Statements...)
	statementsToRun.Push(&statementQueue{statements: statementsToAdd})
	return nil
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
	} else if statement.Detour {
		dr.statementsToRun.Push(&container.Stack[*statementQueue]{&statementQueue{statements: node.Statements}})
		dr.currentNodes.Push(node.Title())
	} else {
		for _, node := range dr.currentNodes {
			dr.incrementNodeTrackingIfAllowed(node)
		}
		dr.variableSnapshot = dr.variableStorer.GetValues()
		dr.visitedNodesSnapshot = maps.Clone(dr.visitedNodes)
		dr.statementsToRun.Clear()
		dr.statementsToRun.Push(&container.Stack[*statementQueue]{&statementQueue{statements: node.Statements}})
		dr.currentNodes.Clear()
		dr.currentNodes.Push(node.Title())
	}
	return nil
}

func (dr *DialogueRunner) incrementNodeTrackingIfAllowed(nodeName string) {
	node, ok := dr.dialogue.FindNode(nodeName)
	if ok && node.Headers != nil && node.Headers["tracking"] != "never" {
		dr.visitedNodes[nodeName]++
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
			statementsToRun := dr.statementsToRun.Peek()
			statementsToRun.Push(&statementQueue{statements: clause.Statements})
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

	errChan := dr.commandStorer.call(commandName, values[1:])
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

	if _, err := dr.functionStorer.call(statement.FunctionID, values); err != nil {
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

// RestoreAt uses a snapshot to restore a dialogue runner to a former state.
func (dr *DialogueRunner) RestoreAt(snapshot *Snapshot) error {
	node, ok := dr.dialogue.FindNode(snapshot.CurrentNode)
	if !ok {
		return fmt.Errorf("dialogue does not contain a node with title [%s]", snapshot.CurrentNode)
	}

	dr.visitedNodes = snapshot.VisitedNodes
	dr.variableStorer.Clear()
	for variable, value := range snapshot.Variables {
		if value.Boolean != nil {
			dr.variableStorer.SetBooleanValue(variable, *value.Boolean)
		}
		if value.Number != nil {
			dr.variableStorer.SetNumberValue(variable, *value.Number)
		}
		if value.String != nil {
			dr.variableStorer.SetStringValue(variable, *value.String)
		}
	}

	dr.statementsToRun.Clear()
	dr.statementsToRun.Push(&container.Stack[*statementQueue]{&statementQueue{statements: node.Statements}})
	dr.currentNodes.Clear()
	dr.currentNodes.Push(node.Title())
	return nil
}

// AddFunction adds a custom function to the library of functions that can be called from a dialogue.
func (dr *DialogueRunner) AddFunction(functionID string, function YarnSpinnerFunction) {
	dr.functionStorer.addFunction(functionID, function)
}

// ConvertAndAddFunction is a convenience wrapper around AddFunction so that manual conversion to YarnSpinnerFunction isn't needed.
// Refer to the unit tests of functionStorer to learn more about the limitations on the accepted functions.
func (dr *DialogueRunner) ConvertAndAddFunction(functionID string, function any) error {
	return dr.functionStorer.convertAndAddFunction(functionID, function)
}

// AddCommand adds a custom command to the library of commands that can be called from a dialogue.
func (dr *DialogueRunner) AddCommand(commandID string, command YarnSpinnerCommand) {
	dr.commandStorer.addCommand(commandID, command)
}

// ConvertAndAddCommand is a convenience wrapper around AddCommand so that manual conversion to YarnSpinnerCommand isn't needed.
// Refer to the unit tests of commandStorer to learn more about the limitations on the accepted functions.
func (dr *DialogueRunner) ConvertAndAddCommand(commandID string, command any) error {
	return dr.commandStorer.convertAndAddCommand(commandID, command)
}

// AddSaliencyStrategy adds a custom saliency strategy, which can be used later to select content.
func (dr *DialogueRunner) AddSaliencyStrategy(name string, strategy ContentSaliencyStrategy) {
	dr.saliencyStrategies[name] = strategy
}

// SetSaliencyStrategy sets the saliency strategy the dialogue runner should use when selecting content, or returns an error.
func (dr *DialogueRunner) SetSaliencyStrategy(name string) error {
	strategy, ok := dr.saliencyStrategies[name]
	if !ok {
		return fmt.Errorf("unknown saliency strategy [%s]", name)
	}
	dr.currentSaliencyStrategy = strategy
	return nil
}

// Snapshot returns the state of the dialogue runner as of the last time a node was entered.
// It can then be used to later restore the state of the dialogue runner.
func (dr *DialogueRunner) Snapshot() *Snapshot {
	return &Snapshot{
		Variables:    dr.variableSnapshot,
		CurrentNode:  dr.currentNodes.Peek(),
		VisitedNodes: dr.visitedNodes,
	}
}

type statementQueue struct {
	statements []*tree.Statement
	pointer    int
}

func (sq *statementQueue) nextStatement() (*tree.Statement, bool) {
	if sq.pointer >= len(sq.statements) {
		return nil, false
	}

	next := sq.statements[sq.pointer]
	sq.pointer++
	return next, true
}

type errWaitingForCommandCompletion string

// Error returns the error message associated with e.
func (e errWaitingForCommandCompletion) Error() string {
	return string(e)
}

// ErrWaitingForCommandCompletion is returned when DialogueRunner#Next(...) is called
// when an ongoing command hasn't ended yet.
const ErrWaitingForCommandCompletion errWaitingForCommandCompletion = "waiting for command completion"
