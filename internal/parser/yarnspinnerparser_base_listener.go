// Code generated from internal/parser/YarnSpinnerParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // YarnSpinnerParser

import "github.com/antlr4-go/antlr/v4"

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

// EnterTitle_header is called when production title_header is entered.
func (s *BaseYarnSpinnerParserListener) EnterTitle_header(ctx *Title_headerContext) {}

// ExitTitle_header is called when production title_header is exited.
func (s *BaseYarnSpinnerParserListener) ExitTitle_header(ctx *Title_headerContext) {}

// EnterWhen_header is called when production when_header is entered.
func (s *BaseYarnSpinnerParserListener) EnterWhen_header(ctx *When_headerContext) {}

// ExitWhen_header is called when production when_header is exited.
func (s *BaseYarnSpinnerParserListener) ExitWhen_header(ctx *When_headerContext) {}

// EnterHeader is called when production header is entered.
func (s *BaseYarnSpinnerParserListener) EnterHeader(ctx *HeaderContext) {}

// ExitHeader is called when production header is exited.
func (s *BaseYarnSpinnerParserListener) ExitHeader(ctx *HeaderContext) {}

// EnterHeader_when_expression is called when production header_when_expression is entered.
func (s *BaseYarnSpinnerParserListener) EnterHeader_when_expression(ctx *Header_when_expressionContext) {
}

// ExitHeader_when_expression is called when production header_when_expression is exited.
func (s *BaseYarnSpinnerParserListener) ExitHeader_when_expression(ctx *Header_when_expressionContext) {
}

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

// EnterLineCondition is called when production lineCondition is entered.
func (s *BaseYarnSpinnerParserListener) EnterLineCondition(ctx *LineConditionContext) {}

// ExitLineCondition is called when production lineCondition is exited.
func (s *BaseYarnSpinnerParserListener) ExitLineCondition(ctx *LineConditionContext) {}

// EnterLineOnceCondition is called when production lineOnceCondition is entered.
func (s *BaseYarnSpinnerParserListener) EnterLineOnceCondition(ctx *LineOnceConditionContext) {}

// ExitLineOnceCondition is called when production lineOnceCondition is exited.
func (s *BaseYarnSpinnerParserListener) ExitLineOnceCondition(ctx *LineOnceConditionContext) {}

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

// EnterValueFunc is called when production valueFunc is entered.
func (s *BaseYarnSpinnerParserListener) EnterValueFunc(ctx *ValueFuncContext) {}

// ExitValueFunc is called when production valueFunc is exited.
func (s *BaseYarnSpinnerParserListener) ExitValueFunc(ctx *ValueFuncContext) {}

// EnterValueTypeMemberReference is called when production valueTypeMemberReference is entered.
func (s *BaseYarnSpinnerParserListener) EnterValueTypeMemberReference(ctx *ValueTypeMemberReferenceContext) {
}

// ExitValueTypeMemberReference is called when production valueTypeMemberReference is exited.
func (s *BaseYarnSpinnerParserListener) ExitValueTypeMemberReference(ctx *ValueTypeMemberReferenceContext) {
}

// EnterVariable is called when production variable is entered.
func (s *BaseYarnSpinnerParserListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaseYarnSpinnerParserListener) ExitVariable(ctx *VariableContext) {}

// EnterFunction_call is called when production function_call is entered.
func (s *BaseYarnSpinnerParserListener) EnterFunction_call(ctx *Function_callContext) {}

// ExitFunction_call is called when production function_call is exited.
func (s *BaseYarnSpinnerParserListener) ExitFunction_call(ctx *Function_callContext) {}

// EnterTypeMemberReference is called when production typeMemberReference is entered.
func (s *BaseYarnSpinnerParserListener) EnterTypeMemberReference(ctx *TypeMemberReferenceContext) {}

// ExitTypeMemberReference is called when production typeMemberReference is exited.
func (s *BaseYarnSpinnerParserListener) ExitTypeMemberReference(ctx *TypeMemberReferenceContext) {}

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

// EnterLine_group_statement is called when production line_group_statement is entered.
func (s *BaseYarnSpinnerParserListener) EnterLine_group_statement(ctx *Line_group_statementContext) {}

// ExitLine_group_statement is called when production line_group_statement is exited.
func (s *BaseYarnSpinnerParserListener) ExitLine_group_statement(ctx *Line_group_statementContext) {}

// EnterLine_group_item is called when production line_group_item is entered.
func (s *BaseYarnSpinnerParserListener) EnterLine_group_item(ctx *Line_group_itemContext) {}

// ExitLine_group_item is called when production line_group_item is exited.
func (s *BaseYarnSpinnerParserListener) ExitLine_group_item(ctx *Line_group_itemContext) {}

// EnterDeclare_statement is called when production declare_statement is entered.
func (s *BaseYarnSpinnerParserListener) EnterDeclare_statement(ctx *Declare_statementContext) {}

// ExitDeclare_statement is called when production declare_statement is exited.
func (s *BaseYarnSpinnerParserListener) ExitDeclare_statement(ctx *Declare_statementContext) {}

// EnterEnum_statement is called when production enum_statement is entered.
func (s *BaseYarnSpinnerParserListener) EnterEnum_statement(ctx *Enum_statementContext) {}

// ExitEnum_statement is called when production enum_statement is exited.
func (s *BaseYarnSpinnerParserListener) ExitEnum_statement(ctx *Enum_statementContext) {}

// EnterEnum_case_statement is called when production enum_case_statement is entered.
func (s *BaseYarnSpinnerParserListener) EnterEnum_case_statement(ctx *Enum_case_statementContext) {}

// ExitEnum_case_statement is called when production enum_case_statement is exited.
func (s *BaseYarnSpinnerParserListener) ExitEnum_case_statement(ctx *Enum_case_statementContext) {}

// EnterJumpToNodeName is called when production jumpToNodeName is entered.
func (s *BaseYarnSpinnerParserListener) EnterJumpToNodeName(ctx *JumpToNodeNameContext) {}

// ExitJumpToNodeName is called when production jumpToNodeName is exited.
func (s *BaseYarnSpinnerParserListener) ExitJumpToNodeName(ctx *JumpToNodeNameContext) {}

// EnterJumpToExpression is called when production jumpToExpression is entered.
func (s *BaseYarnSpinnerParserListener) EnterJumpToExpression(ctx *JumpToExpressionContext) {}

// ExitJumpToExpression is called when production jumpToExpression is exited.
func (s *BaseYarnSpinnerParserListener) ExitJumpToExpression(ctx *JumpToExpressionContext) {}

// EnterDetourToNodeName is called when production detourToNodeName is entered.
func (s *BaseYarnSpinnerParserListener) EnterDetourToNodeName(ctx *DetourToNodeNameContext) {}

// ExitDetourToNodeName is called when production detourToNodeName is exited.
func (s *BaseYarnSpinnerParserListener) ExitDetourToNodeName(ctx *DetourToNodeNameContext) {}

// EnterDetourToExpression is called when production detourToExpression is entered.
func (s *BaseYarnSpinnerParserListener) EnterDetourToExpression(ctx *DetourToExpressionContext) {}

// ExitDetourToExpression is called when production detourToExpression is exited.
func (s *BaseYarnSpinnerParserListener) ExitDetourToExpression(ctx *DetourToExpressionContext) {}

// EnterReturn_statement is called when production return_statement is entered.
func (s *BaseYarnSpinnerParserListener) EnterReturn_statement(ctx *Return_statementContext) {}

// ExitReturn_statement is called when production return_statement is exited.
func (s *BaseYarnSpinnerParserListener) ExitReturn_statement(ctx *Return_statementContext) {}

// EnterOnce_statement is called when production once_statement is entered.
func (s *BaseYarnSpinnerParserListener) EnterOnce_statement(ctx *Once_statementContext) {}

// ExitOnce_statement is called when production once_statement is exited.
func (s *BaseYarnSpinnerParserListener) ExitOnce_statement(ctx *Once_statementContext) {}

// EnterOnce_primary_clause is called when production once_primary_clause is entered.
func (s *BaseYarnSpinnerParserListener) EnterOnce_primary_clause(ctx *Once_primary_clauseContext) {}

// ExitOnce_primary_clause is called when production once_primary_clause is exited.
func (s *BaseYarnSpinnerParserListener) ExitOnce_primary_clause(ctx *Once_primary_clauseContext) {}

// EnterOnce_alternate_clause is called when production once_alternate_clause is entered.
func (s *BaseYarnSpinnerParserListener) EnterOnce_alternate_clause(ctx *Once_alternate_clauseContext) {
}

// ExitOnce_alternate_clause is called when production once_alternate_clause is exited.
func (s *BaseYarnSpinnerParserListener) ExitOnce_alternate_clause(ctx *Once_alternate_clauseContext) {
}

// EnterStructured_command is called when production structured_command is entered.
func (s *BaseYarnSpinnerParserListener) EnterStructured_command(ctx *Structured_commandContext) {}

// ExitStructured_command is called when production structured_command is exited.
func (s *BaseYarnSpinnerParserListener) ExitStructured_command(ctx *Structured_commandContext) {}

// EnterStructured_command_value is called when production structured_command_value is entered.
func (s *BaseYarnSpinnerParserListener) EnterStructured_command_value(ctx *Structured_command_valueContext) {
}

// ExitStructured_command_value is called when production structured_command_value is exited.
func (s *BaseYarnSpinnerParserListener) ExitStructured_command_value(ctx *Structured_command_valueContext) {
}
