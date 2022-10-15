package tree_test

import (
	"os"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/bradleyjkemp/cupaloy"

	"github.com/RemiEven/ysgo/parser"
	"github.com/RemiEven/ysgo/tree"
)

func TestParserListener(t *testing.T) {
	tests := []string{
		"script",
	}

	for _, name := range tests {
		t.Run(name, func(t *testing.T) {
			script, err := os.ReadFile("testdata/" + name + ".yarn")
			if err != nil {
				t.Errorf("failed to read testdata file: %v", err)
				return
			}

			var (
				inputStream = antlr.NewInputStream(string(script))
				lexer       = parser.NewYarnSpinnerLexer(inputStream)
				stream      = antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)
				p           = parser.NewYarnSpinnerParser(stream)
				listener    = &tree.ParserListener{}
			)

			antlr.ParseTreeWalkerDefault.Walk(listener, p.Dialogue())

			dialogue := listener.Dialogue()
			cupaloy.SnapshotT(t, dialogue)
		})
	}
}
