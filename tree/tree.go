package tree

type Dialogue struct {
	Nodes []Node
}

type Node struct {
	Headers    map[string]string
	Statements []*Statement
}

type Statement struct {
	LineStatement           *LineStatement
	ShortcutOptionStatement *ShortcutOptionStatement
}

type LineStatement struct {
	Text *LineFormattedText
	// Condition *Expression
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

type Expression struct {
}
