package tree

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"

	"github.com/RemiEven/ysgo/internal/parser"
)

func FromString(scriptData string) (*Dialogue, error) {
	return fromInputStream(antlr.NewInputStream(scriptData))
}

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
		listener = &ParserListener{}
	)

	antlr.ParseTreeWalkerDefault.Walk(listener, p.Dialogue())

	return listener.Dialogue(), nil
}
