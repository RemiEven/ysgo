package tree

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
}

type LineStatement struct {
	Text      *LineFormattedText
	Condition *Expression
	// Hashtags  []string
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
