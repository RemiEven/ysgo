package tree

import (
	"encoding/json"
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

// FromMarshaledReader creates a dialogue tree by reading a previously
// parsed and marshaled dialogue in reader.
func FromMarshaledReader(reader io.Reader) (*Dialogue, error) {
	var d Dialogue
	if err := json.NewDecoder(reader).Decode(&d); err != nil {
		return nil, fmt.Errorf("failed to decode marshaled dialogue: %w", err)
	}
	return &d, nil
}
