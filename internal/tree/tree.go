package tree

import (
	"strconv"
	"strings"

	"github.com/remieven/ysgo/variable"
)

type Dialogue struct {
	Nodes []Node `json:"n"`
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
	Headers    map[string]string `json:"h,omitempty"`
	Statements []*Statement      `json:"s"`
}

func (n *Node) Title() string {
	return n.Headers["title"]
}

type Statement struct {
	LineStatement           *LineStatement           `json:"ls,omitempty"`
	ShortcutOptionStatement *ShortcutOptionStatement `json:"sos,omitempty"`
	LineGroupStatement      *LineGroupStatement      `json:"lgs,omitempty"`
	SetStatement            *SetStatement            `json:"set-s,omitempty"`
	JumpStatement           *JumpStatement           `json:"js,omitempty"`
	IfStatement             *IfStatement             `json:"is,omitempty"`
	CommandStatement        *CommandStatement        `json:"cs,omitempty"`
	CallStatement           *CallStatement           `json:"call-s,omitempty"`
	DeclareStatement        *DeclareStatement        `json:"ds,omitempty"`
	ReturnStatement         *ReturnStatement         `json:"rs,omitempty"`
}

type LineStatement struct {
	Text      *LineFormattedText `json:"t,omitempty"`
	Condition *Expression        `json:"c,omitempty"`
	Tags      []string           `json:"tags,omitempty"`
}

type LineFormattedText struct {
	Elements []*LineFormattedTextElement `json:"e"`
}

type LineFormattedTextElement struct {
	Text       string      `json:"t"`
	Expression *Expression `json:"e,omitempty"`
}

type ShortcutOptionStatement struct {
	Options []*ShortcutOption `json:"so"`
}

type ShortcutOption struct {
	LineStatement *LineStatement `json:"ls"`
	Statements    []*Statement   `json:"s,omitempty"`
}

type LineGroupStatement struct {
	Items []*LineGroupItem `json:"i"`
}

type LineGroupItem struct {
	LineStatement *LineStatement `json:"ls"`
	Statements    []*Statement   `json:"s,omitempty"`
}

type SetStatement struct {
	VariableID      string      `json:"v"`
	InPlaceOperator int         `json:"ipo"`
	Expression      *Expression `json:"e"`
}

type JumpStatement struct {
	Expression *Expression `json:"e"`
	Detour     bool        `json:"d"`
}

type IfStatement struct {
	Clauses []*Clause `json:"c"`
}

type Clause struct {
	Condition  *Expression  `json:"c"`
	Statements []*Statement `json:"s"`
}

type CommandStatement struct {
	Elements []*CommandStatementElement `json:"e"`
}

type CommandStatementElement struct {
	text       string
	Expression *Expression `json:"e,omitempty"`
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
	*FunctionCall  `json:"fc"`
}

type DeclareStatement struct {
	VariableID string `json:"c"`
	Value      *Expression `json:"val,omitempty"`
}

type ReturnStatement struct {
}
