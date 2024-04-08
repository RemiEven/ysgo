// Code generated from internal/parser/YarnSpinnerParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // YarnSpinnerParser

import "github.com/antlr4-go/antlr/v4"

// YarnSpinnerParserListener is a complete listener for a parse tree produced by YarnSpinnerParser.
type YarnSpinnerParserListener interface {
	antlr.ParseTreeListener

	// EnterDialogue is called when entering the dialogue production.
	EnterDialogue(c *DialogueContext)

	// EnterFile_hashtag is called when entering the file_hashtag production.
	EnterFile_hashtag(c *File_hashtagContext)

	// EnterNode is called when entering the node production.
	EnterNode(c *NodeContext)

	// EnterTitle_header is called when entering the title_header production.
	EnterTitle_header(c *Title_headerContext)

	// EnterWhen_header is called when entering the when_header production.
	EnterWhen_header(c *When_headerContext)

	// EnterHeader is called when entering the header production.
	EnterHeader(c *HeaderContext)

	// EnterHeader_when_expression is called when entering the header_when_expression production.
	EnterHeader_when_expression(c *Header_when_expressionContext)

	// EnterBody is called when entering the body production.
	EnterBody(c *BodyContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterLine_statement is called when entering the line_statement production.
	EnterLine_statement(c *Line_statementContext)

	// EnterLine_formatted_text is called when entering the line_formatted_text production.
	EnterLine_formatted_text(c *Line_formatted_textContext)

	// EnterHashtag is called when entering the hashtag production.
	EnterHashtag(c *HashtagContext)

	// EnterLineCondition is called when entering the lineCondition production.
	EnterLineCondition(c *LineConditionContext)

	// EnterLineOnceCondition is called when entering the lineOnceCondition production.
	EnterLineOnceCondition(c *LineOnceConditionContext)

	// EnterExpParens is called when entering the expParens production.
	EnterExpParens(c *ExpParensContext)

	// EnterExpMultDivMod is called when entering the expMultDivMod production.
	EnterExpMultDivMod(c *ExpMultDivModContext)

	// EnterExpComparison is called when entering the expComparison production.
	EnterExpComparison(c *ExpComparisonContext)

	// EnterExpNegative is called when entering the expNegative production.
	EnterExpNegative(c *ExpNegativeContext)

	// EnterExpAndOrXor is called when entering the expAndOrXor production.
	EnterExpAndOrXor(c *ExpAndOrXorContext)

	// EnterExpAddSub is called when entering the expAddSub production.
	EnterExpAddSub(c *ExpAddSubContext)

	// EnterExpNot is called when entering the expNot production.
	EnterExpNot(c *ExpNotContext)

	// EnterExpValue is called when entering the expValue production.
	EnterExpValue(c *ExpValueContext)

	// EnterExpEquality is called when entering the expEquality production.
	EnterExpEquality(c *ExpEqualityContext)

	// EnterValueNumber is called when entering the valueNumber production.
	EnterValueNumber(c *ValueNumberContext)

	// EnterValueTrue is called when entering the valueTrue production.
	EnterValueTrue(c *ValueTrueContext)

	// EnterValueFalse is called when entering the valueFalse production.
	EnterValueFalse(c *ValueFalseContext)

	// EnterValueVar is called when entering the valueVar production.
	EnterValueVar(c *ValueVarContext)

	// EnterValueString is called when entering the valueString production.
	EnterValueString(c *ValueStringContext)

	// EnterValueFunc is called when entering the valueFunc production.
	EnterValueFunc(c *ValueFuncContext)

	// EnterValueTypeMemberReference is called when entering the valueTypeMemberReference production.
	EnterValueTypeMemberReference(c *ValueTypeMemberReferenceContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterFunction_call is called when entering the function_call production.
	EnterFunction_call(c *Function_callContext)

	// EnterTypeMemberReference is called when entering the typeMemberReference production.
	EnterTypeMemberReference(c *TypeMemberReferenceContext)

	// EnterIf_statement is called when entering the if_statement production.
	EnterIf_statement(c *If_statementContext)

	// EnterIf_clause is called when entering the if_clause production.
	EnterIf_clause(c *If_clauseContext)

	// EnterElse_if_clause is called when entering the else_if_clause production.
	EnterElse_if_clause(c *Else_if_clauseContext)

	// EnterElse_clause is called when entering the else_clause production.
	EnterElse_clause(c *Else_clauseContext)

	// EnterSet_statement is called when entering the set_statement production.
	EnterSet_statement(c *Set_statementContext)

	// EnterCall_statement is called when entering the call_statement production.
	EnterCall_statement(c *Call_statementContext)

	// EnterCommand_statement is called when entering the command_statement production.
	EnterCommand_statement(c *Command_statementContext)

	// EnterCommand_formatted_text is called when entering the command_formatted_text production.
	EnterCommand_formatted_text(c *Command_formatted_textContext)

	// EnterShortcut_option_statement is called when entering the shortcut_option_statement production.
	EnterShortcut_option_statement(c *Shortcut_option_statementContext)

	// EnterShortcut_option is called when entering the shortcut_option production.
	EnterShortcut_option(c *Shortcut_optionContext)

	// EnterLine_group_statement is called when entering the line_group_statement production.
	EnterLine_group_statement(c *Line_group_statementContext)

	// EnterLine_group_item is called when entering the line_group_item production.
	EnterLine_group_item(c *Line_group_itemContext)

	// EnterDeclare_statement is called when entering the declare_statement production.
	EnterDeclare_statement(c *Declare_statementContext)

	// EnterEnum_statement is called when entering the enum_statement production.
	EnterEnum_statement(c *Enum_statementContext)

	// EnterEnum_case_statement is called when entering the enum_case_statement production.
	EnterEnum_case_statement(c *Enum_case_statementContext)

	// EnterJumpToNodeName is called when entering the jumpToNodeName production.
	EnterJumpToNodeName(c *JumpToNodeNameContext)

	// EnterJumpToExpression is called when entering the jumpToExpression production.
	EnterJumpToExpression(c *JumpToExpressionContext)

	// EnterDetourToNodeName is called when entering the detourToNodeName production.
	EnterDetourToNodeName(c *DetourToNodeNameContext)

	// EnterDetourToExpression is called when entering the detourToExpression production.
	EnterDetourToExpression(c *DetourToExpressionContext)

	// EnterReturn_statement is called when entering the return_statement production.
	EnterReturn_statement(c *Return_statementContext)

	// EnterOnce_statement is called when entering the once_statement production.
	EnterOnce_statement(c *Once_statementContext)

	// EnterOnce_primary_clause is called when entering the once_primary_clause production.
	EnterOnce_primary_clause(c *Once_primary_clauseContext)

	// EnterOnce_alternate_clause is called when entering the once_alternate_clause production.
	EnterOnce_alternate_clause(c *Once_alternate_clauseContext)

	// EnterStructured_command is called when entering the structured_command production.
	EnterStructured_command(c *Structured_commandContext)

	// EnterStructured_command_value is called when entering the structured_command_value production.
	EnterStructured_command_value(c *Structured_command_valueContext)

	// ExitDialogue is called when exiting the dialogue production.
	ExitDialogue(c *DialogueContext)

	// ExitFile_hashtag is called when exiting the file_hashtag production.
	ExitFile_hashtag(c *File_hashtagContext)

	// ExitNode is called when exiting the node production.
	ExitNode(c *NodeContext)

	// ExitTitle_header is called when exiting the title_header production.
	ExitTitle_header(c *Title_headerContext)

	// ExitWhen_header is called when exiting the when_header production.
	ExitWhen_header(c *When_headerContext)

	// ExitHeader is called when exiting the header production.
	ExitHeader(c *HeaderContext)

	// ExitHeader_when_expression is called when exiting the header_when_expression production.
	ExitHeader_when_expression(c *Header_when_expressionContext)

	// ExitBody is called when exiting the body production.
	ExitBody(c *BodyContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitLine_statement is called when exiting the line_statement production.
	ExitLine_statement(c *Line_statementContext)

	// ExitLine_formatted_text is called when exiting the line_formatted_text production.
	ExitLine_formatted_text(c *Line_formatted_textContext)

	// ExitHashtag is called when exiting the hashtag production.
	ExitHashtag(c *HashtagContext)

	// ExitLineCondition is called when exiting the lineCondition production.
	ExitLineCondition(c *LineConditionContext)

	// ExitLineOnceCondition is called when exiting the lineOnceCondition production.
	ExitLineOnceCondition(c *LineOnceConditionContext)

	// ExitExpParens is called when exiting the expParens production.
	ExitExpParens(c *ExpParensContext)

	// ExitExpMultDivMod is called when exiting the expMultDivMod production.
	ExitExpMultDivMod(c *ExpMultDivModContext)

	// ExitExpComparison is called when exiting the expComparison production.
	ExitExpComparison(c *ExpComparisonContext)

	// ExitExpNegative is called when exiting the expNegative production.
	ExitExpNegative(c *ExpNegativeContext)

	// ExitExpAndOrXor is called when exiting the expAndOrXor production.
	ExitExpAndOrXor(c *ExpAndOrXorContext)

	// ExitExpAddSub is called when exiting the expAddSub production.
	ExitExpAddSub(c *ExpAddSubContext)

	// ExitExpNot is called when exiting the expNot production.
	ExitExpNot(c *ExpNotContext)

	// ExitExpValue is called when exiting the expValue production.
	ExitExpValue(c *ExpValueContext)

	// ExitExpEquality is called when exiting the expEquality production.
	ExitExpEquality(c *ExpEqualityContext)

	// ExitValueNumber is called when exiting the valueNumber production.
	ExitValueNumber(c *ValueNumberContext)

	// ExitValueTrue is called when exiting the valueTrue production.
	ExitValueTrue(c *ValueTrueContext)

	// ExitValueFalse is called when exiting the valueFalse production.
	ExitValueFalse(c *ValueFalseContext)

	// ExitValueVar is called when exiting the valueVar production.
	ExitValueVar(c *ValueVarContext)

	// ExitValueString is called when exiting the valueString production.
	ExitValueString(c *ValueStringContext)

	// ExitValueFunc is called when exiting the valueFunc production.
	ExitValueFunc(c *ValueFuncContext)

	// ExitValueTypeMemberReference is called when exiting the valueTypeMemberReference production.
	ExitValueTypeMemberReference(c *ValueTypeMemberReferenceContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitFunction_call is called when exiting the function_call production.
	ExitFunction_call(c *Function_callContext)

	// ExitTypeMemberReference is called when exiting the typeMemberReference production.
	ExitTypeMemberReference(c *TypeMemberReferenceContext)

	// ExitIf_statement is called when exiting the if_statement production.
	ExitIf_statement(c *If_statementContext)

	// ExitIf_clause is called when exiting the if_clause production.
	ExitIf_clause(c *If_clauseContext)

	// ExitElse_if_clause is called when exiting the else_if_clause production.
	ExitElse_if_clause(c *Else_if_clauseContext)

	// ExitElse_clause is called when exiting the else_clause production.
	ExitElse_clause(c *Else_clauseContext)

	// ExitSet_statement is called when exiting the set_statement production.
	ExitSet_statement(c *Set_statementContext)

	// ExitCall_statement is called when exiting the call_statement production.
	ExitCall_statement(c *Call_statementContext)

	// ExitCommand_statement is called when exiting the command_statement production.
	ExitCommand_statement(c *Command_statementContext)

	// ExitCommand_formatted_text is called when exiting the command_formatted_text production.
	ExitCommand_formatted_text(c *Command_formatted_textContext)

	// ExitShortcut_option_statement is called when exiting the shortcut_option_statement production.
	ExitShortcut_option_statement(c *Shortcut_option_statementContext)

	// ExitShortcut_option is called when exiting the shortcut_option production.
	ExitShortcut_option(c *Shortcut_optionContext)

	// ExitLine_group_statement is called when exiting the line_group_statement production.
	ExitLine_group_statement(c *Line_group_statementContext)

	// ExitLine_group_item is called when exiting the line_group_item production.
	ExitLine_group_item(c *Line_group_itemContext)

	// ExitDeclare_statement is called when exiting the declare_statement production.
	ExitDeclare_statement(c *Declare_statementContext)

	// ExitEnum_statement is called when exiting the enum_statement production.
	ExitEnum_statement(c *Enum_statementContext)

	// ExitEnum_case_statement is called when exiting the enum_case_statement production.
	ExitEnum_case_statement(c *Enum_case_statementContext)

	// ExitJumpToNodeName is called when exiting the jumpToNodeName production.
	ExitJumpToNodeName(c *JumpToNodeNameContext)

	// ExitJumpToExpression is called when exiting the jumpToExpression production.
	ExitJumpToExpression(c *JumpToExpressionContext)

	// ExitDetourToNodeName is called when exiting the detourToNodeName production.
	ExitDetourToNodeName(c *DetourToNodeNameContext)

	// ExitDetourToExpression is called when exiting the detourToExpression production.
	ExitDetourToExpression(c *DetourToExpressionContext)

	// ExitReturn_statement is called when exiting the return_statement production.
	ExitReturn_statement(c *Return_statementContext)

	// ExitOnce_statement is called when exiting the once_statement production.
	ExitOnce_statement(c *Once_statementContext)

	// ExitOnce_primary_clause is called when exiting the once_primary_clause production.
	ExitOnce_primary_clause(c *Once_primary_clauseContext)

	// ExitOnce_alternate_clause is called when exiting the once_alternate_clause production.
	ExitOnce_alternate_clause(c *Once_alternate_clauseContext)

	// ExitStructured_command is called when exiting the structured_command production.
	ExitStructured_command(c *Structured_commandContext)

	// ExitStructured_command_value is called when exiting the structured_command_value production.
	ExitStructured_command_value(c *Structured_command_valueContext)
}
