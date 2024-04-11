package parser

import (
	"github.com/antlr4-go/antlr/v4"

	"github.com/remieven/ysgo/internal/container"
)

// IndentAwareLexer is a lexer that keeps track of indentations level when it lexes its input.
// It's pretty much a direct port of the IndentAwareLexer.cs file of YarnSpinner from C# to Go.
type IndentAwareLexer struct {
	*antlr.BaseLexer
	hitEOF                       bool
	pendingTokens                container.Queue[antlr.Token]
	unbalancedIndents            container.Stack[int]
	lastIndent                   int
	lineContainsShortcut         bool
	lastToken                    antlr.Token
	lastSeenOptionContentPlusOne int
}

// NextToken returns a token from the lexer source i.e., match a token on the char stream.
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
	case YarnSpinnerLexerSHORTCUT_ARROW:
		ial.pendingTokens.Enqueue(currentToken)
		ial.lineContainsShortcut = true
	case YarnSpinnerLexerBODY_END:
		ial.lineContainsShortcut = false
		ial.lastIndent = 0
		ial.unbalancedIndents.Clear()
		ial.lastSeenOptionContentPlusOne = 0
		ial.pendingTokens.Enqueue(currentToken)
	default:
		ial.pendingTokens.Enqueue(currentToken)
	}
	ial.lastToken = currentToken
}

func (ial *IndentAwareLexer) handleNewLineToken(currentToken antlr.Token) {
	ial.pendingTokens.Enqueue(currentToken)

	currentIndentationLength := ial.getLengthOfNewlineToken(currentToken)

	if ial.lastSeenOptionContentPlusOne != 0 {
		if currentToken.GetTokenType() == ial.lastToken.GetTokenType() {
			if ial.GetLine()-ial.lastSeenOptionContentPlusOne == 0 {
				ial.insertToken("", YarnSpinnerLexerBLANK_LINE_FOLLOWING_OPTION)
			}
			ial.lastSeenOptionContentPlusOne = 0
		}
	}

	if ial.lineContainsShortcut {
		if currentIndentationLength > ial.lastIndent {
			ial.unbalancedIndents.Push(currentIndentationLength)
			ial.insertToken("", YarnSpinnerLexerINDENT)
		}
		ial.lineContainsShortcut = false
		ial.lastSeenOptionContentPlusOne = ial.GetLine() + 1
	}

	if ial.unbalancedIndents.Size() > 0 {
		top := ial.unbalancedIndents.Peek()

		for currentIndentationLength < top {
			ial.insertToken("", YarnSpinnerLexerDEDENT)
			ial.unbalancedIndents.Pop()
			if ial.unbalancedIndents.Size() > 0 {
				top = ial.unbalancedIndents.Peek()
			} else {
				top = 0
				ial.lastSeenOptionContentPlusOne = ial.GetLine() + 1
			}
		}
	}

	ial.lastIndent = currentIndentationLength
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
	for ial.unbalancedIndents.Size() > 0 {
		ial.unbalancedIndents.Pop()
		ial.insertToken("", YarnSpinnerLexerDEDENT)
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
