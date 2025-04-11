package tree

import (
	"strconv"

	"github.com/antlr4-go/antlr/v4"

	"github.com/remieven/ysgo/internal/container"
	"github.com/remieven/ysgo/internal/parser"
	"github.com/remieven/ysgo/variable"
)

// parserListener is a complete listener for a parse tree produced by YarnSpinnerParser.
type parserListener struct {
	*parser.BaseYarnSpinnerParserListener
	dialogue *Dialogue

	node                     *Node
	lineStatement            *LineStatement
	shortcutOptionStatements *container.Stack[*ShortcutOptionStatement]
	shortcutOptions          *container.Stack[*ShortcutOption]
	lineGroupStatements      *container.Stack[*LineGroupStatement]
	lineGroupItems           *container.Stack[*LineGroupItem]
	statementCallbacks       *container.Stack[func(*Statement)]
	textCallback             func(string)
	expressionCallbacks      *container.Stack[func(*Expression)]
	lineStatementCallbacks   *container.Stack[func(*LineStatement)]
	variableCallback         func(string)
	clauseCallbacks          *container.Stack[func(*Clause)]
	functionCallCallback     func(*FunctionCall)
	commandTextCallback      func(string)
	hashtagCallback          func(string)
	protoCommandStatement    *CommandStatement
}

// VisitTerminal is called when a terminal node is visited.
func (pl *parserListener) VisitTerminal(node antlr.TerminalNode) {
	switch node.GetSymbol().GetTokenType() {
	case parser.YarnSpinnerLexerTEXT:
		pl.textCallback(node.GetSymbol().GetText())
	case parser.YarnSpinnerLexerCOMMAND_TEXT:
		pl.commandTextCallback(node.GetSymbol().GetText())
	}
}

// EnterDialogue is called when production dialogue is entered.
func (pl *parserListener) EnterDialogue(ctx *parser.DialogueContext) {
	pl.dialogue = &Dialogue{}
	pl.shortcutOptionStatements = &container.Stack[*ShortcutOptionStatement]{}
	pl.shortcutOptions = &container.Stack[*ShortcutOption]{}
	pl.lineGroupStatements = &container.Stack[*LineGroupStatement]{}
	pl.lineGroupItems = &container.Stack[*LineGroupItem]{}
	pl.statementCallbacks = &container.Stack[func(*Statement)]{}
	pl.lineStatementCallbacks = &container.Stack[func(*LineStatement)]{}
	pl.expressionCallbacks = &container.Stack[func(*Expression)]{}
	pl.clauseCallbacks = &container.Stack[func(*Clause)]{}
}

// EnterNode is called when production node is entered.
func (s *parserListener) EnterNode(ctx *parser.NodeContext) {
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
func (pl *parserListener) ExitNode(ctx *parser.NodeContext) {
	pl.dialogue.Nodes = append(pl.dialogue.Nodes, *pl.node)
	pl.statementCallbacks.Pop()
	pl.lineStatementCallbacks.Pop()
}

// EnterHeader is called when production header is entered.
func (s *parserListener) EnterHeader(ctx *parser.HeaderContext) {
	headerKey := ctx.GetHeader_key().GetText()
	headerValue := ""
	if ctx.GetHeader_value() != nil {
		headerValue = ctx.GetHeader_value().GetText()
	}
	s.node.Headers[headerKey] = headerValue
}

// EnterTitle_header is called when production title_header is entered.
func (s *parserListener) EnterTitle_header(ctx *parser.Title_headerContext) {
	s.node.Headers["title"] = ctx.GetTitle().GetText()
}

// EnterLine_statement is called when production line_statement is entered.
func (s *parserListener) EnterLine_statement(ctx *parser.Line_statementContext) {
	s.lineStatement = &LineStatement{}
	s.hashtagCallback = func(tag string) {
		s.lineStatement.Tags = append(s.lineStatement.Tags, tag)
	}
	s.expressionCallbacks.Push(func(e *Expression) {
		s.lineStatement.Condition = e
	})
}

// ExitLine_statement is called when production line_statement is exited.
func (s *parserListener) ExitLine_statement(ctx *parser.Line_statementContext) {
	s.hashtagCallback = nil
	s.expressionCallbacks.Pop()
	s.lineStatementCallbacks.Peek()(s.lineStatement)
	s.lineStatement = nil
}

// EnterLine_formatted_text is called when production line_formatted_text is entered.
func (s *parserListener) EnterLine_formatted_text(ctx *parser.Line_formatted_textContext) {
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
func (s *parserListener) ExitLine_formatted_text(ctx *parser.Line_formatted_textContext) {
	s.textCallback = nil
	s.expressionCallbacks.Pop()
}

// EnterHashtag is called when production hashtag is entered.
func (s *parserListener) EnterHashtag(ctx *parser.HashtagContext) {
	s.hashtagCallback(ctx.GetText()[1:])
}

// EnterShortcut_option_statement is called when production shortcut_option_statement is entered.
func (s *parserListener) EnterShortcut_option_statement(ctx *parser.Shortcut_option_statementContext) {
	s.shortcutOptionStatements.Push(&ShortcutOptionStatement{})
}

// ExitShortcut_option_statement is called when production shortcut_option_statement is exited.
func (s *parserListener) ExitShortcut_option_statement(ctx *parser.Shortcut_option_statementContext) {
	statement := s.shortcutOptionStatements.Pop()
	s.statementCallbacks.Peek()(&Statement{
		ShortcutOptionStatement: statement,
	})
}

// EnterShortcut_option is called when production shortcut_option is entered.
func (s *parserListener) EnterShortcut_option(ctx *parser.Shortcut_optionContext) {
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
func (s *parserListener) ExitShortcut_option(ctx *parser.Shortcut_optionContext) {
	s.lineStatementCallbacks.Pop()
	s.statementCallbacks.Pop()
}

// EnterLine_group_statement is called when production line_group_statement is entered.
func (s *parserListener) EnterLine_group_statement(ctx *parser.Line_group_statementContext) {
	s.lineGroupStatements.Push(&LineGroupStatement{})
}

// ExitLine_group_statement is called when production line_group_statement is exited.
func (s *parserListener) ExitLine_group_statement(ctx *parser.Line_group_statementContext) {
	statement := s.lineGroupStatements.Pop()
	s.statementCallbacks.Peek()(&Statement{
		LineGroupStatement: statement,
	})
}

// EnterLine_group_item is called when production line_group_item is entered.
func (s *parserListener) EnterLine_group_item(ctx *parser.Line_group_itemContext) {
	lineGroupItem := &LineGroupItem{}
	lineGroupStatement := s.lineGroupStatements.Peek()
	lineGroupStatement.Items = append(lineGroupStatement.Items, lineGroupItem)
	s.lineStatementCallbacks.Push(func(lineStatement *LineStatement) {
		if lineGroupItem.LineStatement == nil {
			lineGroupItem.LineStatement = lineStatement
		} else {
			lineGroupItem.Statements = append(lineGroupItem.Statements, &Statement{
				LineStatement: lineStatement,
			})
		}
	})
	s.statementCallbacks.Push(func(statement *Statement) {
		lineGroupItem.Statements = append(lineGroupItem.Statements, statement)
	})
}

// ExitLine_group_item is called when production line_group_item is exited.
func (s *parserListener) ExitLine_group_item(ctx *parser.Line_group_itemContext) {
	s.lineStatementCallbacks.Pop()
	s.statementCallbacks.Pop()
}

// EnterValueNumber is called when production valueNumber is entered.
func (s *parserListener) EnterValueNumber(ctx *parser.ValueNumberContext) {
	number, _ := strconv.ParseFloat(ctx.GetText(), 64)
	s.expressionCallbacks.Peek()(&Expression{
		Value: variable.NewNumber(number),
	})
}

// EnterValueTrue is called when production valueTrue is entered.
func (s *parserListener) EnterValueTrue(ctx *parser.ValueTrueContext) {
	s.expressionCallbacks.Peek()(&Expression{
		Value: variable.NewBoolean(true),
	})
}

// EnterValueFalse is called when production valueFalse is entered.
func (s *parserListener) EnterValueFalse(ctx *parser.ValueFalseContext) {
	s.expressionCallbacks.Peek()(&Expression{
		Value: variable.NewBoolean(false),
	})
}

// EnterValueVar is called when production valueVar is entered.
func (s *parserListener) EnterValueVar(ctx *parser.ValueVarContext) {
	variableID := ctx.GetText()[1:]
	s.expressionCallbacks.Peek()(&Expression{
		VariableID: &variableID,
	})
}

// EnterValueString is called when production valueString is entered.
func (s *parserListener) EnterValueString(ctx *parser.ValueStringContext) {
	text := ctx.GetText()
	value := text[1 : len(text)-1]
	s.expressionCallbacks.Peek()(&Expression{
		Value: variable.NewString(value),
	})
}

// EnterValueFunc is called when production valueFunc is entered.
func (s *parserListener) EnterValueFunc(ctx *parser.ValueFuncContext) {
	s.functionCallCallback = func(functionCall *FunctionCall) {
		s.functionCallCallback = nil
		s.expressionCallbacks.Peek()(&Expression{
			FunctionCall: functionCall,
		})
	}
}

// EnterFunction_call is called when production function_call is entered.
func (s *parserListener) EnterFunction_call(ctx *parser.Function_callContext) {
	functionID := ctx.FUNC_ID().GetText()
	functionCall := &FunctionCall{
		FunctionID: functionID,
	}
	s.functionCallCallback(functionCall)
	s.expressionCallbacks.Push(func(e *Expression) {
		functionCall.Arguments = append(functionCall.Arguments, e)
	})
}

// ExitFunction_call is called when production function_call is exited.
func (s *parserListener) ExitFunction_call(ctx *parser.Function_callContext) {
	s.expressionCallbacks.Pop()
}

// EnterExpNot is called when production expNot is entered.
func (s *parserListener) EnterExpNot(ctx *parser.ExpNotContext) {
	s.expressionCallbacks.Push(func(e *Expression) {
		s.expressionCallbacks.Pop()
		s.expressionCallbacks.Peek()(&Expression{
			NotExpression: e,
		})
	})
}

// EnterExpNegative is called when production expNegative is entered.
func (s *parserListener) EnterExpNegative(ctx *parser.ExpNegativeContext) {
	s.expressionCallbacks.Push(func(e *Expression) {
		s.expressionCallbacks.Pop()
		s.expressionCallbacks.Peek()(&Expression{
			NegativeExpression: e,
		})
	})
}

// EnterExpMultDivMod is called when production expMultDivMod is entered.
func (s *parserListener) EnterExpMultDivMod(ctx *parser.ExpMultDivModContext) {
	s.enterBinaryOperatorExpression(ctx)
}

// EnterExpComparison is called when production expComparison is entered.
func (s *parserListener) EnterExpComparison(ctx *parser.ExpComparisonContext) {
	s.enterBinaryOperatorExpression(ctx)
}

// EnterExpAndOrXor is called when production expAndOrXor is entered.
func (s *parserListener) EnterExpAndOrXor(ctx *parser.ExpAndOrXorContext) {
	s.enterBinaryOperatorExpression(ctx)
}

// EnterExpAddSub is called when production expAddSub is entered.
func (s *parserListener) EnterExpAddSub(ctx *parser.ExpAddSubContext) {
	s.enterBinaryOperatorExpression(ctx)
}

// EnterExpEquality is called when production expEquality is entered.
func (s *parserListener) EnterExpEquality(ctx *parser.ExpEqualityContext) {
	s.enterBinaryOperatorExpression(ctx)
}

func (s *parserListener) enterBinaryOperatorExpression(ctx interface {
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

// EnterSet_statement is called when production set_statement is entered.
func (s *parserListener) EnterSet_statement(ctx *parser.Set_statementContext) {
	inPlaceOperator, ok := tokenToInplaceOperator(ctx.GetOp().GetTokenType())
	if !ok {
		panic("unknown in-place operator") // TODO: better error handling
	}
	var variableID string
	s.variableCallback = func(id string) {
		variableID = id
		s.variableCallback = nil
	}
	s.expressionCallbacks.Push(func(e *Expression) {
		s.statementCallbacks.Peek()(&Statement{
			SetStatement: &SetStatement{
				VariableID:      variableID,
				InPlaceOperator: *inPlaceOperator,
				Expression:      e,
			},
		})
	})
}

// ExitSet_statement is called when production set_statement is exited.
func (s *parserListener) ExitSet_statement(ctx *parser.Set_statementContext) {
	s.expressionCallbacks.Pop()
}

// EnterVariable is called when production variable is entered.
func (s *parserListener) EnterVariable(ctx *parser.VariableContext) {
	if s.variableCallback == nil {
		return
	}
	variableID := ctx.GetText()[1:]
	s.variableCallback(variableID)
}

// EnterJumpToNodeName is called when production jumpToNodeName is entered.
func (s *parserListener) EnterJumpToNodeName(ctx *parser.JumpToNodeNameContext) {
	s.statementCallbacks.Peek()(&Statement{
		JumpStatement: &JumpStatement{
			Expression: &Expression{
				Value: variable.NewString(ctx.GetDestination().GetText()),
			},
		},
	})
}

// EnterJumpToExpression is called when production jumpToExpression is entered.
func (s *parserListener) EnterJumpToExpression(ctx *parser.JumpToExpressionContext) {
	s.expressionCallbacks.Push(func(e *Expression) {
		s.statementCallbacks.Peek()(&Statement{
			JumpStatement: &JumpStatement{
				Expression: e,
			},
		})
	})
}

// ExitJumpToExpression is called when production jumpToExpression is exited.
func (s *parserListener) ExitJumpToExpression(ctx *parser.JumpToExpressionContext) {
	s.expressionCallbacks.Pop()
}

// EnterDetourToNodeName is called when production detourToNodeName is entered.
func (s *parserListener) EnterDetourToNodeName(ctx *parser.DetourToNodeNameContext) {
	s.statementCallbacks.Peek()(&Statement{
		JumpStatement: &JumpStatement{
			Expression: &Expression{
				Value: variable.NewString(ctx.GetDestination().GetText()),
			},
			Detour: true,
		},
	})
}

// EnterDetourToExpression is called when production detourToExpression is entered.
func (s *parserListener) EnterDetourToExpression(ctx *parser.DetourToExpressionContext) {
	s.expressionCallbacks.Push(func(e *Expression) {
		s.statementCallbacks.Peek()(&Statement{
			JumpStatement: &JumpStatement{
				Expression: e,
				Detour:     true,
			},
		})
	})
}

// ExitDetourToExpression is called when production detourToExpression is exited.
func (s *parserListener) ExitDetourToExpression(ctx *parser.DetourToExpressionContext) {
	s.expressionCallbacks.Pop()
}

// EnterReturn_statement is called when production return_statement is entered.
func (s *parserListener) EnterReturn_statement(ctx *parser.Return_statementContext) {
	s.statementCallbacks.Peek()(&Statement{
		ReturnStatement: &ReturnStatement{},
	})
}

// EnterIf_statement is called when production if_statement is entered.
func (s *parserListener) EnterIf_statement(ctx *parser.If_statementContext) {
	statement := &IfStatement{}
	s.statementCallbacks.Peek()(&Statement{
		IfStatement: statement,
	})
	s.clauseCallbacks.Push(func(clause *Clause) {
		statement.Clauses = append(statement.Clauses, clause)
	})
}

// ExitIf_statement is called when production if_statement is exited.
func (s *parserListener) ExitIf_statement(ctx *parser.If_statementContext) {
	s.clauseCallbacks.Pop()
}

// EnterIf_clause is called when production if_clause is entered.
func (s *parserListener) EnterIf_clause(ctx *parser.If_clauseContext) {
	clause := &Clause{}
	s.expressionCallbacks.Push(func(e *Expression) {
		clause.Condition = e
		s.expressionCallbacks.Pop()
	})
	s.statementCallbacks.Push(func(statement *Statement) {
		clause.Statements = append(clause.Statements, statement)
	})
	s.lineStatementCallbacks.Push(func(ls *LineStatement) {
		clause.Statements = append(clause.Statements, &Statement{LineStatement: ls})
	})
	s.clauseCallbacks.Peek()(clause)
}

// ExitIf_clause is called when production if_clause is exited.
func (s *parserListener) ExitIf_clause(ctx *parser.If_clauseContext) {
	s.statementCallbacks.Pop()
	s.lineStatementCallbacks.Pop()
}

// EnterElse_if_clause is called when production else_if_clause is entered.
func (s *parserListener) EnterElse_if_clause(ctx *parser.Else_if_clauseContext) {
	clause := &Clause{}
	s.expressionCallbacks.Push(func(e *Expression) {
		clause.Condition = e
		s.expressionCallbacks.Pop()
	})
	s.statementCallbacks.Push(func(statement *Statement) {
		clause.Statements = append(clause.Statements, statement)
	})
	s.lineStatementCallbacks.Push(func(ls *LineStatement) {
		clause.Statements = append(clause.Statements, &Statement{LineStatement: ls})
	})
	s.clauseCallbacks.Peek()(clause)
}

// ExitElse_if_clause is called when production else_if_clause is exited.
func (s *parserListener) ExitElse_if_clause(ctx *parser.Else_if_clauseContext) {
	s.statementCallbacks.Pop()
	s.lineStatementCallbacks.Pop()
}

// EnterElse_clause is called when production else_clause is entered.
func (s *parserListener) EnterElse_clause(ctx *parser.Else_clauseContext) {
	clause := &Clause{
		Condition: &Expression{
			Value: variable.NewBoolean(true),
		},
	}
	s.statementCallbacks.Push(func(statement *Statement) {
		clause.Statements = append(clause.Statements, statement)
	})
	s.lineStatementCallbacks.Push(func(ls *LineStatement) {
		clause.Statements = append(clause.Statements, &Statement{LineStatement: ls})
	})
	s.clauseCallbacks.Peek()(clause)
}

// ExitElse_clause is called when production else_clause is exited.
func (s *parserListener) ExitElse_clause(ctx *parser.Else_clauseContext) {
	s.statementCallbacks.Pop()
	s.lineStatementCallbacks.Pop()
}

// EnterCommand_statement is called when production command_statement is entered.
func (s *parserListener) EnterCommand_statement(ctx *parser.Command_statementContext) {
	commandStatement := &CommandStatement{}
	s.commandTextCallback = func(commandText string) {
		commandStatement.Elements = append(commandStatement.Elements, &CommandStatementElement{
			text: commandText,
		})
	}
	s.expressionCallbacks.Push(func(e *Expression) {
		commandStatement.Elements = append(commandStatement.Elements, &CommandStatementElement{
			Expression: e,
		})
	})
	s.statementCallbacks.Peek()(&Statement{
		CommandStatement: commandStatement,
	})
	s.protoCommandStatement = commandStatement
}

// ExitCommand_statement is called when production command_statement is exited.
func (s *parserListener) ExitCommand_statement(ctx *parser.Command_statementContext) {
	s.commandTextCallback = nil
	s.expressionCallbacks.Pop()
	s.protoCommandStatement.rearrange()
	s.protoCommandStatement = nil
}

// EnterCall_statement is called when production call_statement is entered.
func (s *parserListener) EnterCall_statement(ctx *parser.Call_statementContext) {
	s.functionCallCallback = func(call *FunctionCall) {
		s.statementCallbacks.Peek()(&Statement{
			CallStatement: &CallStatement{
				FunctionCall: call,
			},
		})
	}
}

// ExitCall_statement is called when production call_statement is exited.
func (s *parserListener) ExitCall_statement(ctx *parser.Call_statementContext) {
	s.functionCallCallback = nil
}

// EnterDeclare_statement is called when production declare_statement is entered.
func (s *parserListener) EnterDeclare_statement(ctx *parser.Declare_statementContext) {
	declareStatement := &DeclareStatement{}
	s.statementCallbacks.Peek()(&Statement{
		DeclareStatement: declareStatement,
	})
	s.expressionCallbacks.Push(func(e *Expression) {
		declareStatement.Value = e
	})
	s.variableCallback = func(variableID string) {
		declareStatement.VariableID = variableID
		s.variableCallback = nil
	}
}

// ExitDeclare_statement is called when production declare_statement is exited.
func (s *parserListener) ExitDeclare_statement(ctx *parser.Declare_statementContext) {
	s.expressionCallbacks.Pop()
}
