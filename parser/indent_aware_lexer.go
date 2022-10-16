package parser

import (
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"

	"github.com/RemiEven/ysgo/container"
)

type IndentAwareLexer struct {
	*antlr.BaseLexer
	hitEOF        bool
	pendingTokens container.Queue[antlr.Token]
	indents       container.Stack[int]
}

func (ial *IndentAwareLexer) NextToken() antlr.Token {
	if ial.hitEOF && ial.pendingTokens.Size() > 0 {
		return ial.pendingTokens.Dequeue()
	}
	if ial.GetInputStream().Size() == 0 {
		ial.hitEOF = true
		return antlr.NewCommonToken(nil, antlr.TokenEOF, antlr.TokenDefaultChannel, -1, -1)
	}

	ial.checkNextToken()
	if ial.pendingTokens.Size() > 0 {
		return ial.pendingTokens.Dequeue()
	}
	return nil
}

func (ial *IndentAwareLexer) checkNextToken() {
	currentToken := ial.BaseLexer.NextToken()

	switch currentToken.GetTokenType() {
	case YarnSpinnerLexerNEWLINE:
		ial.handleNewLineToken(currentToken)
	case antlr.TokenEOF:
		ial.handleEndOfFileToken(currentToken)
	default:
		ial.pendingTokens.Enqueue(currentToken)
	}
}

func (ial *IndentAwareLexer) handleNewLineToken(currentToken antlr.Token) {
	ial.pendingTokens.Enqueue(currentToken)

	currentIndentationLength := ial.getLengthOfNewlineToken(currentToken)

	previousIndent := 0
	if ial.indents.Size() > 0 {
		previousIndent = ial.indents.Peek()
	}

	if currentIndentationLength > previousIndent {
		ial.indents.Push(currentIndentationLength)
		ial.insertToken("<indent to "+strconv.Itoa(currentIndentationLength)+">", YarnSpinnerLexerINDENT)
	} else if currentIndentationLength < previousIndent {
		for currentIndentationLength < previousIndent {
			previousIndent = ial.indents.Pop()

			ial.insertToken("<dedent from "+strconv.Itoa(previousIndent)+">", YarnSpinnerLexerDEDENT)

			if ial.indents.Size() > 0 {
				previousIndent = ial.indents.Peek()
			} else {
				previousIndent = 0
			}
		}
	}
}

func (ial *IndentAwareLexer) getLengthOfNewlineToken(currentToken antlr.Token) int {
	length := 0
	var sawSpaces, sawTabs bool

	for _, c := range currentToken.GetText() {
		switch c {
		case ' ':
			length++
			sawSpaces = true
		case '\t':
			length += 8
			sawTabs = true
		}
	}

	if sawSpaces && sawTabs {
		panic("Indentation contains tabs and spaces")
	}

	return length
}

func (ial *IndentAwareLexer) handleEndOfFileToken(currentToken antlr.Token) {
	for ial.indents.Size() > 0 {
		previousIndent := ial.indents.Pop()

		ial.insertToken("<dedent from "+strconv.Itoa(previousIndent)+">", YarnSpinnerLexerDEDENT)
	}

	ial.pendingTokens.Enqueue(currentToken)
}

func (ial *IndentAwareLexer) insertToken(text string, tokenType int) {
	startIndex := ial.TokenStartCharIndex + len(ial.GetText())
	stopIndex := startIndex - 1
	token := antlr.NewCommonToken(ial.GetTokenSourceCharStreamPair(), tokenType, antlr.TokenDefaultChannel, startIndex, stopIndex)
	token.SetText(text)
	ial.pendingTokens.Enqueue(token)
}
