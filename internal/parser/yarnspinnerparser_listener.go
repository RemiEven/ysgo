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

	// EnterHeader is called when entering the header production.
	EnterHeader(c *HeaderContext)

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

	// EnterLine_condition is called when entering the line_condition production.
	EnterLine_condition(c *Line_conditionContext)

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

	// EnterValueNull is called when entering the valueNull production.
	EnterValueNull(c *ValueNullContext)

	// EnterValueFunc is called when entering the valueFunc production.
	EnterValueFunc(c *ValueFuncContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterFunction_call is called when entering the function_call production.
	EnterFunction_call(c *Function_callContext)

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

	// EnterDeclare_statement is called when entering the declare_statement production.
	EnterDeclare_statement(c *Declare_statementContext)

	// EnterJumpToNodeName is called when entering the jumpToNodeName production.
	EnterJumpToNodeName(c *JumpToNodeNameContext)

	// EnterJumpToExpression is called when entering the jumpToExpression production.
	EnterJumpToExpression(c *JumpToExpressionContext)

	// ExitDialogue is called when exiting the dialogue production.
	ExitDialogue(c *DialogueContext)

	// ExitFile_hashtag is called when exiting the file_hashtag production.
	ExitFile_hashtag(c *File_hashtagContext)

	// ExitNode is called when exiting the node production.
	ExitNode(c *NodeContext)

	// ExitHeader is called when exiting the header production.
	ExitHeader(c *HeaderContext)

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

	// ExitLine_condition is called when exiting the line_condition production.
	ExitLine_condition(c *Line_conditionContext)

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

	// ExitValueNull is called when exiting the valueNull production.
	ExitValueNull(c *ValueNullContext)

	// ExitValueFunc is called when exiting the valueFunc production.
	ExitValueFunc(c *ValueFuncContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitFunction_call is called when exiting the function_call production.
	ExitFunction_call(c *Function_callContext)

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

	// ExitDeclare_statement is called when exiting the declare_statement production.
	ExitDeclare_statement(c *Declare_statementContext)

	// ExitJumpToNodeName is called when exiting the jumpToNodeName production.
	ExitJumpToNodeName(c *JumpToNodeNameContext)

	// ExitJumpToExpression is called when exiting the jumpToExpression production.
	ExitJumpToExpression(c *JumpToExpressionContext)
}
