package main

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/RemiEven/ysgo/parser"
	"github.com/RemiEven/ysgo/tree"
)

func main() {
	inputStream, err := antlr.NewFileStream("/home/remi/projects/ysgo/script.yarn")
	if err != nil {
		panic(err)
	}

	lexer := parser.NewYarnSpinnerLexer(inputStream)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewYarnSpinnerParser(stream)

	listener := &tree.ParserListener{}
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Dialogue())

	di := listener.Dialogue()
	fmt.Println(di.Nodes[0].Headers["title"])
}
