package tree

import (
	"strconv"
	"strings"

	"github.com/remieven/ysgo/variable"
)

type Dialogue struct {
	Nodes []Node
}

func (d *Dialogue) FindNode(title string) (*Node, bool) {
	for _, node := range d.Nodes {
		if node.Title() == title {
			return &node, true
		}
	}
	return nil, false
}

type Node struct {
	Headers    map[string]string
	Statements []*Statement
}

func (n *Node) Title() string {
	return n.Headers["title"]
}

type Statement struct {
	LineStatement           *LineStatement
	ShortcutOptionStatement *ShortcutOptionStatement
	SetStatement            *SetStatement
	JumpStatement           *JumpStatement
	IfStatement             *IfStatement
	CommandStatement        *CommandStatement
	CallStatement           *CallStatement
	DeclareStatement        *DeclareStatement
	ReturnStatement         *ReturnStatement
}

type LineStatement struct {
	Text      *LineFormattedText
	Condition *Expression
	Tags      []string
}

type LineFormattedText struct {
	Elements []*LineFormattedTextElement
}

type LineFormattedTextElement struct {
	Text       string
	Expression *Expression
}

type ShortcutOptionStatement struct {
	Options []*ShortcutOption
}

type ShortcutOption struct {
	LineStatement *LineStatement
	Statements    []*Statement
}

type SetStatement struct {
	VariableID      string
	InPlaceOperator int
	Expression      *Expression
}

type JumpStatement struct {
	Expression *Expression
	Detour     bool
}

type IfStatement struct {
	Clauses []*Clause
}

type Clause struct {
	Condition  *Expression
	Statements []*Statement
}

type CommandStatement struct {
	Elements []*CommandStatementElement
}

type CommandStatementElement struct {
	text       string
	Expression *Expression
}

func (cs *CommandStatement) rearrange() {
	arrangedElements := []*CommandStatementElement{}
	stringAcc := ""
	for i, element := range cs.Elements {
		isText := element.text != ""
		if isText {
			stringAcc += element.text
		}
		if i == len(cs.Elements)-1 || !isText {
			arrangedElements = append(arrangedElements, cs.split(stringAcc)...)
			stringAcc = ""
		}
		if !isText {
			arrangedElements = append(arrangedElements, element)
		}
	}

	cs.Elements = arrangedElements
}

func (cs *CommandStatement) split(str string) []*CommandStatementElement {
	split := strings.Split(str, " ")
	elements := make([]*CommandStatementElement, 0, len(split))
	for _, word := range split {
		if word == "" {
			continue
		}
		value := valueFromCommandText(word)
		elements = append(elements, &CommandStatementElement{
			Expression: &Expression{Value: value},
		})
	}

	return elements
}

func valueFromCommandText(commandText string) *variable.Value {
	if commandText == "true" {
		return variable.NewBoolean(true)
	} else if commandText == "false" {
		return variable.NewBoolean(false)
	}

	if commandText[0] == '+' { // see Antlr grammar, numbers don't start with + even though Go would be happy to parse them
		return variable.NewString(commandText)
	}
	numberValue, err := strconv.ParseFloat(commandText, 64)
	if err == nil {
		return variable.NewNumber(numberValue)
	}
	return variable.NewString(commandText)
}

type CallStatement struct {
	*FunctionCall
}

type DeclareStatement struct {
	VariableID string
	Value      *Expression
}

type ReturnStatement struct {
}
