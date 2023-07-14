package tree

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/antlr4-go/antlr/v4"

	"github.com/remieven/ysgo/internal/parser"
)

// FromFile creates a dialogue tree by reading the file located at scriptPath.
func FromFile(scriptPath string) (*Dialogue, error) {
	fileReader, err := os.Open(scriptPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open dialogue file: %w", err)
	}
	defer fileReader.Close()
	return FromReader(fileReader)
}

// FromReaders creates a dialogue tree by reading the content of readers.
func FromReaders(readers ...io.Reader) (*Dialogue, error) {
	if len(readers) == 0 {
		return nil, errors.New("at least one argument is required")
	}

	dialogue := &Dialogue{}

	for _, reader := range readers {
		d, err := FromReader(reader)
		if err != nil {
			return nil, err
		}
		dialogue.Nodes = append(dialogue.Nodes, d.Nodes...)
	}

	return dialogue, nil
}

// FromReader creates a dialogue tree by reading the content of reader.
func FromReader(reader io.Reader) (*Dialogue, error) {
	scriptData, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to read content: %w", err)
	}
	input := antlr.NewInputStream(string(scriptData))
	var (
		lexer    = parser.NewYarnSpinnerLexer(input)
		stream   = antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)
		p        = parser.NewYarnSpinnerParser(stream)
		listener = &parserListener{}
	)

	antlr.ParseTreeWalkerDefault.Walk(listener, p.Dialogue())

	return listener.dialogue, nil
}
