package parser

import (
	"strings"

	"github.com/antlr4-go/antlr/v4"

	"github.com/remieven/ysgo/internal/container"
)

// IndentAwareLexer is a lexer that keeps track of indentations level when it lexes its input.
// It's pretty much a direct port of the IndentAwareLexer.cs file of YarnSpinner from C# to Go.
type IndentAwareLexer struct {
	*antlr.BaseLexer
	hitEOF                               bool
	pendingTokens                        container.Queue[antlr.Token]
	unbalancedIndents                    container.Stack[int]
	lastIndent                           int
	lineContainsIndentTrackingToken      bool
	lastToken                            antlr.Token
	lastSeenIndentTrackingContentPlusOne int
	inWhenClause                         bool
	tokens                               container.Queue[antlr.Token]
}

// NextToken returns a token from the lexer source i.e., match a token on the char stream.
func (ial *IndentAwareLexer) NextToken() antlr.Token {
	var tokenToReturn antlr.Token

	if ial.hitEOF && ial.pendingTokens.Size() > 0 {
		tokenToReturn = ial.pendingTokens.Dequeue()
	} else if ial.GetInputStream().Size() == 0 {
		ial.hitEOF = true
		tokenToReturn = antlr.NewCommonToken(nil, antlr.TokenEOF, antlr.TokenDefaultChannel, -1, -1)
	} else {
		ial.checkNextToken()
		if ial.pendingTokens.Size() > 0 {
			tokenToReturn = ial.pendingTokens.Dequeue()
		}
	}

	if tokenToReturn != nil {
		ial.pushToken(tokenToReturn)
	}
	return tokenToReturn
}

func (ial *IndentAwareLexer) checkNextToken() {
	currentToken := ial.BaseLexer.NextToken()

	switch currentToken.GetTokenType() {
	case YarnSpinnerLexerNEWLINE:
		ial.handleNewLineToken(currentToken)
	case antlr.TokenEOF:
		ial.handleEndOfFileToken(currentToken)
	case YarnSpinnerLexerSHORTCUT_ARROW, YarnSpinnerLexerLINE_GROUP_ARROW:
		ial.pendingTokens.Enqueue(currentToken)
		ial.lineContainsIndentTrackingToken = true
	case YarnSpinnerLexerBODY_END:
		ial.lineContainsIndentTrackingToken = false
		ial.lastIndent = 0
		ial.unbalancedIndents.Clear()
		ial.lastSeenIndentTrackingContentPlusOne = 0
		ial.pendingTokens.Enqueue(currentToken)
	default:
		ial.pendingTokens.Enqueue(currentToken)
	}
	ial.lastToken = currentToken
}

func (ial *IndentAwareLexer) handleNewLineToken(currentToken antlr.Token) {
	ial.pendingTokens.Enqueue(currentToken)

	currentIndentationLength := ial.getLengthOfNewlineToken(currentToken)

	if ial.lastSeenIndentTrackingContentPlusOne != 0 {
		if currentToken.GetTokenType() == ial.lastToken.GetTokenType() {
			if ial.GetLine()-ial.lastSeenIndentTrackingContentPlusOne == 0 {
				ial.insertToken("", YarnSpinnerLexerBLANK_LINE_FOLLOWING_OPTION)
			}
			ial.lastSeenIndentTrackingContentPlusOne = 0
		}
	}

	if ial.lineContainsIndentTrackingToken {
		if currentIndentationLength > ial.lastIndent {
			ial.unbalancedIndents.Push(currentIndentationLength)
			ial.insertToken("", YarnSpinnerLexerINDENT)
		}
		ial.lineContainsIndentTrackingToken = false
		ial.lastSeenIndentTrackingContentPlusOne = ial.GetLine() + 1
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
				ial.lastSeenIndentTrackingContentPlusOne = ial.GetLine() + 1
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

func (ial *IndentAwareLexer) pushToken(token antlr.Token) {
	ial.tokens.Enqueue(token)
	if ial.tokens.Size() > 5 {
		ial.tokens.Dequeue()
	}
}

func (ial *IndentAwareLexer) lastTokenWas(tokenType int, text string) bool {
	if ial.tokens.Size() == 0 {
		return false
	}

	token := ial.tokens.Peek()
	return token.GetTokenType() == tokenType && strings.EqualFold(token.GetText(), text)
}
