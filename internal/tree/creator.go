package tree

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"

	"github.com/RemiEven/ysgo/internal/parser"
)

// FromString creates a dialogue tree by reading scriptData.
func FromString(scriptData string) (*Dialogue, error) {
	return fromInputStream(antlr.NewInputStream(scriptData))
}

// FromFile creates a dialogue tree by reading the file located at scriptPath.
func FromFile(scriptPath string) (*Dialogue, error) {
	inputStream, err := antlr.NewFileStream(scriptPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create file input stream: %w", err)
	}

	return fromInputStream(inputStream)
}

func fromInputStream(input antlr.CharStream) (*Dialogue, error) {
	var (
		lexer    = parser.NewYarnSpinnerLexer(input)
		stream   = antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)
		p        = parser.NewYarnSpinnerParser(stream)
		listener = &parserListener{}
	)

	antlr.ParseTreeWalkerDefault.Walk(listener, p.Dialogue())

	return listener.dialogue, nil
}
