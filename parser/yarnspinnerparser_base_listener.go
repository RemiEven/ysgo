// Code generated from parser/YarnSpinnerParser.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // YarnSpinnerParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseYarnSpinnerParserListener is a complete listener for a parse tree produced by YarnSpinnerParser.
type BaseYarnSpinnerParserListener struct{}

var _ YarnSpinnerParserListener = &BaseYarnSpinnerParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseYarnSpinnerParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseYarnSpinnerParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseYarnSpinnerParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseYarnSpinnerParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterDialogue is called when production dialogue is entered.
func (s *BaseYarnSpinnerParserListener) EnterDialogue(ctx *DialogueContext) {}

// ExitDialogue is called when production dialogue is exited.
func (s *BaseYarnSpinnerParserListener) ExitDialogue(ctx *DialogueContext) {}

// EnterFile_hashtag is called when production file_hashtag is entered.
func (s *BaseYarnSpinnerParserListener) EnterFile_hashtag(ctx *File_hashtagContext) {}

// ExitFile_hashtag is called when production file_hashtag is exited.
func (s *BaseYarnSpinnerParserListener) ExitFile_hashtag(ctx *File_hashtagContext) {}

// EnterNode is called when production node is entered.
func (s *BaseYarnSpinnerParserListener) EnterNode(ctx *NodeContext) {}

// ExitNode is called when production node is exited.
func (s *BaseYarnSpinnerParserListener) ExitNode(ctx *NodeContext) {}

// EnterHeader is called when production header is entered.
func (s *BaseYarnSpinnerParserListener) EnterHeader(ctx *HeaderContext) {}

// ExitHeader is called when production header is exited.
func (s *BaseYarnSpinnerParserListener) ExitHeader(ctx *HeaderContext) {}

// EnterBody is called when production body is entered.
func (s *BaseYarnSpinnerParserListener) EnterBody(ctx *BodyContext) {}

// ExitBody is called when production body is exited.
func (s *BaseYarnSpinnerParserListener) ExitBody(ctx *BodyContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseYarnSpinnerParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseYarnSpinnerParserListener) ExitStatement(ctx *StatementContext) {}

// EnterLine_statement is called when production line_statement is entered.
func (s *BaseYarnSpinnerParserListener) EnterLine_statement(ctx *Line_statementContext) {}

// ExitLine_statement is called when production line_statement is exited.
func (s *BaseYarnSpinnerParserListener) ExitLine_statement(ctx *Line_statementContext) {}

// EnterLine_formatted_text is called when production line_formatted_text is entered.
func (s *BaseYarnSpinnerParserListener) EnterLine_formatted_text(ctx *Line_formatted_textContext) {}

// ExitLine_formatted_text is called when production line_formatted_text is exited.
func (s *BaseYarnSpinnerParserListener) ExitLine_formatted_text(ctx *Line_formatted_textContext) {}

// EnterHashtag is called when production hashtag is entered.
func (s *BaseYarnSpinnerParserListener) EnterHashtag(ctx *HashtagContext) {}

// ExitHashtag is called when production hashtag is exited.
func (s *BaseYarnSpinnerParserListener) ExitHashtag(ctx *HashtagContext) {}

// EnterLine_condition is called when production line_condition is entered.
func (s *BaseYarnSpinnerParserListener) EnterLine_condition(ctx *Line_conditionContext) {}

// ExitLine_condition is called when production line_condition is exited.
func (s *BaseYarnSpinnerParserListener) ExitLine_condition(ctx *Line_conditionContext) {}

// EnterExpParens is called when production expParens is entered.
func (s *BaseYarnSpinnerParserListener) EnterExpParens(ctx *ExpParensContext) {}

// ExitExpParens is called when production expParens is exited.
func (s *BaseYarnSpinnerParserListener) ExitExpParens(ctx *ExpParensContext) {}

// EnterExpMultDivMod is called when production expMultDivMod is entered.
func (s *BaseYarnSpinnerParserListener) EnterExpMultDivMod(ctx *ExpMultDivModContext) {}

// ExitExpMultDivMod is called when production expMultDivMod is exited.
func (s *BaseYarnSpinnerParserListener) ExitExpMultDivMod(ctx *ExpMultDivModContext) {}

// EnterExpComparison is called when production expComparison is entered.
func (s *BaseYarnSpinnerParserListener) EnterExpComparison(ctx *ExpComparisonContext) {}

// ExitExpComparison is called when production expComparison is exited.
func (s *BaseYarnSpinnerParserListener) ExitExpComparison(ctx *ExpComparisonContext) {}

// EnterExpNegative is called when production expNegative is entered.
func (s *BaseYarnSpinnerParserListener) EnterExpNegative(ctx *ExpNegativeContext) {}

// ExitExpNegative is called when production expNegative is exited.
func (s *BaseYarnSpinnerParserListener) ExitExpNegative(ctx *ExpNegativeContext) {}

// EnterExpAndOrXor is called when production expAndOrXor is entered.
func (s *BaseYarnSpinnerParserListener) EnterExpAndOrXor(ctx *ExpAndOrXorContext) {}

// ExitExpAndOrXor is called when production expAndOrXor is exited.
func (s *BaseYarnSpinnerParserListener) ExitExpAndOrXor(ctx *ExpAndOrXorContext) {}

// EnterExpAddSub is called when production expAddSub is entered.
func (s *BaseYarnSpinnerParserListener) EnterExpAddSub(ctx *ExpAddSubContext) {}

// ExitExpAddSub is called when production expAddSub is exited.
func (s *BaseYarnSpinnerParserListener) ExitExpAddSub(ctx *ExpAddSubContext) {}

// EnterExpNot is called when production expNot is entered.
func (s *BaseYarnSpinnerParserListener) EnterExpNot(ctx *ExpNotContext) {}

// ExitExpNot is called when production expNot is exited.
func (s *BaseYarnSpinnerParserListener) ExitExpNot(ctx *ExpNotContext) {}

// EnterExpValue is called when production expValue is entered.
func (s *BaseYarnSpinnerParserListener) EnterExpValue(ctx *ExpValueContext) {}

// ExitExpValue is called when production expValue is exited.
func (s *BaseYarnSpinnerParserListener) ExitExpValue(ctx *ExpValueContext) {}

// EnterExpEquality is called when production expEquality is entered.
func (s *BaseYarnSpinnerParserListener) EnterExpEquality(ctx *ExpEqualityContext) {}

// ExitExpEquality is called when production expEquality is exited.
func (s *BaseYarnSpinnerParserListener) ExitExpEquality(ctx *ExpEqualityContext) {}

// EnterValueNumber is called when production valueNumber is entered.
func (s *BaseYarnSpinnerParserListener) EnterValueNumber(ctx *ValueNumberContext) {}

// ExitValueNumber is called when production valueNumber is exited.
func (s *BaseYarnSpinnerParserListener) ExitValueNumber(ctx *ValueNumberContext) {}

// EnterValueTrue is called when production valueTrue is entered.
func (s *BaseYarnSpinnerParserListener) EnterValueTrue(ctx *ValueTrueContext) {}

// ExitValueTrue is called when production valueTrue is exited.
func (s *BaseYarnSpinnerParserListener) ExitValueTrue(ctx *ValueTrueContext) {}

// EnterValueFalse is called when production valueFalse is entered.
func (s *BaseYarnSpinnerParserListener) EnterValueFalse(ctx *ValueFalseContext) {}

// ExitValueFalse is called when production valueFalse is exited.
func (s *BaseYarnSpinnerParserListener) ExitValueFalse(ctx *ValueFalseContext) {}

// EnterValueVar is called when production valueVar is entered.
func (s *BaseYarnSpinnerParserListener) EnterValueVar(ctx *ValueVarContext) {}

// ExitValueVar is called when production valueVar is exited.
func (s *BaseYarnSpinnerParserListener) ExitValueVar(ctx *ValueVarContext) {}

// EnterValueString is called when production valueString is entered.
func (s *BaseYarnSpinnerParserListener) EnterValueString(ctx *ValueStringContext) {}

// ExitValueString is called when production valueString is exited.
func (s *BaseYarnSpinnerParserListener) ExitValueString(ctx *ValueStringContext) {}

// EnterValueNull is called when production valueNull is entered.
func (s *BaseYarnSpinnerParserListener) EnterValueNull(ctx *ValueNullContext) {}

// ExitValueNull is called when production valueNull is exited.
func (s *BaseYarnSpinnerParserListener) ExitValueNull(ctx *ValueNullContext) {}

// EnterValueFunc is called when production valueFunc is entered.
func (s *BaseYarnSpinnerParserListener) EnterValueFunc(ctx *ValueFuncContext) {}

// ExitValueFunc is called when production valueFunc is exited.
func (s *BaseYarnSpinnerParserListener) ExitValueFunc(ctx *ValueFuncContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaseYarnSpinnerParserListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaseYarnSpinnerParserListener) ExitVariable(ctx *VariableContext) {}

// EnterFunction_call is called when production function_call is entered.
func (s *BaseYarnSpinnerParserListener) EnterFunction_call(ctx *Function_callContext) {}

// ExitFunction_call is called when production function_call is exited.
func (s *BaseYarnSpinnerParserListener) ExitFunction_call(ctx *Function_callContext) {}

// EnterIf_statement is called when production if_statement is entered.
func (s *BaseYarnSpinnerParserListener) EnterIf_statement(ctx *If_statementContext) {}

// ExitIf_statement is called when production if_statement is exited.
func (s *BaseYarnSpinnerParserListener) ExitIf_statement(ctx *If_statementContext) {}

// EnterIf_clause is called when production if_clause is entered.
func (s *BaseYarnSpinnerParserListener) EnterIf_clause(ctx *If_clauseContext) {}

// ExitIf_clause is called when production if_clause is exited.
func (s *BaseYarnSpinnerParserListener) ExitIf_clause(ctx *If_clauseContext) {}

// EnterElse_if_clause is called when production else_if_clause is entered.
func (s *BaseYarnSpinnerParserListener) EnterElse_if_clause(ctx *Else_if_clauseContext) {}

// ExitElse_if_clause is called when production else_if_clause is exited.
func (s *BaseYarnSpinnerParserListener) ExitElse_if_clause(ctx *Else_if_clauseContext) {}

// EnterElse_clause is called when production else_clause is entered.
func (s *BaseYarnSpinnerParserListener) EnterElse_clause(ctx *Else_clauseContext) {}

// ExitElse_clause is called when production else_clause is exited.
func (s *BaseYarnSpinnerParserListener) ExitElse_clause(ctx *Else_clauseContext) {}

// EnterSet_statement is called when production set_statement is entered.
func (s *BaseYarnSpinnerParserListener) EnterSet_statement(ctx *Set_statementContext) {}

// ExitSet_statement is called when production set_statement is exited.
func (s *BaseYarnSpinnerParserListener) ExitSet_statement(ctx *Set_statementContext) {}

// EnterCall_statement is called when production call_statement is entered.
func (s *BaseYarnSpinnerParserListener) EnterCall_statement(ctx *Call_statementContext) {}

// ExitCall_statement is called when production call_statement is exited.
func (s *BaseYarnSpinnerParserListener) ExitCall_statement(ctx *Call_statementContext) {}

// EnterCommand_statement is called when production command_statement is entered.
func (s *BaseYarnSpinnerParserListener) EnterCommand_statement(ctx *Command_statementContext) {}

// ExitCommand_statement is called when production command_statement is exited.
func (s *BaseYarnSpinnerParserListener) ExitCommand_statement(ctx *Command_statementContext) {}

// EnterCommand_formatted_text is called when production command_formatted_text is entered.
func (s *BaseYarnSpinnerParserListener) EnterCommand_formatted_text(ctx *Command_formatted_textContext) {
}

// ExitCommand_formatted_text is called when production command_formatted_text is exited.
func (s *BaseYarnSpinnerParserListener) ExitCommand_formatted_text(ctx *Command_formatted_textContext) {
}

// EnterShortcut_option_statement is called when production shortcut_option_statement is entered.
func (s *BaseYarnSpinnerParserListener) EnterShortcut_option_statement(ctx *Shortcut_option_statementContext) {
}

// ExitShortcut_option_statement is called when production shortcut_option_statement is exited.
func (s *BaseYarnSpinnerParserListener) ExitShortcut_option_statement(ctx *Shortcut_option_statementContext) {
}

// EnterShortcut_option is called when production shortcut_option is entered.
func (s *BaseYarnSpinnerParserListener) EnterShortcut_option(ctx *Shortcut_optionContext) {}

// ExitShortcut_option is called when production shortcut_option is exited.
func (s *BaseYarnSpinnerParserListener) ExitShortcut_option(ctx *Shortcut_optionContext) {}

// EnterDeclare_statement is called when production declare_statement is entered.
func (s *BaseYarnSpinnerParserListener) EnterDeclare_statement(ctx *Declare_statementContext) {}

// ExitDeclare_statement is called when production declare_statement is exited.
func (s *BaseYarnSpinnerParserListener) ExitDeclare_statement(ctx *Declare_statementContext) {}

// EnterJumpToNodeName is called when production jumpToNodeName is entered.
func (s *BaseYarnSpinnerParserListener) EnterJumpToNodeName(ctx *JumpToNodeNameContext) {}

// ExitJumpToNodeName is called when production jumpToNodeName is exited.
func (s *BaseYarnSpinnerParserListener) ExitJumpToNodeName(ctx *JumpToNodeNameContext) {}

// EnterJumpToExpression is called when production jumpToExpression is entered.
func (s *BaseYarnSpinnerParserListener) EnterJumpToExpression(ctx *JumpToExpressionContext) {}

// ExitJumpToExpression is called when production jumpToExpression is exited.
func (s *BaseYarnSpinnerParserListener) ExitJumpToExpression(ctx *JumpToExpressionContext) {}
