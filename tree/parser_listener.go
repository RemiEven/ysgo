package tree

import (
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/RemiEven/ysgo/container"
	"github.com/RemiEven/ysgo/parser"
)

// ParserListener is a complete listener for a parse tree produced by YarnSpinnerParser.
type ParserListener struct {
	*parser.BaseYarnSpinnerParserListener
	dialogue *Dialogue

	node                     *Node
	lineStatement            *LineStatement
	shortcutOptionStatements *container.Stack[*ShortcutOptionStatement]
	shortcutOptions          *container.Stack[*ShortcutOption]
	statementCallbacks       *container.Stack[func(*Statement)]
	textCallback             func(string)
	expressionCallbacks      *container.Stack[func(*Expression)]
	lineStatementCallbacks   *container.Stack[func(*LineStatement)]
}

func (pl *ParserListener) Dialogue() *Dialogue {
	return pl.dialogue
}

// VisitTerminal is called when a terminal node is visited.
func (pl *ParserListener) VisitTerminal(node antlr.TerminalNode) {
	switch node.GetSymbol().GetTokenType() {
	case parser.YarnSpinnerLexerTEXT:
		pl.textCallback(node.GetSymbol().GetText())
	}
}

// EnterDialogue is called when production dialogue is entered.
func (pl *ParserListener) EnterDialogue(ctx *parser.DialogueContext) {
	pl.dialogue = &Dialogue{}
	pl.shortcutOptionStatements = &container.Stack[*ShortcutOptionStatement]{}
	pl.shortcutOptions = &container.Stack[*ShortcutOption]{}
	pl.statementCallbacks = &container.Stack[func(*Statement)]{}
	pl.lineStatementCallbacks = &container.Stack[func(*LineStatement)]{}
	pl.expressionCallbacks = &container.Stack[func(*Expression)]{}
}

// EnterNode is called when production node is entered.
func (s *ParserListener) EnterNode(ctx *parser.NodeContext) {
	s.node = &Node{
		Headers: map[string]string{},
	}
	s.statementCallbacks.Push(func(statement *Statement) {
		s.node.Statements = append(s.node.Statements, statement)
	})
	s.lineStatementCallbacks.Push(func(lineStatement *LineStatement) {
		s.node.Statements = append(s.node.Statements, &Statement{
			LineStatement: lineStatement,
		})
	})
}

// ExitNode is called when production node is exited.
func (pl *ParserListener) ExitNode(ctx *parser.NodeContext) {
	pl.dialogue.Nodes = append(pl.dialogue.Nodes, *pl.node)
	pl.statementCallbacks.Pop()
	pl.lineStatementCallbacks.Pop()
}

// EnterHeader is called when production header is entered.
func (s *ParserListener) EnterHeader(ctx *parser.HeaderContext) {
	headerKey := ctx.GetHeader_key().GetText()
	headerValue := ctx.GetHeader_value().GetText()
	s.node.Headers[headerKey] = headerValue
}

// EnterLine_statement is called when production line_statement is entered.
func (s *ParserListener) EnterLine_statement(ctx *parser.Line_statementContext) {
	s.lineStatement = &LineStatement{}
}

// ExitLine_statement is called when production line_statement is exited.
func (s *ParserListener) ExitLine_statement(ctx *parser.Line_statementContext) {
	s.lineStatementCallbacks.Peek()(s.lineStatement)
	s.lineStatement = nil
}

// EnterLine_formatted_text is called when production line_formatted_text is entered.
func (s *ParserListener) EnterLine_formatted_text(ctx *parser.Line_formatted_textContext) {
	lineStatementText := &LineFormattedText{}
	s.lineStatement.Text = lineStatementText
	s.textCallback = func(text string) {
		numberOfElements := len(lineStatementText.Elements)
		if numberOfElements > 0 && lineStatementText.Elements[numberOfElements-1].Text != "" {
			lineStatementText.Elements[numberOfElements-1].Text += text
			return
		}
		element := &LineFormattedTextElement{
			Text: text,
		}
		lineStatementText.Elements = append(lineStatementText.Elements, element)
	}
	s.expressionCallbacks.Push(func(expression *Expression) {
		element := &LineFormattedTextElement{
			Expression: expression,
		}
		lineStatementText.Elements = append(lineStatementText.Elements, element)
	})
}

// ExitLine_formatted_text is called when production line_formatted_text is exited.
func (s *ParserListener) ExitLine_formatted_text(ctx *parser.Line_formatted_textContext) {
	s.textCallback = nil
	s.expressionCallbacks.Pop()
}

// EnterShortcut_option_statement is called when production shortcut_option_statement is entered.
func (s *ParserListener) EnterShortcut_option_statement(ctx *parser.Shortcut_option_statementContext) {
	s.shortcutOptionStatements.Push(&ShortcutOptionStatement{})
}

// ExitShortcut_option_statement is called when production shortcut_option_statement is exited.
func (s *ParserListener) ExitShortcut_option_statement(ctx *parser.Shortcut_option_statementContext) {
	statement := s.shortcutOptionStatements.Pop()
	s.statementCallbacks.Peek()(&Statement{
		ShortcutOptionStatement: statement,
	})
}

// EnterShortcut_option is called when production shortcut_option is entered.
func (s *ParserListener) EnterShortcut_option(ctx *parser.Shortcut_optionContext) {
	shortcutOption := &ShortcutOption{}
	shortcutOptionStatement := s.shortcutOptionStatements.Peek()
	shortcutOptionStatement.Options = append(shortcutOptionStatement.Options, shortcutOption)
	s.lineStatementCallbacks.Push(func(lineStatement *LineStatement) {
		if shortcutOption.LineStatement == nil {
			shortcutOption.LineStatement = lineStatement
		} else {
			shortcutOption.Statements = append(shortcutOption.Statements, &Statement{
				LineStatement: lineStatement,
			})
		}
	})
	s.statementCallbacks.Push(func(statement *Statement) {
		shortcutOption.Statements = append(shortcutOption.Statements, statement)
	})
}

// ExitShortcut_option is called when production shortcut_option is exited.
func (s *ParserListener) ExitShortcut_option(ctx *parser.Shortcut_optionContext) {
	s.lineStatementCallbacks.Pop()
	s.statementCallbacks.Pop()
}

// EnterValueNumber is called when production valueNumber is entered.
func (s *ParserListener) EnterValueNumber(ctx *parser.ValueNumberContext) {
	number, _ := strconv.ParseFloat(ctx.GetText(), 64)
	s.expressionCallbacks.Peek()(&Expression{
		Value: NewNumberValue(number),
	})
}

// EnterValueTrue is called when production valueTrue is entered.
func (s *ParserListener) EnterValueTrue(ctx *parser.ValueTrueContext) {
	s.expressionCallbacks.Peek()(&Expression{
		Value: NewBooleanValue(true),
	})
}

// EnterValueFalse is called when production valueFalse is entered.
func (s *ParserListener) EnterValueFalse(ctx *parser.ValueFalseContext) {
	s.expressionCallbacks.Peek()(&Expression{
		Value: NewBooleanValue(false),
	})
}

// EnterValueVar is called when production valueVar is entered.
func (s *ParserListener) EnterValueVar(ctx *parser.ValueVarContext) {
	variableID := ctx.GetText()[1:]
	s.expressionCallbacks.Peek()(&Expression{
		Value: &Value{
			VariableID: &variableID,
		},
	})
}

// EnterValueString is called when production valueString is entered.
func (s *ParserListener) EnterValueString(ctx *parser.ValueStringContext) {
	text := ctx.GetText()
	value := text[1 : len(text)-1]
	s.expressionCallbacks.Peek()(&Expression{
		Value: NewStringValue(value),
	})
}

// EnterValueFunc is called when production valueFunc is entered.
func (s *ParserListener) EnterValueFunc(ctx *parser.ValueFuncContext) {
	// TODO: complete this
}

// EnterExpNot is called when production expNot is entered.
func (s *ParserListener) EnterExpNot(ctx *parser.ExpNotContext) {
	s.expressionCallbacks.Push(func(e *Expression) {
		s.expressionCallbacks.Pop()
		s.expressionCallbacks.Peek()(&Expression{
			NotExpression: e,
		})
	})
}

// EnterExpNegative is called when production expNegative is entered.
func (s *ParserListener) EnterExpNegative(ctx *parser.ExpNegativeContext) {
	s.expressionCallbacks.Push(func(e *Expression) {
		s.expressionCallbacks.Pop()
		s.expressionCallbacks.Peek()(&Expression{
			NegativeExpression: e,
		})
	})
}

// EnterExpMultDivMod is called when production expMultDivMod is entered.
func (s *ParserListener) EnterExpMultDivMod(ctx *parser.ExpMultDivModContext) {
	s.enterBinaryOperatorExpression(ctx)
}

// EnterExpComparison is called when production expComparison is entered.
func (s *ParserListener) EnterExpComparison(ctx *parser.ExpComparisonContext) {
	s.enterBinaryOperatorExpression(ctx)
}

// EnterExpAndOrXor is called when production expAndOrXor is entered.
func (s *ParserListener) EnterExpAndOrXor(ctx *parser.ExpAndOrXorContext) {
	s.enterBinaryOperatorExpression(ctx)
}

// EnterExpAddSub is called when production expAddSub is entered.
func (s *ParserListener) EnterExpAddSub(ctx *parser.ExpAddSubContext) {
	s.enterBinaryOperatorExpression(ctx)
}

// EnterExpEquality is called when production expEquality is entered.
func (s *ParserListener) EnterExpEquality(ctx *parser.ExpEqualityContext) {
	s.enterBinaryOperatorExpression(ctx)
}

func (s *ParserListener) enterBinaryOperatorExpression(ctx interface {
	GetOp() antlr.Token
}) {
	binaryOperator, ok := tokenToBinaryOperator(ctx.GetOp().GetTokenType())
	if !ok {
		panic("unknown binary operator") // TODO: better error handling
	}
	var leftOperand *Expression
	s.expressionCallbacks.Push(func(e *Expression) {
		s.expressionCallbacks.Pop()
		s.expressionCallbacks.Peek()(&Expression{
			Operator:     binaryOperator,
			LeftOperand:  leftOperand,
			RightOperand: e,
		})
	})
	s.expressionCallbacks.Push(func(e *Expression) {
		s.expressionCallbacks.Pop()
		leftOperand = e
	})
}
