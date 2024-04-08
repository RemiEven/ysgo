// Code generated from internal/parser/YarnSpinnerParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // YarnSpinnerParser

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type YarnSpinnerParser struct {
	*antlr.BaseParser
}

var YarnSpinnerParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func yarnspinnerparserParserInit() {
	staticData := &YarnSpinnerParserParserStaticData
	staticData.LiteralNames = []string{
		"", "", "", "", "", "", "", "'when'", "'title'", "", "'---'", "", "'#'",
		"", "", "", "'==='", "'->'", "'=>'", "'<<'", "", "", "", "", "", "",
		"", "", "", "", "", "", "'always'", "'true'", "'false'", "'null'", "",
		"", "", "", "", "", "", "", "", "", "", "'+='", "'-='", "'*='", "'%='",
		"'/='", "'+'", "'-'", "'*'", "'/'", "'%'", "'('", "')'", "','", "'as'",
		"", "", "'}'", "", "'.'", "", "", "", "", "", "", "", "'endif'", "",
		"", "", "", "", "", "", "", "'once'", "'endonce'", "", "", "", "", "'{'",
		"", "", "'string'", "'number'", "'bool'",
	}
	staticData.SymbolicNames = []string{
		"", "INDENT", "DEDENT", "BLANK_LINE_FOLLOWING_OPTION", "WS", "COMMENT",
		"NEWLINE", "HEADER_WHEN", "HEADER_TITLE", "ID", "BODY_START", "HEADER_DELIMITER",
		"HASHTAG", "HEADER_WHEN_UNKNOWN", "REST_OF_LINE", "BODY_WS", "BODY_END",
		"SHORTCUT_ARROW", "LINE_GROUP_ARROW", "COMMAND_START", "EXPRESSION_START",
		"ESCAPED_ANY", "TEXT_ESCAPE", "TEXT_COMMENT", "TEXT", "UNESCAPABLE_CHARACTER",
		"TEXT_COMMANDHASHTAG_WS", "TEXT_COMMANDHASHTAG_COMMENT", "TEXT_COMMANDHASHTAG_ERROR",
		"HASHTAG_WS", "HASHTAG_TEXT", "EXPR_WS", "EXPRESSION_WHEN_ALWAYS", "KEYWORD_TRUE",
		"KEYWORD_FALSE", "KEYWORD_NULL", "OPERATOR_ASSIGNMENT", "OPERATOR_LOGICAL_LESS_THAN_EQUALS",
		"OPERATOR_LOGICAL_GREATER_THAN_EQUALS", "OPERATOR_LOGICAL_EQUALS", "OPERATOR_LOGICAL_LESS",
		"OPERATOR_LOGICAL_GREATER", "OPERATOR_LOGICAL_NOT_EQUALS", "OPERATOR_LOGICAL_AND",
		"OPERATOR_LOGICAL_OR", "OPERATOR_LOGICAL_XOR", "OPERATOR_LOGICAL_NOT",
		"OPERATOR_MATHS_ADDITION_EQUALS", "OPERATOR_MATHS_SUBTRACTION_EQUALS",
		"OPERATOR_MATHS_MULTIPLICATION_EQUALS", "OPERATOR_MATHS_MODULUS_EQUALS",
		"OPERATOR_MATHS_DIVISION_EQUALS", "OPERATOR_MATHS_ADDITION", "OPERATOR_MATHS_SUBTRACTION",
		"OPERATOR_MATHS_MULTIPLICATION", "OPERATOR_MATHS_DIVISION", "OPERATOR_MATHS_MODULUS",
		"LPAREN", "RPAREN", "COMMA", "EXPRESSION_AS", "STRING", "FUNC_ID", "EXPRESSION_END",
		"VAR_ID", "DOT", "NUMBER", "COMMAND_NEWLINE", "COMMAND_WS", "COMMAND_IF",
		"COMMAND_ELSEIF", "COMMAND_ELSE", "COMMAND_SET", "COMMAND_ENDIF", "COMMAND_CALL",
		"COMMAND_DECLARE", "COMMAND_JUMP", "COMMAND_DETOUR", "COMMAND_RETURN",
		"COMMAND_ENUM", "COMMAND_CASE", "COMMAND_ENDENUM", "COMMAND_ONCE", "COMMAND_ENDONCE",
		"COMMAND_LOCAL", "COMMAND_END", "COMMAND_TEXT_NEWLINE", "COMMAND_TEXT_END",
		"COMMAND_EXPRESSION_START", "COMMAND_TEXT", "COMMAND_ID_NEWLINE", "TYPE_STRING",
		"TYPE_NUMBER", "TYPE_BOOL",
	}
	staticData.RuleNames = []string{
		"dialogue", "file_hashtag", "node", "title_header", "when_header", "header",
		"header_when_expression", "body", "statement", "line_statement", "line_formatted_text",
		"hashtag", "line_condition", "expression", "value", "variable", "function_call",
		"typeMemberReference", "if_statement", "if_clause", "else_if_clause",
		"else_clause", "set_statement", "call_statement", "command_statement",
		"command_formatted_text", "shortcut_option_statement", "shortcut_option",
		"line_group_statement", "line_group_item", "declare_statement", "enum_statement",
		"enum_case_statement", "jump_statement", "return_statement", "once_statement",
		"once_primary_clause", "once_alternate_clause", "structured_command",
		"structured_command_value",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 93, 488, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 1, 0, 5, 0, 82, 8, 0, 10, 0,
		12, 0, 85, 9, 0, 1, 0, 4, 0, 88, 8, 0, 11, 0, 12, 0, 89, 1, 1, 1, 1, 1,
		1, 1, 2, 1, 2, 1, 2, 4, 2, 98, 8, 2, 11, 2, 12, 2, 99, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5,
		1, 5, 1, 5, 3, 5, 119, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 126, 8,
		6, 3, 6, 128, 8, 6, 1, 7, 5, 7, 131, 8, 7, 10, 7, 12, 7, 134, 9, 7, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 5, 8, 150, 8, 8, 10, 8, 12, 8, 153, 9, 8, 1, 8, 3, 8, 156, 8,
		8, 1, 9, 1, 9, 3, 9, 160, 8, 9, 1, 9, 5, 9, 163, 8, 9, 10, 9, 12, 9, 166,
		9, 9, 1, 9, 1, 9, 1, 10, 4, 10, 171, 8, 10, 11, 10, 12, 10, 172, 1, 10,
		1, 10, 1, 10, 1, 10, 4, 10, 179, 8, 10, 11, 10, 12, 10, 180, 1, 11, 1,
		11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		3, 12, 195, 8, 12, 1, 12, 3, 12, 198, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 210, 8, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 5, 13, 227, 8, 13, 10, 13, 12, 13, 230, 9, 13, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 239, 8, 14, 1, 15,
		1, 15, 1, 16, 1, 16, 1, 16, 3, 16, 246, 8, 16, 1, 16, 1, 16, 5, 16, 250,
		8, 16, 10, 16, 12, 16, 253, 9, 16, 1, 16, 1, 16, 1, 17, 3, 17, 258, 8,
		17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 5, 18, 265, 8, 18, 10, 18, 12, 18,
		268, 9, 18, 1, 18, 3, 18, 271, 8, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 19, 5, 19, 282, 8, 19, 10, 19, 12, 19, 285, 9,
		19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 5, 20, 292, 8, 20, 10, 20, 12, 20,
		295, 9, 20, 1, 21, 1, 21, 1, 21, 1, 21, 5, 21, 301, 8, 21, 10, 21, 12,
		21, 304, 9, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23,
		1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 5, 24, 322, 8,
		24, 10, 24, 12, 24, 325, 9, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 5, 25,
		332, 8, 25, 10, 25, 12, 25, 335, 9, 25, 1, 26, 5, 26, 338, 8, 26, 10, 26,
		12, 26, 341, 9, 26, 1, 26, 1, 26, 3, 26, 345, 8, 26, 1, 27, 1, 27, 1, 27,
		1, 27, 5, 27, 351, 8, 27, 10, 27, 12, 27, 354, 9, 27, 1, 27, 3, 27, 357,
		8, 27, 1, 28, 5, 28, 360, 8, 28, 10, 28, 12, 28, 363, 9, 28, 1, 28, 1,
		28, 3, 28, 367, 8, 28, 1, 29, 1, 29, 1, 29, 1, 29, 5, 29, 373, 8, 29, 10,
		29, 12, 29, 376, 9, 29, 1, 29, 3, 29, 379, 8, 29, 1, 30, 1, 30, 1, 30,
		1, 30, 1, 30, 1, 30, 1, 30, 3, 30, 388, 8, 30, 1, 30, 1, 30, 1, 31, 1,
		31, 1, 31, 1, 31, 1, 31, 4, 31, 397, 8, 31, 11, 31, 12, 31, 398, 1, 31,
		1, 31, 1, 31, 1, 31, 1, 32, 3, 32, 406, 8, 32, 1, 32, 1, 32, 1, 32, 1,
		32, 1, 32, 3, 32, 413, 8, 32, 1, 32, 1, 32, 3, 32, 417, 8, 32, 1, 33, 1,
		33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33,
		1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 3,
		33, 441, 8, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 3, 35, 449, 8,
		35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 36, 3, 36, 459,
		8, 36, 1, 36, 1, 36, 5, 36, 463, 8, 36, 10, 36, 12, 36, 466, 9, 36, 1,
		37, 1, 37, 1, 37, 1, 37, 5, 37, 472, 8, 37, 10, 37, 12, 37, 475, 9, 37,
		1, 38, 1, 38, 5, 38, 479, 8, 38, 10, 38, 12, 38, 482, 9, 38, 1, 39, 1,
		39, 3, 39, 486, 8, 39, 1, 39, 0, 1, 26, 40, 0, 2, 4, 6, 8, 10, 12, 14,
		16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50,
		52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 0, 6, 1, 0, 54,
		56, 1, 0, 52, 53, 2, 0, 37, 38, 40, 41, 2, 0, 39, 39, 42, 42, 1, 0, 43,
		45, 2, 0, 36, 36, 47, 51, 524, 0, 83, 1, 0, 0, 0, 2, 91, 1, 0, 0, 0, 4,
		97, 1, 0, 0, 0, 6, 105, 1, 0, 0, 0, 8, 110, 1, 0, 0, 0, 10, 115, 1, 0,
		0, 0, 12, 127, 1, 0, 0, 0, 14, 132, 1, 0, 0, 0, 16, 155, 1, 0, 0, 0, 18,
		157, 1, 0, 0, 0, 20, 178, 1, 0, 0, 0, 22, 182, 1, 0, 0, 0, 24, 197, 1,
		0, 0, 0, 26, 209, 1, 0, 0, 0, 28, 238, 1, 0, 0, 0, 30, 240, 1, 0, 0, 0,
		32, 242, 1, 0, 0, 0, 34, 257, 1, 0, 0, 0, 36, 262, 1, 0, 0, 0, 38, 276,
		1, 0, 0, 0, 40, 286, 1, 0, 0, 0, 42, 296, 1, 0, 0, 0, 44, 305, 1, 0, 0,
		0, 46, 312, 1, 0, 0, 0, 48, 317, 1, 0, 0, 0, 50, 333, 1, 0, 0, 0, 52, 339,
		1, 0, 0, 0, 54, 346, 1, 0, 0, 0, 56, 361, 1, 0, 0, 0, 58, 368, 1, 0, 0,
		0, 60, 380, 1, 0, 0, 0, 62, 391, 1, 0, 0, 0, 64, 405, 1, 0, 0, 0, 66, 440,
		1, 0, 0, 0, 68, 442, 1, 0, 0, 0, 70, 446, 1, 0, 0, 0, 72, 454, 1, 0, 0,
		0, 74, 467, 1, 0, 0, 0, 76, 476, 1, 0, 0, 0, 78, 485, 1, 0, 0, 0, 80, 82,
		3, 2, 1, 0, 81, 80, 1, 0, 0, 0, 82, 85, 1, 0, 0, 0, 83, 81, 1, 0, 0, 0,
		83, 84, 1, 0, 0, 0, 84, 87, 1, 0, 0, 0, 85, 83, 1, 0, 0, 0, 86, 88, 3,
		4, 2, 0, 87, 86, 1, 0, 0, 0, 88, 89, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 89,
		90, 1, 0, 0, 0, 90, 1, 1, 0, 0, 0, 91, 92, 5, 12, 0, 0, 92, 93, 5, 30,
		0, 0, 93, 3, 1, 0, 0, 0, 94, 98, 3, 10, 5, 0, 95, 98, 3, 8, 4, 0, 96, 98,
		3, 6, 3, 0, 97, 94, 1, 0, 0, 0, 97, 95, 1, 0, 0, 0, 97, 96, 1, 0, 0, 0,
		98, 99, 1, 0, 0, 0, 99, 97, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0, 100, 101,
		1, 0, 0, 0, 101, 102, 5, 10, 0, 0, 102, 103, 3, 14, 7, 0, 103, 104, 5,
		16, 0, 0, 104, 5, 1, 0, 0, 0, 105, 106, 5, 8, 0, 0, 106, 107, 5, 11, 0,
		0, 107, 108, 5, 9, 0, 0, 108, 109, 5, 6, 0, 0, 109, 7, 1, 0, 0, 0, 110,
		111, 5, 7, 0, 0, 111, 112, 5, 11, 0, 0, 112, 113, 3, 12, 6, 0, 113, 114,
		5, 6, 0, 0, 114, 9, 1, 0, 0, 0, 115, 116, 5, 9, 0, 0, 116, 118, 5, 11,
		0, 0, 117, 119, 5, 14, 0, 0, 118, 117, 1, 0, 0, 0, 118, 119, 1, 0, 0, 0,
		119, 11, 1, 0, 0, 0, 120, 128, 3, 26, 13, 0, 121, 128, 5, 32, 0, 0, 122,
		125, 5, 82, 0, 0, 123, 124, 5, 69, 0, 0, 124, 126, 3, 26, 13, 0, 125, 123,
		1, 0, 0, 0, 125, 126, 1, 0, 0, 0, 126, 128, 1, 0, 0, 0, 127, 120, 1, 0,
		0, 0, 127, 121, 1, 0, 0, 0, 127, 122, 1, 0, 0, 0, 128, 13, 1, 0, 0, 0,
		129, 131, 3, 16, 8, 0, 130, 129, 1, 0, 0, 0, 131, 134, 1, 0, 0, 0, 132,
		130, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 15, 1, 0, 0, 0, 134, 132, 1,
		0, 0, 0, 135, 156, 3, 18, 9, 0, 136, 156, 3, 36, 18, 0, 137, 156, 3, 44,
		22, 0, 138, 156, 3, 52, 26, 0, 139, 156, 3, 46, 23, 0, 140, 156, 3, 48,
		24, 0, 141, 156, 3, 60, 30, 0, 142, 156, 3, 62, 31, 0, 143, 156, 3, 66,
		33, 0, 144, 156, 3, 68, 34, 0, 145, 156, 3, 56, 28, 0, 146, 156, 3, 70,
		35, 0, 147, 151, 5, 1, 0, 0, 148, 150, 3, 16, 8, 0, 149, 148, 1, 0, 0,
		0, 150, 153, 1, 0, 0, 0, 151, 149, 1, 0, 0, 0, 151, 152, 1, 0, 0, 0, 152,
		154, 1, 0, 0, 0, 153, 151, 1, 0, 0, 0, 154, 156, 5, 2, 0, 0, 155, 135,
		1, 0, 0, 0, 155, 136, 1, 0, 0, 0, 155, 137, 1, 0, 0, 0, 155, 138, 1, 0,
		0, 0, 155, 139, 1, 0, 0, 0, 155, 140, 1, 0, 0, 0, 155, 141, 1, 0, 0, 0,
		155, 142, 1, 0, 0, 0, 155, 143, 1, 0, 0, 0, 155, 144, 1, 0, 0, 0, 155,
		145, 1, 0, 0, 0, 155, 146, 1, 0, 0, 0, 155, 147, 1, 0, 0, 0, 156, 17, 1,
		0, 0, 0, 157, 159, 3, 20, 10, 0, 158, 160, 3, 24, 12, 0, 159, 158, 1, 0,
		0, 0, 159, 160, 1, 0, 0, 0, 160, 164, 1, 0, 0, 0, 161, 163, 3, 22, 11,
		0, 162, 161, 1, 0, 0, 0, 163, 166, 1, 0, 0, 0, 164, 162, 1, 0, 0, 0, 164,
		165, 1, 0, 0, 0, 165, 167, 1, 0, 0, 0, 166, 164, 1, 0, 0, 0, 167, 168,
		5, 6, 0, 0, 168, 19, 1, 0, 0, 0, 169, 171, 5, 24, 0, 0, 170, 169, 1, 0,
		0, 0, 171, 172, 1, 0, 0, 0, 172, 170, 1, 0, 0, 0, 172, 173, 1, 0, 0, 0,
		173, 179, 1, 0, 0, 0, 174, 175, 5, 20, 0, 0, 175, 176, 3, 26, 13, 0, 176,
		177, 5, 63, 0, 0, 177, 179, 1, 0, 0, 0, 178, 170, 1, 0, 0, 0, 178, 174,
		1, 0, 0, 0, 179, 180, 1, 0, 0, 0, 180, 178, 1, 0, 0, 0, 180, 181, 1, 0,
		0, 0, 181, 21, 1, 0, 0, 0, 182, 183, 5, 12, 0, 0, 183, 184, 5, 30, 0, 0,
		184, 23, 1, 0, 0, 0, 185, 186, 5, 19, 0, 0, 186, 187, 5, 69, 0, 0, 187,
		188, 3, 26, 13, 0, 188, 189, 5, 85, 0, 0, 189, 198, 1, 0, 0, 0, 190, 191,
		5, 19, 0, 0, 191, 194, 5, 82, 0, 0, 192, 193, 5, 69, 0, 0, 193, 195, 3,
		26, 13, 0, 194, 192, 1, 0, 0, 0, 194, 195, 1, 0, 0, 0, 195, 196, 1, 0,
		0, 0, 196, 198, 5, 85, 0, 0, 197, 185, 1, 0, 0, 0, 197, 190, 1, 0, 0, 0,
		198, 25, 1, 0, 0, 0, 199, 200, 6, 13, -1, 0, 200, 201, 5, 57, 0, 0, 201,
		202, 3, 26, 13, 0, 202, 203, 5, 58, 0, 0, 203, 210, 1, 0, 0, 0, 204, 205,
		5, 53, 0, 0, 205, 210, 3, 26, 13, 8, 206, 207, 5, 46, 0, 0, 207, 210, 3,
		26, 13, 7, 208, 210, 3, 28, 14, 0, 209, 199, 1, 0, 0, 0, 209, 204, 1, 0,
		0, 0, 209, 206, 1, 0, 0, 0, 209, 208, 1, 0, 0, 0, 210, 228, 1, 0, 0, 0,
		211, 212, 10, 6, 0, 0, 212, 213, 7, 0, 0, 0, 213, 227, 3, 26, 13, 7, 214,
		215, 10, 5, 0, 0, 215, 216, 7, 1, 0, 0, 216, 227, 3, 26, 13, 6, 217, 218,
		10, 4, 0, 0, 218, 219, 7, 2, 0, 0, 219, 227, 3, 26, 13, 5, 220, 221, 10,
		3, 0, 0, 221, 222, 7, 3, 0, 0, 222, 227, 3, 26, 13, 4, 223, 224, 10, 2,
		0, 0, 224, 225, 7, 4, 0, 0, 225, 227, 3, 26, 13, 3, 226, 211, 1, 0, 0,
		0, 226, 214, 1, 0, 0, 0, 226, 217, 1, 0, 0, 0, 226, 220, 1, 0, 0, 0, 226,
		223, 1, 0, 0, 0, 227, 230, 1, 0, 0, 0, 228, 226, 1, 0, 0, 0, 228, 229,
		1, 0, 0, 0, 229, 27, 1, 0, 0, 0, 230, 228, 1, 0, 0, 0, 231, 239, 5, 66,
		0, 0, 232, 239, 5, 33, 0, 0, 233, 239, 5, 34, 0, 0, 234, 239, 3, 30, 15,
		0, 235, 239, 5, 61, 0, 0, 236, 239, 3, 32, 16, 0, 237, 239, 3, 34, 17,
		0, 238, 231, 1, 0, 0, 0, 238, 232, 1, 0, 0, 0, 238, 233, 1, 0, 0, 0, 238,
		234, 1, 0, 0, 0, 238, 235, 1, 0, 0, 0, 238, 236, 1, 0, 0, 0, 238, 237,
		1, 0, 0, 0, 239, 29, 1, 0, 0, 0, 240, 241, 5, 64, 0, 0, 241, 31, 1, 0,
		0, 0, 242, 243, 5, 62, 0, 0, 243, 245, 5, 57, 0, 0, 244, 246, 3, 26, 13,
		0, 245, 244, 1, 0, 0, 0, 245, 246, 1, 0, 0, 0, 246, 251, 1, 0, 0, 0, 247,
		248, 5, 59, 0, 0, 248, 250, 3, 26, 13, 0, 249, 247, 1, 0, 0, 0, 250, 253,
		1, 0, 0, 0, 251, 249, 1, 0, 0, 0, 251, 252, 1, 0, 0, 0, 252, 254, 1, 0,
		0, 0, 253, 251, 1, 0, 0, 0, 254, 255, 5, 58, 0, 0, 255, 33, 1, 0, 0, 0,
		256, 258, 5, 62, 0, 0, 257, 256, 1, 0, 0, 0, 257, 258, 1, 0, 0, 0, 258,
		259, 1, 0, 0, 0, 259, 260, 5, 65, 0, 0, 260, 261, 5, 62, 0, 0, 261, 35,
		1, 0, 0, 0, 262, 266, 3, 38, 19, 0, 263, 265, 3, 40, 20, 0, 264, 263, 1,
		0, 0, 0, 265, 268, 1, 0, 0, 0, 266, 264, 1, 0, 0, 0, 266, 267, 1, 0, 0,
		0, 267, 270, 1, 0, 0, 0, 268, 266, 1, 0, 0, 0, 269, 271, 3, 42, 21, 0,
		270, 269, 1, 0, 0, 0, 270, 271, 1, 0, 0, 0, 271, 272, 1, 0, 0, 0, 272,
		273, 5, 19, 0, 0, 273, 274, 5, 73, 0, 0, 274, 275, 5, 85, 0, 0, 275, 37,
		1, 0, 0, 0, 276, 277, 5, 19, 0, 0, 277, 278, 5, 69, 0, 0, 278, 279, 3,
		26, 13, 0, 279, 283, 5, 85, 0, 0, 280, 282, 3, 16, 8, 0, 281, 280, 1, 0,
		0, 0, 282, 285, 1, 0, 0, 0, 283, 281, 1, 0, 0, 0, 283, 284, 1, 0, 0, 0,
		284, 39, 1, 0, 0, 0, 285, 283, 1, 0, 0, 0, 286, 287, 5, 19, 0, 0, 287,
		288, 5, 70, 0, 0, 288, 289, 3, 26, 13, 0, 289, 293, 5, 85, 0, 0, 290, 292,
		3, 16, 8, 0, 291, 290, 1, 0, 0, 0, 292, 295, 1, 0, 0, 0, 293, 291, 1, 0,
		0, 0, 293, 294, 1, 0, 0, 0, 294, 41, 1, 0, 0, 0, 295, 293, 1, 0, 0, 0,
		296, 297, 5, 19, 0, 0, 297, 298, 5, 71, 0, 0, 298, 302, 5, 85, 0, 0, 299,
		301, 3, 16, 8, 0, 300, 299, 1, 0, 0, 0, 301, 304, 1, 0, 0, 0, 302, 300,
		1, 0, 0, 0, 302, 303, 1, 0, 0, 0, 303, 43, 1, 0, 0, 0, 304, 302, 1, 0,
		0, 0, 305, 306, 5, 19, 0, 0, 306, 307, 5, 72, 0, 0, 307, 308, 3, 30, 15,
		0, 308, 309, 7, 5, 0, 0, 309, 310, 3, 26, 13, 0, 310, 311, 5, 85, 0, 0,
		311, 45, 1, 0, 0, 0, 312, 313, 5, 19, 0, 0, 313, 314, 5, 74, 0, 0, 314,
		315, 3, 32, 16, 0, 315, 316, 5, 85, 0, 0, 316, 47, 1, 0, 0, 0, 317, 318,
		5, 19, 0, 0, 318, 319, 3, 50, 25, 0, 319, 323, 5, 87, 0, 0, 320, 322, 3,
		22, 11, 0, 321, 320, 1, 0, 0, 0, 322, 325, 1, 0, 0, 0, 323, 321, 1, 0,
		0, 0, 323, 324, 1, 0, 0, 0, 324, 49, 1, 0, 0, 0, 325, 323, 1, 0, 0, 0,
		326, 332, 5, 89, 0, 0, 327, 328, 5, 88, 0, 0, 328, 329, 3, 26, 13, 0, 329,
		330, 5, 63, 0, 0, 330, 332, 1, 0, 0, 0, 331, 326, 1, 0, 0, 0, 331, 327,
		1, 0, 0, 0, 332, 335, 1, 0, 0, 0, 333, 331, 1, 0, 0, 0, 333, 334, 1, 0,
		0, 0, 334, 51, 1, 0, 0, 0, 335, 333, 1, 0, 0, 0, 336, 338, 3, 54, 27, 0,
		337, 336, 1, 0, 0, 0, 338, 341, 1, 0, 0, 0, 339, 337, 1, 0, 0, 0, 339,
		340, 1, 0, 0, 0, 340, 342, 1, 0, 0, 0, 341, 339, 1, 0, 0, 0, 342, 344,
		3, 54, 27, 0, 343, 345, 5, 3, 0, 0, 344, 343, 1, 0, 0, 0, 344, 345, 1,
		0, 0, 0, 345, 53, 1, 0, 0, 0, 346, 347, 5, 17, 0, 0, 347, 356, 3, 18, 9,
		0, 348, 352, 5, 1, 0, 0, 349, 351, 3, 16, 8, 0, 350, 349, 1, 0, 0, 0, 351,
		354, 1, 0, 0, 0, 352, 350, 1, 0, 0, 0, 352, 353, 1, 0, 0, 0, 353, 355,
		1, 0, 0, 0, 354, 352, 1, 0, 0, 0, 355, 357, 5, 2, 0, 0, 356, 348, 1, 0,
		0, 0, 356, 357, 1, 0, 0, 0, 357, 55, 1, 0, 0, 0, 358, 360, 3, 58, 29, 0,
		359, 358, 1, 0, 0, 0, 360, 363, 1, 0, 0, 0, 361, 359, 1, 0, 0, 0, 361,
		362, 1, 0, 0, 0, 362, 364, 1, 0, 0, 0, 363, 361, 1, 0, 0, 0, 364, 366,
		3, 58, 29, 0, 365, 367, 5, 3, 0, 0, 366, 365, 1, 0, 0, 0, 366, 367, 1,
		0, 0, 0, 367, 57, 1, 0, 0, 0, 368, 369, 5, 18, 0, 0, 369, 378, 3, 18, 9,
		0, 370, 374, 5, 1, 0, 0, 371, 373, 3, 16, 8, 0, 372, 371, 1, 0, 0, 0, 373,
		376, 1, 0, 0, 0, 374, 372, 1, 0, 0, 0, 374, 375, 1, 0, 0, 0, 375, 377,
		1, 0, 0, 0, 376, 374, 1, 0, 0, 0, 377, 379, 5, 2, 0, 0, 378, 370, 1, 0,
		0, 0, 378, 379, 1, 0, 0, 0, 379, 59, 1, 0, 0, 0, 380, 381, 5, 19, 0, 0,
		381, 382, 5, 75, 0, 0, 382, 383, 3, 30, 15, 0, 383, 384, 5, 36, 0, 0, 384,
		387, 3, 26, 13, 0, 385, 386, 5, 60, 0, 0, 386, 388, 5, 62, 0, 0, 387, 385,
		1, 0, 0, 0, 387, 388, 1, 0, 0, 0, 388, 389, 1, 0, 0, 0, 389, 390, 5, 85,
		0, 0, 390, 61, 1, 0, 0, 0, 391, 392, 5, 19, 0, 0, 392, 393, 5, 79, 0, 0,
		393, 394, 5, 9, 0, 0, 394, 396, 5, 85, 0, 0, 395, 397, 3, 64, 32, 0, 396,
		395, 1, 0, 0, 0, 397, 398, 1, 0, 0, 0, 398, 396, 1, 0, 0, 0, 398, 399,
		1, 0, 0, 0, 399, 400, 1, 0, 0, 0, 400, 401, 5, 19, 0, 0, 401, 402, 5, 81,
		0, 0, 402, 403, 5, 85, 0, 0, 403, 63, 1, 0, 0, 0, 404, 406, 5, 1, 0, 0,
		405, 404, 1, 0, 0, 0, 405, 406, 1, 0, 0, 0, 406, 407, 1, 0, 0, 0, 407,
		408, 5, 19, 0, 0, 408, 409, 5, 80, 0, 0, 409, 412, 5, 62, 0, 0, 410, 411,
		5, 36, 0, 0, 411, 413, 3, 28, 14, 0, 412, 410, 1, 0, 0, 0, 412, 413, 1,
		0, 0, 0, 413, 414, 1, 0, 0, 0, 414, 416, 5, 85, 0, 0, 415, 417, 5, 2, 0,
		0, 416, 415, 1, 0, 0, 0, 416, 417, 1, 0, 0, 0, 417, 65, 1, 0, 0, 0, 418,
		419, 5, 19, 0, 0, 419, 420, 5, 76, 0, 0, 420, 421, 5, 9, 0, 0, 421, 441,
		5, 85, 0, 0, 422, 423, 5, 19, 0, 0, 423, 424, 5, 76, 0, 0, 424, 425, 5,
		20, 0, 0, 425, 426, 3, 26, 13, 0, 426, 427, 5, 63, 0, 0, 427, 428, 5, 85,
		0, 0, 428, 441, 1, 0, 0, 0, 429, 430, 5, 19, 0, 0, 430, 431, 5, 77, 0,
		0, 431, 432, 5, 9, 0, 0, 432, 441, 5, 85, 0, 0, 433, 434, 5, 19, 0, 0,
		434, 435, 5, 77, 0, 0, 435, 436, 5, 20, 0, 0, 436, 437, 3, 26, 13, 0, 437,
		438, 5, 63, 0, 0, 438, 439, 5, 85, 0, 0, 439, 441, 1, 0, 0, 0, 440, 418,
		1, 0, 0, 0, 440, 422, 1, 0, 0, 0, 440, 429, 1, 0, 0, 0, 440, 433, 1, 0,
		0, 0, 441, 67, 1, 0, 0, 0, 442, 443, 5, 19, 0, 0, 443, 444, 5, 78, 0, 0,
		444, 445, 5, 85, 0, 0, 445, 69, 1, 0, 0, 0, 446, 448, 3, 72, 36, 0, 447,
		449, 3, 74, 37, 0, 448, 447, 1, 0, 0, 0, 448, 449, 1, 0, 0, 0, 449, 450,
		1, 0, 0, 0, 450, 451, 5, 19, 0, 0, 451, 452, 5, 83, 0, 0, 452, 453, 5,
		85, 0, 0, 453, 71, 1, 0, 0, 0, 454, 455, 5, 19, 0, 0, 455, 458, 5, 82,
		0, 0, 456, 457, 5, 69, 0, 0, 457, 459, 3, 26, 13, 0, 458, 456, 1, 0, 0,
		0, 458, 459, 1, 0, 0, 0, 459, 460, 1, 0, 0, 0, 460, 464, 5, 85, 0, 0, 461,
		463, 3, 16, 8, 0, 462, 461, 1, 0, 0, 0, 463, 466, 1, 0, 0, 0, 464, 462,
		1, 0, 0, 0, 464, 465, 1, 0, 0, 0, 465, 73, 1, 0, 0, 0, 466, 464, 1, 0,
		0, 0, 467, 468, 5, 19, 0, 0, 468, 469, 5, 71, 0, 0, 469, 473, 5, 85, 0,
		0, 470, 472, 3, 16, 8, 0, 471, 470, 1, 0, 0, 0, 472, 475, 1, 0, 0, 0, 473,
		471, 1, 0, 0, 0, 473, 474, 1, 0, 0, 0, 474, 75, 1, 0, 0, 0, 475, 473, 1,
		0, 0, 0, 476, 480, 5, 62, 0, 0, 477, 479, 3, 78, 39, 0, 478, 477, 1, 0,
		0, 0, 479, 482, 1, 0, 0, 0, 480, 478, 1, 0, 0, 0, 480, 481, 1, 0, 0, 0,
		481, 77, 1, 0, 0, 0, 482, 480, 1, 0, 0, 0, 483, 486, 3, 26, 13, 0, 484,
		486, 5, 62, 0, 0, 485, 483, 1, 0, 0, 0, 485, 484, 1, 0, 0, 0, 486, 79,
		1, 0, 0, 0, 52, 83, 89, 97, 99, 118, 125, 127, 132, 151, 155, 159, 164,
		172, 178, 180, 194, 197, 209, 226, 228, 238, 245, 251, 257, 266, 270, 283,
		293, 302, 323, 331, 333, 339, 344, 352, 356, 361, 366, 374, 378, 387, 398,
		405, 412, 416, 440, 448, 458, 464, 473, 480, 485,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// YarnSpinnerParserInit initializes any static state used to implement YarnSpinnerParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewYarnSpinnerParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func YarnSpinnerParserInit() {
	staticData := &YarnSpinnerParserParserStaticData
	staticData.once.Do(yarnspinnerparserParserInit)
}

// NewYarnSpinnerParser produces a new parser instance for the optional input antlr.TokenStream.
func NewYarnSpinnerParser(input antlr.TokenStream) *YarnSpinnerParser {
	YarnSpinnerParserInit()
	this := new(YarnSpinnerParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &YarnSpinnerParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "YarnSpinnerParser.g4"

	return this
}

// YarnSpinnerParser tokens.
const (
	YarnSpinnerParserEOF                                  = antlr.TokenEOF
	YarnSpinnerParserINDENT                               = 1
	YarnSpinnerParserDEDENT                               = 2
	YarnSpinnerParserBLANK_LINE_FOLLOWING_OPTION          = 3
	YarnSpinnerParserWS                                   = 4
	YarnSpinnerParserCOMMENT                              = 5
	YarnSpinnerParserNEWLINE                              = 6
	YarnSpinnerParserHEADER_WHEN                          = 7
	YarnSpinnerParserHEADER_TITLE                         = 8
	YarnSpinnerParserID                                   = 9
	YarnSpinnerParserBODY_START                           = 10
	YarnSpinnerParserHEADER_DELIMITER                     = 11
	YarnSpinnerParserHASHTAG                              = 12
	YarnSpinnerParserHEADER_WHEN_UNKNOWN                  = 13
	YarnSpinnerParserREST_OF_LINE                         = 14
	YarnSpinnerParserBODY_WS                              = 15
	YarnSpinnerParserBODY_END                             = 16
	YarnSpinnerParserSHORTCUT_ARROW                       = 17
	YarnSpinnerParserLINE_GROUP_ARROW                     = 18
	YarnSpinnerParserCOMMAND_START                        = 19
	YarnSpinnerParserEXPRESSION_START                     = 20
	YarnSpinnerParserESCAPED_ANY                          = 21
	YarnSpinnerParserTEXT_ESCAPE                          = 22
	YarnSpinnerParserTEXT_COMMENT                         = 23
	YarnSpinnerParserTEXT                                 = 24
	YarnSpinnerParserUNESCAPABLE_CHARACTER                = 25
	YarnSpinnerParserTEXT_COMMANDHASHTAG_WS               = 26
	YarnSpinnerParserTEXT_COMMANDHASHTAG_COMMENT          = 27
	YarnSpinnerParserTEXT_COMMANDHASHTAG_ERROR            = 28
	YarnSpinnerParserHASHTAG_WS                           = 29
	YarnSpinnerParserHASHTAG_TEXT                         = 30
	YarnSpinnerParserEXPR_WS                              = 31
	YarnSpinnerParserEXPRESSION_WHEN_ALWAYS               = 32
	YarnSpinnerParserKEYWORD_TRUE                         = 33
	YarnSpinnerParserKEYWORD_FALSE                        = 34
	YarnSpinnerParserKEYWORD_NULL                         = 35
	YarnSpinnerParserOPERATOR_ASSIGNMENT                  = 36
	YarnSpinnerParserOPERATOR_LOGICAL_LESS_THAN_EQUALS    = 37
	YarnSpinnerParserOPERATOR_LOGICAL_GREATER_THAN_EQUALS = 38
	YarnSpinnerParserOPERATOR_LOGICAL_EQUALS              = 39
	YarnSpinnerParserOPERATOR_LOGICAL_LESS                = 40
	YarnSpinnerParserOPERATOR_LOGICAL_GREATER             = 41
	YarnSpinnerParserOPERATOR_LOGICAL_NOT_EQUALS          = 42
	YarnSpinnerParserOPERATOR_LOGICAL_AND                 = 43
	YarnSpinnerParserOPERATOR_LOGICAL_OR                  = 44
	YarnSpinnerParserOPERATOR_LOGICAL_XOR                 = 45
	YarnSpinnerParserOPERATOR_LOGICAL_NOT                 = 46
	YarnSpinnerParserOPERATOR_MATHS_ADDITION_EQUALS       = 47
	YarnSpinnerParserOPERATOR_MATHS_SUBTRACTION_EQUALS    = 48
	YarnSpinnerParserOPERATOR_MATHS_MULTIPLICATION_EQUALS = 49
	YarnSpinnerParserOPERATOR_MATHS_MODULUS_EQUALS        = 50
	YarnSpinnerParserOPERATOR_MATHS_DIVISION_EQUALS       = 51
	YarnSpinnerParserOPERATOR_MATHS_ADDITION              = 52
	YarnSpinnerParserOPERATOR_MATHS_SUBTRACTION           = 53
	YarnSpinnerParserOPERATOR_MATHS_MULTIPLICATION        = 54
	YarnSpinnerParserOPERATOR_MATHS_DIVISION              = 55
	YarnSpinnerParserOPERATOR_MATHS_MODULUS               = 56
	YarnSpinnerParserLPAREN                               = 57
	YarnSpinnerParserRPAREN                               = 58
	YarnSpinnerParserCOMMA                                = 59
	YarnSpinnerParserEXPRESSION_AS                        = 60
	YarnSpinnerParserSTRING                               = 61
	YarnSpinnerParserFUNC_ID                              = 62
	YarnSpinnerParserEXPRESSION_END                       = 63
	YarnSpinnerParserVAR_ID                               = 64
	YarnSpinnerParserDOT                                  = 65
	YarnSpinnerParserNUMBER                               = 66
	YarnSpinnerParserCOMMAND_NEWLINE                      = 67
	YarnSpinnerParserCOMMAND_WS                           = 68
	YarnSpinnerParserCOMMAND_IF                           = 69
	YarnSpinnerParserCOMMAND_ELSEIF                       = 70
	YarnSpinnerParserCOMMAND_ELSE                         = 71
	YarnSpinnerParserCOMMAND_SET                          = 72
	YarnSpinnerParserCOMMAND_ENDIF                        = 73
	YarnSpinnerParserCOMMAND_CALL                         = 74
	YarnSpinnerParserCOMMAND_DECLARE                      = 75
	YarnSpinnerParserCOMMAND_JUMP                         = 76
	YarnSpinnerParserCOMMAND_DETOUR                       = 77
	YarnSpinnerParserCOMMAND_RETURN                       = 78
	YarnSpinnerParserCOMMAND_ENUM                         = 79
	YarnSpinnerParserCOMMAND_CASE                         = 80
	YarnSpinnerParserCOMMAND_ENDENUM                      = 81
	YarnSpinnerParserCOMMAND_ONCE                         = 82
	YarnSpinnerParserCOMMAND_ENDONCE                      = 83
	YarnSpinnerParserCOMMAND_LOCAL                        = 84
	YarnSpinnerParserCOMMAND_END                          = 85
	YarnSpinnerParserCOMMAND_TEXT_NEWLINE                 = 86
	YarnSpinnerParserCOMMAND_TEXT_END                     = 87
	YarnSpinnerParserCOMMAND_EXPRESSION_START             = 88
	YarnSpinnerParserCOMMAND_TEXT                         = 89
	YarnSpinnerParserCOMMAND_ID_NEWLINE                   = 90
	YarnSpinnerParserTYPE_STRING                          = 91
	YarnSpinnerParserTYPE_NUMBER                          = 92
	YarnSpinnerParserTYPE_BOOL                            = 93
)

// YarnSpinnerParser rules.
const (
	YarnSpinnerParserRULE_dialogue                  = 0
	YarnSpinnerParserRULE_file_hashtag              = 1
	YarnSpinnerParserRULE_node                      = 2
	YarnSpinnerParserRULE_title_header              = 3
	YarnSpinnerParserRULE_when_header               = 4
	YarnSpinnerParserRULE_header                    = 5
	YarnSpinnerParserRULE_header_when_expression    = 6
	YarnSpinnerParserRULE_body                      = 7
	YarnSpinnerParserRULE_statement                 = 8
	YarnSpinnerParserRULE_line_statement            = 9
	YarnSpinnerParserRULE_line_formatted_text       = 10
	YarnSpinnerParserRULE_hashtag                   = 11
	YarnSpinnerParserRULE_line_condition            = 12
	YarnSpinnerParserRULE_expression                = 13
	YarnSpinnerParserRULE_value                     = 14
	YarnSpinnerParserRULE_variable                  = 15
	YarnSpinnerParserRULE_function_call             = 16
	YarnSpinnerParserRULE_typeMemberReference       = 17
	YarnSpinnerParserRULE_if_statement              = 18
	YarnSpinnerParserRULE_if_clause                 = 19
	YarnSpinnerParserRULE_else_if_clause            = 20
	YarnSpinnerParserRULE_else_clause               = 21
	YarnSpinnerParserRULE_set_statement             = 22
	YarnSpinnerParserRULE_call_statement            = 23
	YarnSpinnerParserRULE_command_statement         = 24
	YarnSpinnerParserRULE_command_formatted_text    = 25
	YarnSpinnerParserRULE_shortcut_option_statement = 26
	YarnSpinnerParserRULE_shortcut_option           = 27
	YarnSpinnerParserRULE_line_group_statement      = 28
	YarnSpinnerParserRULE_line_group_item           = 29
	YarnSpinnerParserRULE_declare_statement         = 30
	YarnSpinnerParserRULE_enum_statement            = 31
	YarnSpinnerParserRULE_enum_case_statement       = 32
	YarnSpinnerParserRULE_jump_statement            = 33
	YarnSpinnerParserRULE_return_statement          = 34
	YarnSpinnerParserRULE_once_statement            = 35
	YarnSpinnerParserRULE_once_primary_clause       = 36
	YarnSpinnerParserRULE_once_alternate_clause     = 37
	YarnSpinnerParserRULE_structured_command        = 38
	YarnSpinnerParserRULE_structured_command_value  = 39
)

// IDialogueContext is an interface to support dynamic dispatch.
type IDialogueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNode() []INodeContext
	Node(i int) INodeContext
	AllFile_hashtag() []IFile_hashtagContext
	File_hashtag(i int) IFile_hashtagContext

	// IsDialogueContext differentiates from other interfaces.
	IsDialogueContext()
}

type DialogueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDialogueContext() *DialogueContext {
	var p = new(DialogueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_dialogue
	return p
}

func InitEmptyDialogueContext(p *DialogueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_dialogue
}

func (*DialogueContext) IsDialogueContext() {}

func NewDialogueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DialogueContext {
	var p = new(DialogueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YarnSpinnerParserRULE_dialogue

	return p
}

func (s *DialogueContext) GetParser() antlr.Parser { return s.parser }

func (s *DialogueContext) AllNode() []INodeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INodeContext); ok {
			len++
		}
	}

	tst := make([]INodeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INodeContext); ok {
			tst[i] = t.(INodeContext)
			i++
		}
	}

	return tst
}

func (s *DialogueContext) Node(i int) INodeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INodeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INodeContext)
}

func (s *DialogueContext) AllFile_hashtag() []IFile_hashtagContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFile_hashtagContext); ok {
			len++
		}
	}

	tst := make([]IFile_hashtagContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFile_hashtagContext); ok {
			tst[i] = t.(IFile_hashtagContext)
			i++
		}
	}

	return tst
}

func (s *DialogueContext) File_hashtag(i int) IFile_hashtagContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFile_hashtagContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFile_hashtagContext)
}

func (s *DialogueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DialogueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DialogueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterDialogue(s)
	}
}

func (s *DialogueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitDialogue(s)
	}
}

func (p *YarnSpinnerParser) Dialogue() (localctx IDialogueContext) {
	localctx = NewDialogueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, YarnSpinnerParserRULE_dialogue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == YarnSpinnerParserHASHTAG {
		{
			p.SetState(80)
			p.File_hashtag()
		}

		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&896) != 0) {
		{
			p.SetState(86)
			p.Node()
		}

		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFile_hashtagContext is an interface to support dynamic dispatch.
type IFile_hashtagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetInnerText returns the text token.
	GetInnerText() antlr.Token

	// SetText sets the text token.
	SetText(antlr.Token)

	// Getter signatures
	HASHTAG() antlr.TerminalNode
	HASHTAG_TEXT() antlr.TerminalNode

	// IsFile_hashtagContext differentiates from other interfaces.
	IsFile_hashtagContext()
}

type File_hashtagContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	text   antlr.Token
}

func NewEmptyFile_hashtagContext() *File_hashtagContext {
	var p = new(File_hashtagContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_file_hashtag
	return p
}

func InitEmptyFile_hashtagContext(p *File_hashtagContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_file_hashtag
}

func (*File_hashtagContext) IsFile_hashtagContext() {}

func NewFile_hashtagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *File_hashtagContext {
	var p = new(File_hashtagContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YarnSpinnerParserRULE_file_hashtag

	return p
}

func (s *File_hashtagContext) GetParser() antlr.Parser { return s.parser }

func (s *File_hashtagContext) GetInnerText() antlr.Token { return s.text }

func (s *File_hashtagContext) SetText(v antlr.Token) { s.text = v }

func (s *File_hashtagContext) HASHTAG() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserHASHTAG, 0)
}

func (s *File_hashtagContext) HASHTAG_TEXT() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserHASHTAG_TEXT, 0)
}

func (s *File_hashtagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *File_hashtagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *File_hashtagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterFile_hashtag(s)
	}
}

func (s *File_hashtagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitFile_hashtag(s)
	}
}

func (p *YarnSpinnerParser) File_hashtag() (localctx IFile_hashtagContext) {
	localctx = NewFile_hashtagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, YarnSpinnerParserRULE_file_hashtag)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(91)
		p.Match(YarnSpinnerParserHASHTAG)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(92)

		var _m = p.Match(YarnSpinnerParserHASHTAG_TEXT)

		localctx.(*File_hashtagContext).text = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INodeContext is an interface to support dynamic dispatch.
type INodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BODY_START() antlr.TerminalNode
	Body() IBodyContext
	BODY_END() antlr.TerminalNode
	AllHeader() []IHeaderContext
	Header(i int) IHeaderContext
	AllWhen_header() []IWhen_headerContext
	When_header(i int) IWhen_headerContext
	AllTitle_header() []ITitle_headerContext
	Title_header(i int) ITitle_headerContext

	// IsNodeContext differentiates from other interfaces.
	IsNodeContext()
}

type NodeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNodeContext() *NodeContext {
	var p = new(NodeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_node
	return p
}

func InitEmptyNodeContext(p *NodeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_node
}

func (*NodeContext) IsNodeContext() {}

func NewNodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NodeContext {
	var p = new(NodeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YarnSpinnerParserRULE_node

	return p
}

func (s *NodeContext) GetParser() antlr.Parser { return s.parser }

func (s *NodeContext) BODY_START() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserBODY_START, 0)
}

func (s *NodeContext) Body() IBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBodyContext)
}

func (s *NodeContext) BODY_END() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserBODY_END, 0)
}

func (s *NodeContext) AllHeader() []IHeaderContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IHeaderContext); ok {
			len++
		}
	}

	tst := make([]IHeaderContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IHeaderContext); ok {
			tst[i] = t.(IHeaderContext)
			i++
		}
	}

	return tst
}

func (s *NodeContext) Header(i int) IHeaderContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHeaderContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHeaderContext)
}

func (s *NodeContext) AllWhen_header() []IWhen_headerContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IWhen_headerContext); ok {
			len++
		}
	}

	tst := make([]IWhen_headerContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IWhen_headerContext); ok {
			tst[i] = t.(IWhen_headerContext)
			i++
		}
	}

	return tst
}

func (s *NodeContext) When_header(i int) IWhen_headerContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhen_headerContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhen_headerContext)
}

func (s *NodeContext) AllTitle_header() []ITitle_headerContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITitle_headerContext); ok {
			len++
		}
	}

	tst := make([]ITitle_headerContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITitle_headerContext); ok {
			tst[i] = t.(ITitle_headerContext)
			i++
		}
	}

	return tst
}

func (s *NodeContext) Title_header(i int) ITitle_headerContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITitle_headerContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITitle_headerContext)
}

func (s *NodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterNode(s)
	}
}

func (s *NodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitNode(s)
	}
}

func (p *YarnSpinnerParser) Node() (localctx INodeContext) {
	localctx = NewNodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, YarnSpinnerParserRULE_node)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&896) != 0) {
		p.SetState(97)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case YarnSpinnerParserID:
			{
				p.SetState(94)
				p.Header()
			}

		case YarnSpinnerParserHEADER_WHEN:
			{
				p.SetState(95)
				p.When_header()
			}

		case YarnSpinnerParserHEADER_TITLE:
			{
				p.SetState(96)
				p.Title_header()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(99)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(101)
		p.Match(YarnSpinnerParserBODY_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(102)
		p.Body()
	}
	{
		p.SetState(103)
		p.Match(YarnSpinnerParserBODY_END)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITitle_headerContext is an interface to support dynamic dispatch.
type ITitle_headerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTitle returns the title token.
	GetTitle() antlr.Token

	// SetTitle sets the title token.
	SetTitle(antlr.Token)

	// Getter signatures
	HEADER_TITLE() antlr.TerminalNode
	HEADER_DELIMITER() antlr.TerminalNode
	NEWLINE() antlr.TerminalNode
	ID() antlr.TerminalNode

	// IsTitle_headerContext differentiates from other interfaces.
	IsTitle_headerContext()
}

type Title_headerContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	title  antlr.Token
}

func NewEmptyTitle_headerContext() *Title_headerContext {
	var p = new(Title_headerContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_title_header
	return p
}

func InitEmptyTitle_headerContext(p *Title_headerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_title_header
}

func (*Title_headerContext) IsTitle_headerContext() {}

func NewTitle_headerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Title_headerContext {
	var p = new(Title_headerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YarnSpinnerParserRULE_title_header

	return p
}

func (s *Title_headerContext) GetParser() antlr.Parser { return s.parser }

func (s *Title_headerContext) GetTitle() antlr.Token { return s.title }

func (s *Title_headerContext) SetTitle(v antlr.Token) { s.title = v }

func (s *Title_headerContext) HEADER_TITLE() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserHEADER_TITLE, 0)
}

func (s *Title_headerContext) HEADER_DELIMITER() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserHEADER_DELIMITER, 0)
}

func (s *Title_headerContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserNEWLINE, 0)
}

func (s *Title_headerContext) ID() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserID, 0)
}

func (s *Title_headerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Title_headerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Title_headerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterTitle_header(s)
	}
}

func (s *Title_headerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitTitle_header(s)
	}
}

func (p *YarnSpinnerParser) Title_header() (localctx ITitle_headerContext) {
	localctx = NewTitle_headerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, YarnSpinnerParserRULE_title_header)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(105)
		p.Match(YarnSpinnerParserHEADER_TITLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(106)
		p.Match(YarnSpinnerParserHEADER_DELIMITER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(107)

		var _m = p.Match(YarnSpinnerParserID)

		localctx.(*Title_headerContext).title = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(108)
		p.Match(YarnSpinnerParserNEWLINE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWhen_headerContext is an interface to support dynamic dispatch.
type IWhen_headerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetHeader_expression returns the header_expression rule contexts.
	GetHeader_expression() IHeader_when_expressionContext

	// SetHeader_expression sets the header_expression rule contexts.
	SetHeader_expression(IHeader_when_expressionContext)

	// Getter signatures
	HEADER_WHEN() antlr.TerminalNode
	HEADER_DELIMITER() antlr.TerminalNode
	NEWLINE() antlr.TerminalNode
	Header_when_expression() IHeader_when_expressionContext

	// IsWhen_headerContext differentiates from other interfaces.
	IsWhen_headerContext()
}

type When_headerContext struct {
	antlr.BaseParserRuleContext
	parser            antlr.Parser
	header_expression IHeader_when_expressionContext
}

func NewEmptyWhen_headerContext() *When_headerContext {
	var p = new(When_headerContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_when_header
	return p
}

func InitEmptyWhen_headerContext(p *When_headerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_when_header
}

func (*When_headerContext) IsWhen_headerContext() {}

func NewWhen_headerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *When_headerContext {
	var p = new(When_headerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YarnSpinnerParserRULE_when_header

	return p
}

func (s *When_headerContext) GetParser() antlr.Parser { return s.parser }

func (s *When_headerContext) GetHeader_expression() IHeader_when_expressionContext {
	return s.header_expression
}

func (s *When_headerContext) SetHeader_expression(v IHeader_when_expressionContext) {
	s.header_expression = v
}

func (s *When_headerContext) HEADER_WHEN() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserHEADER_WHEN, 0)
}

func (s *When_headerContext) HEADER_DELIMITER() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserHEADER_DELIMITER, 0)
}

func (s *When_headerContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserNEWLINE, 0)
}

func (s *When_headerContext) Header_when_expression() IHeader_when_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHeader_when_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHeader_when_expressionContext)
}

func (s *When_headerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *When_headerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *When_headerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterWhen_header(s)
	}
}

func (s *When_headerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitWhen_header(s)
	}
}

func (p *YarnSpinnerParser) When_header() (localctx IWhen_headerContext) {
	localctx = NewWhen_headerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, YarnSpinnerParserRULE_when_header)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(110)
		p.Match(YarnSpinnerParserHEADER_WHEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(111)
		p.Match(YarnSpinnerParserHEADER_DELIMITER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(112)

		var _x = p.Header_when_expression()

		localctx.(*When_headerContext).header_expression = _x
	}
	{
		p.SetState(113)
		p.Match(YarnSpinnerParserNEWLINE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IHeaderContext is an interface to support dynamic dispatch.
type IHeaderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetHeader_key returns the header_key token.
	GetHeader_key() antlr.Token

	// GetHeader_value returns the header_value token.
	GetHeader_value() antlr.Token

	// SetHeader_key sets the header_key token.
	SetHeader_key(antlr.Token)

	// SetHeader_value sets the header_value token.
	SetHeader_value(antlr.Token)

	// Getter signatures
	HEADER_DELIMITER() antlr.TerminalNode
	ID() antlr.TerminalNode
	REST_OF_LINE() antlr.TerminalNode

	// IsHeaderContext differentiates from other interfaces.
	IsHeaderContext()
}

type HeaderContext struct {
	antlr.BaseParserRuleContext
	parser       antlr.Parser
	header_key   antlr.Token
	header_value antlr.Token
}

func NewEmptyHeaderContext() *HeaderContext {
	var p = new(HeaderContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_header
	return p
}

func InitEmptyHeaderContext(p *HeaderContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_header
}

func (*HeaderContext) IsHeaderContext() {}

func NewHeaderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HeaderContext {
	var p = new(HeaderContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YarnSpinnerParserRULE_header

	return p
}

func (s *HeaderContext) GetParser() antlr.Parser { return s.parser }

func (s *HeaderContext) GetHeader_key() antlr.Token { return s.header_key }

func (s *HeaderContext) GetHeader_value() antlr.Token { return s.header_value }

func (s *HeaderContext) SetHeader_key(v antlr.Token) { s.header_key = v }

func (s *HeaderContext) SetHeader_value(v antlr.Token) { s.header_value = v }

func (s *HeaderContext) HEADER_DELIMITER() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserHEADER_DELIMITER, 0)
}

func (s *HeaderContext) ID() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserID, 0)
}

func (s *HeaderContext) REST_OF_LINE() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserREST_OF_LINE, 0)
}

func (s *HeaderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HeaderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HeaderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterHeader(s)
	}
}

func (s *HeaderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitHeader(s)
	}
}

func (p *YarnSpinnerParser) Header() (localctx IHeaderContext) {
	localctx = NewHeaderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, YarnSpinnerParserRULE_header)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(115)

		var _m = p.Match(YarnSpinnerParserID)

		localctx.(*HeaderContext).header_key = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(116)
		p.Match(YarnSpinnerParserHEADER_DELIMITER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == YarnSpinnerParserREST_OF_LINE {
		{
			p.SetState(117)

			var _m = p.Match(YarnSpinnerParserREST_OF_LINE)

			localctx.(*HeaderContext).header_value = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IHeader_when_expressionContext is an interface to support dynamic dispatch.
type IHeader_when_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetAlways returns the always token.
	GetAlways() antlr.Token

	// GetOnce returns the once token.
	GetOnce() antlr.Token

	// SetAlways sets the always token.
	SetAlways(antlr.Token)

	// SetOnce sets the once token.
	SetOnce(antlr.Token)

	// Getter signatures
	Expression() IExpressionContext
	EXPRESSION_WHEN_ALWAYS() antlr.TerminalNode
	COMMAND_ONCE() antlr.TerminalNode
	COMMAND_IF() antlr.TerminalNode

	// IsHeader_when_expressionContext differentiates from other interfaces.
	IsHeader_when_expressionContext()
}

type Header_when_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	always antlr.Token
	once   antlr.Token
}

func NewEmptyHeader_when_expressionContext() *Header_when_expressionContext {
	var p = new(Header_when_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_header_when_expression
	return p
}

func InitEmptyHeader_when_expressionContext(p *Header_when_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_header_when_expression
}

func (*Header_when_expressionContext) IsHeader_when_expressionContext() {}

func NewHeader_when_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Header_when_expressionContext {
	var p = new(Header_when_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YarnSpinnerParserRULE_header_when_expression

	return p
}

func (s *Header_when_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Header_when_expressionContext) GetAlways() antlr.Token { return s.always }

func (s *Header_when_expressionContext) GetOnce() antlr.Token { return s.once }

func (s *Header_when_expressionContext) SetAlways(v antlr.Token) { s.always = v }

func (s *Header_when_expressionContext) SetOnce(v antlr.Token) { s.once = v }

func (s *Header_when_expressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Header_when_expressionContext) EXPRESSION_WHEN_ALWAYS() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserEXPRESSION_WHEN_ALWAYS, 0)
}

func (s *Header_when_expressionContext) COMMAND_ONCE() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_ONCE, 0)
}

func (s *Header_when_expressionContext) COMMAND_IF() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_IF, 0)
}

func (s *Header_when_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Header_when_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Header_when_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterHeader_when_expression(s)
	}
}

func (s *Header_when_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitHeader_when_expression(s)
	}
}

func (p *YarnSpinnerParser) Header_when_expression() (localctx IHeader_when_expressionContext) {
	localctx = NewHeader_when_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, YarnSpinnerParserRULE_header_when_expression)
	var _la int

	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case YarnSpinnerParserKEYWORD_TRUE, YarnSpinnerParserKEYWORD_FALSE, YarnSpinnerParserOPERATOR_LOGICAL_NOT, YarnSpinnerParserOPERATOR_MATHS_SUBTRACTION, YarnSpinnerParserLPAREN, YarnSpinnerParserSTRING, YarnSpinnerParserFUNC_ID, YarnSpinnerParserVAR_ID, YarnSpinnerParserDOT, YarnSpinnerParserNUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(120)
			p.expression(0)
		}

	case YarnSpinnerParserEXPRESSION_WHEN_ALWAYS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(121)

			var _m = p.Match(YarnSpinnerParserEXPRESSION_WHEN_ALWAYS)

			localctx.(*Header_when_expressionContext).always = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case YarnSpinnerParserCOMMAND_ONCE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(122)

			var _m = p.Match(YarnSpinnerParserCOMMAND_ONCE)

			localctx.(*Header_when_expressionContext).once = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(125)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == YarnSpinnerParserCOMMAND_IF {
			{
				p.SetState(123)
				p.Match(YarnSpinnerParserCOMMAND_IF)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(124)
				p.expression(0)
			}

		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBodyContext is an interface to support dynamic dispatch.
type IBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsBodyContext differentiates from other interfaces.
	IsBodyContext()
}

type BodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBodyContext() *BodyContext {
	var p = new(BodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_body
	return p
}

func InitEmptyBodyContext(p *BodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_body
}

func (*BodyContext) IsBodyContext() {}

func NewBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BodyContext {
	var p = new(BodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YarnSpinnerParserRULE_body

	return p
}

func (s *BodyContext) GetParser() antlr.Parser { return s.parser }

func (s *BodyContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *BodyContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterBody(s)
	}
}

func (s *BodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitBody(s)
	}
}

func (p *YarnSpinnerParser) Body() (localctx IBodyContext) {
	localctx = NewBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, YarnSpinnerParserRULE_body)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&18743298) != 0 {
		{
			p.SetState(129)
			p.Statement()
		}

		p.SetState(134)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Line_statement() ILine_statementContext
	If_statement() IIf_statementContext
	Set_statement() ISet_statementContext
	Shortcut_option_statement() IShortcut_option_statementContext
	Call_statement() ICall_statementContext
	Command_statement() ICommand_statementContext
	Declare_statement() IDeclare_statementContext
	Enum_statement() IEnum_statementContext
	Jump_statement() IJump_statementContext
	Return_statement() IReturn_statementContext
	Line_group_statement() ILine_group_statementContext
	Once_statement() IOnce_statementContext
	INDENT() antlr.TerminalNode
	DEDENT() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YarnSpinnerParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Line_statement() ILine_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILine_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILine_statementContext)
}

func (s *StatementContext) If_statement() IIf_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIf_statementContext)
}

func (s *StatementContext) Set_statement() ISet_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISet_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISet_statementContext)
}

func (s *StatementContext) Shortcut_option_statement() IShortcut_option_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IShortcut_option_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IShortcut_option_statementContext)
}

func (s *StatementContext) Call_statement() ICall_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICall_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICall_statementContext)
}

func (s *StatementContext) Command_statement() ICommand_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommand_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommand_statementContext)
}

func (s *StatementContext) Declare_statement() IDeclare_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclare_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclare_statementContext)
}

func (s *StatementContext) Enum_statement() IEnum_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnum_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnum_statementContext)
}

func (s *StatementContext) Jump_statement() IJump_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJump_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJump_statementContext)
}

func (s *StatementContext) Return_statement() IReturn_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturn_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturn_statementContext)
}

func (s *StatementContext) Line_group_statement() ILine_group_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILine_group_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILine_group_statementContext)
}

func (s *StatementContext) Once_statement() IOnce_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOnce_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOnce_statementContext)
}

func (s *StatementContext) INDENT() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserINDENT, 0)
}

func (s *StatementContext) DEDENT() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserDEDENT, 0)
}

func (s *StatementContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *StatementContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *YarnSpinnerParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, YarnSpinnerParserRULE_statement)
	var _la int

	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(135)
			p.Line_statement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(136)
			p.If_statement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(137)
			p.Set_statement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(138)
			p.Shortcut_option_statement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(139)
			p.Call_statement()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(140)
			p.Command_statement()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(141)
			p.Declare_statement()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(142)
			p.Enum_statement()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(143)
			p.Jump_statement()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(144)
			p.Return_statement()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(145)
			p.Line_group_statement()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(146)
			p.Once_statement()
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(147)
			p.Match(YarnSpinnerParserINDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&18743298) != 0 {
			{
				p.SetState(148)
				p.Statement()
			}

			p.SetState(153)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(154)
			p.Match(YarnSpinnerParserDEDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILine_statementContext is an interface to support dynamic dispatch.
type ILine_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Line_formatted_text() ILine_formatted_textContext
	NEWLINE() antlr.TerminalNode
	Line_condition() ILine_conditionContext
	AllHashtag() []IHashtagContext
	Hashtag(i int) IHashtagContext

	// IsLine_statementContext differentiates from other interfaces.
	IsLine_statementContext()
}

type Line_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLine_statementContext() *Line_statementContext {
	var p = new(Line_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_line_statement
	return p
}

func InitEmptyLine_statementContext(p *Line_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_line_statement
}

func (*Line_statementContext) IsLine_statementContext() {}

func NewLine_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Line_statementContext {
	var p = new(Line_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YarnSpinnerParserRULE_line_statement

	return p
}

func (s *Line_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Line_statementContext) Line_formatted_text() ILine_formatted_textContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILine_formatted_textContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILine_formatted_textContext)
}

func (s *Line_statementContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserNEWLINE, 0)
}

func (s *Line_statementContext) Line_condition() ILine_conditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILine_conditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILine_conditionContext)
}

func (s *Line_statementContext) AllHashtag() []IHashtagContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IHashtagContext); ok {
			len++
		}
	}

	tst := make([]IHashtagContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IHashtagContext); ok {
			tst[i] = t.(IHashtagContext)
			i++
		}
	}

	return tst
}

func (s *Line_statementContext) Hashtag(i int) IHashtagContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHashtagContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHashtagContext)
}

func (s *Line_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Line_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Line_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterLine_statement(s)
	}
}

func (s *Line_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitLine_statement(s)
	}
}

func (p *YarnSpinnerParser) Line_statement() (localctx ILine_statementContext) {
	localctx = NewLine_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, YarnSpinnerParserRULE_line_statement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(157)
		p.Line_formatted_text()
	}
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == YarnSpinnerParserCOMMAND_START {
		{
			p.SetState(158)
			p.Line_condition()
		}

	}
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == YarnSpinnerParserHASHTAG {
		{
			p.SetState(161)
			p.Hashtag()
		}

		p.SetState(166)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(167)
		p.Match(YarnSpinnerParserNEWLINE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILine_formatted_textContext is an interface to support dynamic dispatch.
type ILine_formatted_textContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllEXPRESSION_START() []antlr.TerminalNode
	EXPRESSION_START(i int) antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllEXPRESSION_END() []antlr.TerminalNode
	EXPRESSION_END(i int) antlr.TerminalNode
	AllTEXT() []antlr.TerminalNode
	TEXT(i int) antlr.TerminalNode

	// IsLine_formatted_textContext differentiates from other interfaces.
	IsLine_formatted_textContext()
}

type Line_formatted_textContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLine_formatted_textContext() *Line_formatted_textContext {
	var p = new(Line_formatted_textContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_line_formatted_text
	return p
}

func InitEmptyLine_formatted_textContext(p *Line_formatted_textContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_line_formatted_text
}

func (*Line_formatted_textContext) IsLine_formatted_textContext() {}

func NewLine_formatted_textContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Line_formatted_textContext {
	var p = new(Line_formatted_textContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YarnSpinnerParserRULE_line_formatted_text

	return p
}

func (s *Line_formatted_textContext) GetParser() antlr.Parser { return s.parser }

func (s *Line_formatted_textContext) AllEXPRESSION_START() []antlr.TerminalNode {
	return s.GetTokens(YarnSpinnerParserEXPRESSION_START)
}

func (s *Line_formatted_textContext) EXPRESSION_START(i int) antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserEXPRESSION_START, i)
}

func (s *Line_formatted_textContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *Line_formatted_textContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Line_formatted_textContext) AllEXPRESSION_END() []antlr.TerminalNode {
	return s.GetTokens(YarnSpinnerParserEXPRESSION_END)
}

func (s *Line_formatted_textContext) EXPRESSION_END(i int) antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserEXPRESSION_END, i)
}

func (s *Line_formatted_textContext) AllTEXT() []antlr.TerminalNode {
	return s.GetTokens(YarnSpinnerParserTEXT)
}

func (s *Line_formatted_textContext) TEXT(i int) antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserTEXT, i)
}

func (s *Line_formatted_textContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Line_formatted_textContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Line_formatted_textContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterLine_formatted_text(s)
	}
}

func (s *Line_formatted_textContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitLine_formatted_text(s)
	}
}

func (p *YarnSpinnerParser) Line_formatted_text() (localctx ILine_formatted_textContext) {
	localctx = NewLine_formatted_textContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, YarnSpinnerParserRULE_line_formatted_text)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == YarnSpinnerParserEXPRESSION_START || _la == YarnSpinnerParserTEXT {
		p.SetState(178)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case YarnSpinnerParserTEXT:
			p.SetState(170)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = 1
			for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				switch _alt {
				case 1:
					{
						p.SetState(169)
						p.Match(YarnSpinnerParserTEXT)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				default:
					p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
					goto errorExit
				}

				p.SetState(172)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
				if p.HasError() {
					goto errorExit
				}
			}

		case YarnSpinnerParserEXPRESSION_START:
			{
				p.SetState(174)
				p.Match(YarnSpinnerParserEXPRESSION_START)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(175)
				p.expression(0)
			}
			{
				p.SetState(176)
				p.Match(YarnSpinnerParserEXPRESSION_END)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(180)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IHashtagContext is an interface to support dynamic dispatch.
type IHashtagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetInnerText returns the text token.
	GetInnerText() antlr.Token

	// SetText sets the text token.
	SetText(antlr.Token)

	// Getter signatures
	HASHTAG() antlr.TerminalNode
	HASHTAG_TEXT() antlr.TerminalNode

	// IsHashtagContext differentiates from other interfaces.
	IsHashtagContext()
}

type HashtagContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	text   antlr.Token
}

func NewEmptyHashtagContext() *HashtagContext {
	var p = new(HashtagContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_hashtag
	return p
}

func InitEmptyHashtagContext(p *HashtagContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_hashtag
}

func (*HashtagContext) IsHashtagContext() {}

func NewHashtagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HashtagContext {
	var p = new(HashtagContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YarnSpinnerParserRULE_hashtag

	return p
}

func (s *HashtagContext) GetParser() antlr.Parser { return s.parser }

func (s *HashtagContext) GetInnerText() antlr.Token { return s.text }

func (s *HashtagContext) SetText(v antlr.Token) { s.text = v }

func (s *HashtagContext) HASHTAG() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserHASHTAG, 0)
}

func (s *HashtagContext) HASHTAG_TEXT() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserHASHTAG_TEXT, 0)
}

func (s *HashtagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HashtagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HashtagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterHashtag(s)
	}
}

func (s *HashtagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitHashtag(s)
	}
}

func (p *YarnSpinnerParser) Hashtag() (localctx IHashtagContext) {
	localctx = NewHashtagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, YarnSpinnerParserRULE_hashtag)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(182)
		p.Match(YarnSpinnerParserHASHTAG)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(183)

		var _m = p.Match(YarnSpinnerParserHASHTAG_TEXT)

		localctx.(*HashtagContext).text = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILine_conditionContext is an interface to support dynamic dispatch.
type ILine_conditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLine_conditionContext differentiates from other interfaces.
	IsLine_conditionContext()
}

type Line_conditionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLine_conditionContext() *Line_conditionContext {
	var p = new(Line_conditionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_line_condition
	return p
}

func InitEmptyLine_conditionContext(p *Line_conditionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_line_condition
}

func (*Line_conditionContext) IsLine_conditionContext() {}

func NewLine_conditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Line_conditionContext {
	var p = new(Line_conditionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YarnSpinnerParserRULE_line_condition

	return p
}

func (s *Line_conditionContext) GetParser() antlr.Parser { return s.parser }

func (s *Line_conditionContext) CopyAll(ctx *Line_conditionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Line_conditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Line_conditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type LineOnceConditionContext struct {
	Line_conditionContext
}

func NewLineOnceConditionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LineOnceConditionContext {
	var p = new(LineOnceConditionContext)

	InitEmptyLine_conditionContext(&p.Line_conditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Line_conditionContext))

	return p
}

func (s *LineOnceConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineOnceConditionContext) COMMAND_START() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_START, 0)
}

func (s *LineOnceConditionContext) COMMAND_ONCE() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_ONCE, 0)
}

func (s *LineOnceConditionContext) COMMAND_END() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_END, 0)
}

func (s *LineOnceConditionContext) COMMAND_IF() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_IF, 0)
}

func (s *LineOnceConditionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LineOnceConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterLineOnceCondition(s)
	}
}

func (s *LineOnceConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitLineOnceCondition(s)
	}
}

type LineConditionContext struct {
	Line_conditionContext
}

func NewLineConditionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LineConditionContext {
	var p = new(LineConditionContext)

	InitEmptyLine_conditionContext(&p.Line_conditionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Line_conditionContext))

	return p
}

func (s *LineConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineConditionContext) COMMAND_START() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_START, 0)
}

func (s *LineConditionContext) COMMAND_IF() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_IF, 0)
}

func (s *LineConditionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LineConditionContext) COMMAND_END() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_END, 0)
}

func (s *LineConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterLineCondition(s)
	}
}

func (s *LineConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitLineCondition(s)
	}
}

func (p *YarnSpinnerParser) Line_condition() (localctx ILine_conditionContext) {
	localctx = NewLine_conditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, YarnSpinnerParserRULE_line_condition)
	var _la int

	p.SetState(197)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		localctx = NewLineConditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(185)
			p.Match(YarnSpinnerParserCOMMAND_START)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(186)
			p.Match(YarnSpinnerParserCOMMAND_IF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(187)
			p.expression(0)
		}
		{
			p.SetState(188)
			p.Match(YarnSpinnerParserCOMMAND_END)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewLineOnceConditionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(190)
			p.Match(YarnSpinnerParserCOMMAND_START)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(191)
			p.Match(YarnSpinnerParserCOMMAND_ONCE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(194)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == YarnSpinnerParserCOMMAND_IF {
			{
				p.SetState(192)
				p.Match(YarnSpinnerParserCOMMAND_IF)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(193)
				p.expression(0)
			}

		}
		{
			p.SetState(196)
			p.Match(YarnSpinnerParserCOMMAND_END)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YarnSpinnerParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyAll(ctx *ExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExpParensContext struct {
	ExpressionContext
}

func NewExpParensContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpParensContext {
	var p = new(ExpParensContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpParensContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpParensContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserLPAREN, 0)
}

func (s *ExpParensContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpParensContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserRPAREN, 0)
}

func (s *ExpParensContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterExpParens(s)
	}
}

func (s *ExpParensContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitExpParens(s)
	}
}

type ExpMultDivModContext struct {
	ExpressionContext
	op antlr.Token
}

func NewExpMultDivModContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpMultDivModContext {
	var p = new(ExpMultDivModContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpMultDivModContext) GetOp() antlr.Token { return s.op }

func (s *ExpMultDivModContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExpMultDivModContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpMultDivModContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpMultDivModContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpMultDivModContext) OPERATOR_MATHS_MULTIPLICATION() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserOPERATOR_MATHS_MULTIPLICATION, 0)
}

func (s *ExpMultDivModContext) OPERATOR_MATHS_DIVISION() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserOPERATOR_MATHS_DIVISION, 0)
}

func (s *ExpMultDivModContext) OPERATOR_MATHS_MODULUS() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserOPERATOR_MATHS_MODULUS, 0)
}

func (s *ExpMultDivModContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterExpMultDivMod(s)
	}
}

func (s *ExpMultDivModContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitExpMultDivMod(s)
	}
}

type ExpComparisonContext struct {
	ExpressionContext
	op antlr.Token
}

func NewExpComparisonContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpComparisonContext {
	var p = new(ExpComparisonContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpComparisonContext) GetOp() antlr.Token { return s.op }

func (s *ExpComparisonContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExpComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpComparisonContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpComparisonContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpComparisonContext) OPERATOR_LOGICAL_LESS_THAN_EQUALS() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserOPERATOR_LOGICAL_LESS_THAN_EQUALS, 0)
}

func (s *ExpComparisonContext) OPERATOR_LOGICAL_GREATER_THAN_EQUALS() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserOPERATOR_LOGICAL_GREATER_THAN_EQUALS, 0)
}

func (s *ExpComparisonContext) OPERATOR_LOGICAL_LESS() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserOPERATOR_LOGICAL_LESS, 0)
}

func (s *ExpComparisonContext) OPERATOR_LOGICAL_GREATER() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserOPERATOR_LOGICAL_GREATER, 0)
}

func (s *ExpComparisonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterExpComparison(s)
	}
}

func (s *ExpComparisonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitExpComparison(s)
	}
}

type ExpNegativeContext struct {
	ExpressionContext
	op antlr.Token
}

func NewExpNegativeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpNegativeContext {
	var p = new(ExpNegativeContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpNegativeContext) GetOp() antlr.Token { return s.op }

func (s *ExpNegativeContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExpNegativeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpNegativeContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpNegativeContext) OPERATOR_MATHS_SUBTRACTION() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserOPERATOR_MATHS_SUBTRACTION, 0)
}

func (s *ExpNegativeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterExpNegative(s)
	}
}

func (s *ExpNegativeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitExpNegative(s)
	}
}

type ExpAndOrXorContext struct {
	ExpressionContext
	op antlr.Token
}

func NewExpAndOrXorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpAndOrXorContext {
	var p = new(ExpAndOrXorContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpAndOrXorContext) GetOp() antlr.Token { return s.op }

func (s *ExpAndOrXorContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExpAndOrXorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpAndOrXorContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpAndOrXorContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpAndOrXorContext) OPERATOR_LOGICAL_AND() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserOPERATOR_LOGICAL_AND, 0)
}

func (s *ExpAndOrXorContext) OPERATOR_LOGICAL_OR() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserOPERATOR_LOGICAL_OR, 0)
}

func (s *ExpAndOrXorContext) OPERATOR_LOGICAL_XOR() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserOPERATOR_LOGICAL_XOR, 0)
}

func (s *ExpAndOrXorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterExpAndOrXor(s)
	}
}

func (s *ExpAndOrXorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitExpAndOrXor(s)
	}
}

type ExpAddSubContext struct {
	ExpressionContext
	op antlr.Token
}

func NewExpAddSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpAddSubContext {
	var p = new(ExpAddSubContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpAddSubContext) GetOp() antlr.Token { return s.op }

func (s *ExpAddSubContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExpAddSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpAddSubContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpAddSubContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpAddSubContext) OPERATOR_MATHS_ADDITION() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserOPERATOR_MATHS_ADDITION, 0)
}

func (s *ExpAddSubContext) OPERATOR_MATHS_SUBTRACTION() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserOPERATOR_MATHS_SUBTRACTION, 0)
}

func (s *ExpAddSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterExpAddSub(s)
	}
}

func (s *ExpAddSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitExpAddSub(s)
	}
}

type ExpNotContext struct {
	ExpressionContext
	op antlr.Token
}

func NewExpNotContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpNotContext {
	var p = new(ExpNotContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpNotContext) GetOp() antlr.Token { return s.op }

func (s *ExpNotContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExpNotContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpNotContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpNotContext) OPERATOR_LOGICAL_NOT() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserOPERATOR_LOGICAL_NOT, 0)
}

func (s *ExpNotContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterExpNot(s)
	}
}

func (s *ExpNotContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitExpNot(s)
	}
}

type ExpValueContext struct {
	ExpressionContext
}

func NewExpValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpValueContext {
	var p = new(ExpValueContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpValueContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ExpValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterExpValue(s)
	}
}

func (s *ExpValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitExpValue(s)
	}
}

type ExpEqualityContext struct {
	ExpressionContext
	op antlr.Token
}

func NewExpEqualityContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpEqualityContext {
	var p = new(ExpEqualityContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpEqualityContext) GetOp() antlr.Token { return s.op }

func (s *ExpEqualityContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExpEqualityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpEqualityContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpEqualityContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpEqualityContext) OPERATOR_LOGICAL_EQUALS() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserOPERATOR_LOGICAL_EQUALS, 0)
}

func (s *ExpEqualityContext) OPERATOR_LOGICAL_NOT_EQUALS() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserOPERATOR_LOGICAL_NOT_EQUALS, 0)
}

func (s *ExpEqualityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterExpEquality(s)
	}
}

func (s *ExpEqualityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitExpEquality(s)
	}
}

func (p *YarnSpinnerParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *YarnSpinnerParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 26
	p.EnterRecursionRule(localctx, 26, YarnSpinnerParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(209)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case YarnSpinnerParserLPAREN:
		localctx = NewExpParensContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(200)
			p.Match(YarnSpinnerParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(201)
			p.expression(0)
		}
		{
			p.SetState(202)
			p.Match(YarnSpinnerParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case YarnSpinnerParserOPERATOR_MATHS_SUBTRACTION:
		localctx = NewExpNegativeContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(204)

			var _m = p.Match(YarnSpinnerParserOPERATOR_MATHS_SUBTRACTION)

			localctx.(*ExpNegativeContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(205)
			p.expression(8)
		}

	case YarnSpinnerParserOPERATOR_LOGICAL_NOT:
		localctx = NewExpNotContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(206)

			var _m = p.Match(YarnSpinnerParserOPERATOR_LOGICAL_NOT)

			localctx.(*ExpNotContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(207)
			p.expression(7)
		}

	case YarnSpinnerParserKEYWORD_TRUE, YarnSpinnerParserKEYWORD_FALSE, YarnSpinnerParserSTRING, YarnSpinnerParserFUNC_ID, YarnSpinnerParserVAR_ID, YarnSpinnerParserDOT, YarnSpinnerParserNUMBER:
		localctx = NewExpValueContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(208)
			p.Value()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(228)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(226)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpMultDivModContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, YarnSpinnerParserRULE_expression)
				p.SetState(211)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(212)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpMultDivModContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&126100789566373888) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpMultDivModContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(213)
					p.expression(7)
				}

			case 2:
				localctx = NewExpAddSubContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, YarnSpinnerParserRULE_expression)
				p.SetState(214)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(215)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpAddSubContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == YarnSpinnerParserOPERATOR_MATHS_ADDITION || _la == YarnSpinnerParserOPERATOR_MATHS_SUBTRACTION) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpAddSubContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(216)
					p.expression(6)
				}

			case 3:
				localctx = NewExpComparisonContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, YarnSpinnerParserRULE_expression)
				p.SetState(217)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(218)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpComparisonContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3710851743744) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpComparisonContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(219)
					p.expression(5)
				}

			case 4:
				localctx = NewExpEqualityContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, YarnSpinnerParserRULE_expression)
				p.SetState(220)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(221)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpEqualityContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == YarnSpinnerParserOPERATOR_LOGICAL_EQUALS || _la == YarnSpinnerParserOPERATOR_LOGICAL_NOT_EQUALS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpEqualityContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(222)
					p.expression(4)
				}

			case 5:
				localctx = NewExpAndOrXorContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, YarnSpinnerParserRULE_expression)
				p.SetState(223)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(224)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpAndOrXorContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&61572651155456) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpAndOrXorContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(225)
					p.expression(3)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(230)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YarnSpinnerParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) CopyAll(ctx *ValueContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ValueNumberContext struct {
	ValueContext
}

func NewValueNumberContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValueNumberContext {
	var p = new(ValueNumberContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *ValueNumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueNumberContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserNUMBER, 0)
}

func (s *ValueNumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterValueNumber(s)
	}
}

func (s *ValueNumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitValueNumber(s)
	}
}

type ValueTrueContext struct {
	ValueContext
}

func NewValueTrueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValueTrueContext {
	var p = new(ValueTrueContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *ValueTrueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueTrueContext) KEYWORD_TRUE() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserKEYWORD_TRUE, 0)
}

func (s *ValueTrueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterValueTrue(s)
	}
}

func (s *ValueTrueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitValueTrue(s)
	}
}

type ValueFalseContext struct {
	ValueContext
}

func NewValueFalseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValueFalseContext {
	var p = new(ValueFalseContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *ValueFalseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueFalseContext) KEYWORD_FALSE() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserKEYWORD_FALSE, 0)
}

func (s *ValueFalseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterValueFalse(s)
	}
}

func (s *ValueFalseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitValueFalse(s)
	}
}

type ValueFuncContext struct {
	ValueContext
}

func NewValueFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValueFuncContext {
	var p = new(ValueFuncContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *ValueFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueFuncContext) Function_call() IFunction_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_callContext)
}

func (s *ValueFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterValueFunc(s)
	}
}

func (s *ValueFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitValueFunc(s)
	}
}

type ValueVarContext struct {
	ValueContext
}

func NewValueVarContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValueVarContext {
	var p = new(ValueVarContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *ValueVarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueVarContext) Variable() IVariableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *ValueVarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterValueVar(s)
	}
}

func (s *ValueVarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitValueVar(s)
	}
}

type ValueStringContext struct {
	ValueContext
}

func NewValueStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValueStringContext {
	var p = new(ValueStringContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *ValueStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueStringContext) STRING() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserSTRING, 0)
}

func (s *ValueStringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterValueString(s)
	}
}

func (s *ValueStringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitValueString(s)
	}
}

type ValueTypeMemberReferenceContext struct {
	ValueContext
}

func NewValueTypeMemberReferenceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValueTypeMemberReferenceContext {
	var p = new(ValueTypeMemberReferenceContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *ValueTypeMemberReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueTypeMemberReferenceContext) TypeMemberReference() ITypeMemberReferenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeMemberReferenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeMemberReferenceContext)
}

func (s *ValueTypeMemberReferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterValueTypeMemberReference(s)
	}
}

func (s *ValueTypeMemberReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitValueTypeMemberReference(s)
	}
}

func (p *YarnSpinnerParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, YarnSpinnerParserRULE_value)
	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		localctx = NewValueNumberContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(231)
			p.Match(YarnSpinnerParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewValueTrueContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(232)
			p.Match(YarnSpinnerParserKEYWORD_TRUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewValueFalseContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(233)
			p.Match(YarnSpinnerParserKEYWORD_FALSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewValueVarContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(234)
			p.Variable()
		}

	case 5:
		localctx = NewValueStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(235)
			p.Match(YarnSpinnerParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewValueFuncContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(236)
			p.Function_call()
		}

	case 7:
		localctx = NewValueTypeMemberReferenceContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(237)
			p.TypeMemberReference()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VAR_ID() antlr.TerminalNode

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_variable
	return p
}

func InitEmptyVariableContext(p *VariableContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_variable
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YarnSpinnerParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) VAR_ID() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserVAR_ID, 0)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (p *YarnSpinnerParser) Variable() (localctx IVariableContext) {
	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, YarnSpinnerParserRULE_variable)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(240)
		p.Match(YarnSpinnerParserVAR_ID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunction_callContext is an interface to support dynamic dispatch.
type IFunction_callContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNC_ID() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsFunction_callContext differentiates from other interfaces.
	IsFunction_callContext()
}

type Function_callContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_callContext() *Function_callContext {
	var p = new(Function_callContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_function_call
	return p
}

func InitEmptyFunction_callContext(p *Function_callContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_function_call
}

func (*Function_callContext) IsFunction_callContext() {}

func NewFunction_callContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_callContext {
	var p = new(Function_callContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YarnSpinnerParserRULE_function_call

	return p
}

func (s *Function_callContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_callContext) FUNC_ID() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserFUNC_ID, 0)
}

func (s *Function_callContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserLPAREN, 0)
}

func (s *Function_callContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserRPAREN, 0)
}

func (s *Function_callContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *Function_callContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Function_callContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(YarnSpinnerParserCOMMA)
}

func (s *Function_callContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMA, i)
}

func (s *Function_callContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_callContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_callContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterFunction_call(s)
	}
}

func (s *Function_callContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitFunction_call(s)
	}
}

func (p *YarnSpinnerParser) Function_call() (localctx IFunction_callContext) {
	localctx = NewFunction_callContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, YarnSpinnerParserRULE_function_call)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(242)
		p.Match(YarnSpinnerParserFUNC_ID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(243)
		p.Match(YarnSpinnerParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-33)) & ^0x3f) == 0 && ((int64(1)<<(_la-33))&15855525891) != 0 {
		{
			p.SetState(244)
			p.expression(0)
		}

	}
	p.SetState(251)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == YarnSpinnerParserCOMMA {
		{
			p.SetState(247)
			p.Match(YarnSpinnerParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(248)
			p.expression(0)
		}

		p.SetState(253)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(254)
		p.Match(YarnSpinnerParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeMemberReferenceContext is an interface to support dynamic dispatch.
type ITypeMemberReferenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTypeName returns the typeName token.
	GetTypeName() antlr.Token

	// GetMemberName returns the memberName token.
	GetMemberName() antlr.Token

	// SetTypeName sets the typeName token.
	SetTypeName(antlr.Token)

	// SetMemberName sets the memberName token.
	SetMemberName(antlr.Token)

	// Getter signatures
	DOT() antlr.TerminalNode
	AllFUNC_ID() []antlr.TerminalNode
	FUNC_ID(i int) antlr.TerminalNode

	// IsTypeMemberReferenceContext differentiates from other interfaces.
	IsTypeMemberReferenceContext()
}

type TypeMemberReferenceContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	typeName   antlr.Token
	memberName antlr.Token
}

func NewEmptyTypeMemberReferenceContext() *TypeMemberReferenceContext {
	var p = new(TypeMemberReferenceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_typeMemberReference
	return p
}

func InitEmptyTypeMemberReferenceContext(p *TypeMemberReferenceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_typeMemberReference
}

func (*TypeMemberReferenceContext) IsTypeMemberReferenceContext() {}

func NewTypeMemberReferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeMemberReferenceContext {
	var p = new(TypeMemberReferenceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YarnSpinnerParserRULE_typeMemberReference

	return p
}

func (s *TypeMemberReferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeMemberReferenceContext) GetTypeName() antlr.Token { return s.typeName }

func (s *TypeMemberReferenceContext) GetMemberName() antlr.Token { return s.memberName }

func (s *TypeMemberReferenceContext) SetTypeName(v antlr.Token) { s.typeName = v }

func (s *TypeMemberReferenceContext) SetMemberName(v antlr.Token) { s.memberName = v }

func (s *TypeMemberReferenceContext) DOT() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserDOT, 0)
}

func (s *TypeMemberReferenceContext) AllFUNC_ID() []antlr.TerminalNode {
	return s.GetTokens(YarnSpinnerParserFUNC_ID)
}

func (s *TypeMemberReferenceContext) FUNC_ID(i int) antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserFUNC_ID, i)
}

func (s *TypeMemberReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeMemberReferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeMemberReferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterTypeMemberReference(s)
	}
}

func (s *TypeMemberReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitTypeMemberReference(s)
	}
}

func (p *YarnSpinnerParser) TypeMemberReference() (localctx ITypeMemberReferenceContext) {
	localctx = NewTypeMemberReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, YarnSpinnerParserRULE_typeMemberReference)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(257)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == YarnSpinnerParserFUNC_ID {
		{
			p.SetState(256)

			var _m = p.Match(YarnSpinnerParserFUNC_ID)

			localctx.(*TypeMemberReferenceContext).typeName = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(259)
		p.Match(YarnSpinnerParserDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(260)

		var _m = p.Match(YarnSpinnerParserFUNC_ID)

		localctx.(*TypeMemberReferenceContext).memberName = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIf_statementContext is an interface to support dynamic dispatch.
type IIf_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	If_clause() IIf_clauseContext
	COMMAND_START() antlr.TerminalNode
	COMMAND_ENDIF() antlr.TerminalNode
	COMMAND_END() antlr.TerminalNode
	AllElse_if_clause() []IElse_if_clauseContext
	Else_if_clause(i int) IElse_if_clauseContext
	Else_clause() IElse_clauseContext

	// IsIf_statementContext differentiates from other interfaces.
	IsIf_statementContext()
}

type If_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_statementContext() *If_statementContext {
	var p = new(If_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_if_statement
	return p
}

func InitEmptyIf_statementContext(p *If_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_if_statement
}

func (*If_statementContext) IsIf_statementContext() {}

func NewIf_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_statementContext {
	var p = new(If_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YarnSpinnerParserRULE_if_statement

	return p
}

func (s *If_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *If_statementContext) If_clause() IIf_clauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_clauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIf_clauseContext)
}

func (s *If_statementContext) COMMAND_START() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_START, 0)
}

func (s *If_statementContext) COMMAND_ENDIF() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_ENDIF, 0)
}

func (s *If_statementContext) COMMAND_END() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_END, 0)
}

func (s *If_statementContext) AllElse_if_clause() []IElse_if_clauseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IElse_if_clauseContext); ok {
			len++
		}
	}

	tst := make([]IElse_if_clauseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IElse_if_clauseContext); ok {
			tst[i] = t.(IElse_if_clauseContext)
			i++
		}
	}

	return tst
}

func (s *If_statementContext) Else_if_clause(i int) IElse_if_clauseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElse_if_clauseContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElse_if_clauseContext)
}

func (s *If_statementContext) Else_clause() IElse_clauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElse_clauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElse_clauseContext)
}

func (s *If_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterIf_statement(s)
	}
}

func (s *If_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitIf_statement(s)
	}
}

func (p *YarnSpinnerParser) If_statement() (localctx IIf_statementContext) {
	localctx = NewIf_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, YarnSpinnerParserRULE_if_statement)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(262)
		p.If_clause()
	}
	p.SetState(266)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(263)
				p.Else_if_clause()
			}

		}
		p.SetState(268)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(270)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(269)
			p.Else_clause()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(272)
		p.Match(YarnSpinnerParserCOMMAND_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(273)
		p.Match(YarnSpinnerParserCOMMAND_ENDIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(274)
		p.Match(YarnSpinnerParserCOMMAND_END)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIf_clauseContext is an interface to support dynamic dispatch.
type IIf_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COMMAND_START() antlr.TerminalNode
	COMMAND_IF() antlr.TerminalNode
	Expression() IExpressionContext
	COMMAND_END() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsIf_clauseContext differentiates from other interfaces.
	IsIf_clauseContext()
}

type If_clauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_clauseContext() *If_clauseContext {
	var p = new(If_clauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_if_clause
	return p
}

func InitEmptyIf_clauseContext(p *If_clauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_if_clause
}

func (*If_clauseContext) IsIf_clauseContext() {}

func NewIf_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_clauseContext {
	var p = new(If_clauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YarnSpinnerParserRULE_if_clause

	return p
}

func (s *If_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *If_clauseContext) COMMAND_START() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_START, 0)
}

func (s *If_clauseContext) COMMAND_IF() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_IF, 0)
}

func (s *If_clauseContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *If_clauseContext) COMMAND_END() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_END, 0)
}

func (s *If_clauseContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *If_clauseContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *If_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_clauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterIf_clause(s)
	}
}

func (s *If_clauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitIf_clause(s)
	}
}

func (p *YarnSpinnerParser) If_clause() (localctx IIf_clauseContext) {
	localctx = NewIf_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, YarnSpinnerParserRULE_if_clause)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(276)
		p.Match(YarnSpinnerParserCOMMAND_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(277)
		p.Match(YarnSpinnerParserCOMMAND_IF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(278)
		p.expression(0)
	}
	{
		p.SetState(279)
		p.Match(YarnSpinnerParserCOMMAND_END)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(283)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(280)
				p.Statement()
			}

		}
		p.SetState(285)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElse_if_clauseContext is an interface to support dynamic dispatch.
type IElse_if_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COMMAND_START() antlr.TerminalNode
	COMMAND_ELSEIF() antlr.TerminalNode
	Expression() IExpressionContext
	COMMAND_END() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsElse_if_clauseContext differentiates from other interfaces.
	IsElse_if_clauseContext()
}

type Else_if_clauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElse_if_clauseContext() *Else_if_clauseContext {
	var p = new(Else_if_clauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_else_if_clause
	return p
}

func InitEmptyElse_if_clauseContext(p *Else_if_clauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_else_if_clause
}

func (*Else_if_clauseContext) IsElse_if_clauseContext() {}

func NewElse_if_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Else_if_clauseContext {
	var p = new(Else_if_clauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YarnSpinnerParserRULE_else_if_clause

	return p
}

func (s *Else_if_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *Else_if_clauseContext) COMMAND_START() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_START, 0)
}

func (s *Else_if_clauseContext) COMMAND_ELSEIF() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_ELSEIF, 0)
}

func (s *Else_if_clauseContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Else_if_clauseContext) COMMAND_END() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_END, 0)
}

func (s *Else_if_clauseContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *Else_if_clauseContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *Else_if_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Else_if_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Else_if_clauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterElse_if_clause(s)
	}
}

func (s *Else_if_clauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitElse_if_clause(s)
	}
}

func (p *YarnSpinnerParser) Else_if_clause() (localctx IElse_if_clauseContext) {
	localctx = NewElse_if_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, YarnSpinnerParserRULE_else_if_clause)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(286)
		p.Match(YarnSpinnerParserCOMMAND_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(287)
		p.Match(YarnSpinnerParserCOMMAND_ELSEIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(288)
		p.expression(0)
	}
	{
		p.SetState(289)
		p.Match(YarnSpinnerParserCOMMAND_END)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(293)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(290)
				p.Statement()
			}

		}
		p.SetState(295)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElse_clauseContext is an interface to support dynamic dispatch.
type IElse_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COMMAND_START() antlr.TerminalNode
	COMMAND_ELSE() antlr.TerminalNode
	COMMAND_END() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsElse_clauseContext differentiates from other interfaces.
	IsElse_clauseContext()
}

type Else_clauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElse_clauseContext() *Else_clauseContext {
	var p = new(Else_clauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_else_clause
	return p
}

func InitEmptyElse_clauseContext(p *Else_clauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_else_clause
}

func (*Else_clauseContext) IsElse_clauseContext() {}

func NewElse_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Else_clauseContext {
	var p = new(Else_clauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YarnSpinnerParserRULE_else_clause

	return p
}

func (s *Else_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *Else_clauseContext) COMMAND_START() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_START, 0)
}

func (s *Else_clauseContext) COMMAND_ELSE() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_ELSE, 0)
}

func (s *Else_clauseContext) COMMAND_END() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_END, 0)
}

func (s *Else_clauseContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *Else_clauseContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *Else_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Else_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Else_clauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterElse_clause(s)
	}
}

func (s *Else_clauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitElse_clause(s)
	}
}

func (p *YarnSpinnerParser) Else_clause() (localctx IElse_clauseContext) {
	localctx = NewElse_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, YarnSpinnerParserRULE_else_clause)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(296)
		p.Match(YarnSpinnerParserCOMMAND_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(297)
		p.Match(YarnSpinnerParserCOMMAND_ELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(298)
		p.Match(YarnSpinnerParserCOMMAND_END)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(302)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(299)
				p.Statement()
			}

		}
		p.SetState(304)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISet_statementContext is an interface to support dynamic dispatch.
type ISet_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Getter signatures
	COMMAND_START() antlr.TerminalNode
	COMMAND_SET() antlr.TerminalNode
	Variable() IVariableContext
	Expression() IExpressionContext
	COMMAND_END() antlr.TerminalNode
	OPERATOR_ASSIGNMENT() antlr.TerminalNode
	OPERATOR_MATHS_MULTIPLICATION_EQUALS() antlr.TerminalNode
	OPERATOR_MATHS_DIVISION_EQUALS() antlr.TerminalNode
	OPERATOR_MATHS_MODULUS_EQUALS() antlr.TerminalNode
	OPERATOR_MATHS_ADDITION_EQUALS() antlr.TerminalNode
	OPERATOR_MATHS_SUBTRACTION_EQUALS() antlr.TerminalNode

	// IsSet_statementContext differentiates from other interfaces.
	IsSet_statementContext()
}

type Set_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptySet_statementContext() *Set_statementContext {
	var p = new(Set_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_set_statement
	return p
}

func InitEmptySet_statementContext(p *Set_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_set_statement
}

func (*Set_statementContext) IsSet_statementContext() {}

func NewSet_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Set_statementContext {
	var p = new(Set_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YarnSpinnerParserRULE_set_statement

	return p
}

func (s *Set_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Set_statementContext) GetOp() antlr.Token { return s.op }

func (s *Set_statementContext) SetOp(v antlr.Token) { s.op = v }

func (s *Set_statementContext) COMMAND_START() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_START, 0)
}

func (s *Set_statementContext) COMMAND_SET() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_SET, 0)
}

func (s *Set_statementContext) Variable() IVariableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *Set_statementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Set_statementContext) COMMAND_END() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_END, 0)
}

func (s *Set_statementContext) OPERATOR_ASSIGNMENT() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserOPERATOR_ASSIGNMENT, 0)
}

func (s *Set_statementContext) OPERATOR_MATHS_MULTIPLICATION_EQUALS() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserOPERATOR_MATHS_MULTIPLICATION_EQUALS, 0)
}

func (s *Set_statementContext) OPERATOR_MATHS_DIVISION_EQUALS() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserOPERATOR_MATHS_DIVISION_EQUALS, 0)
}

func (s *Set_statementContext) OPERATOR_MATHS_MODULUS_EQUALS() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserOPERATOR_MATHS_MODULUS_EQUALS, 0)
}

func (s *Set_statementContext) OPERATOR_MATHS_ADDITION_EQUALS() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserOPERATOR_MATHS_ADDITION_EQUALS, 0)
}

func (s *Set_statementContext) OPERATOR_MATHS_SUBTRACTION_EQUALS() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserOPERATOR_MATHS_SUBTRACTION_EQUALS, 0)
}

func (s *Set_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Set_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Set_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterSet_statement(s)
	}
}

func (s *Set_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitSet_statement(s)
	}
}

func (p *YarnSpinnerParser) Set_statement() (localctx ISet_statementContext) {
	localctx = NewSet_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, YarnSpinnerParserRULE_set_statement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(305)
		p.Match(YarnSpinnerParserCOMMAND_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(306)
		p.Match(YarnSpinnerParserCOMMAND_SET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(307)
		p.Variable()
	}
	{
		p.SetState(308)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*Set_statementContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4362930858491904) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*Set_statementContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(309)
		p.expression(0)
	}
	{
		p.SetState(310)
		p.Match(YarnSpinnerParserCOMMAND_END)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICall_statementContext is an interface to support dynamic dispatch.
type ICall_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COMMAND_START() antlr.TerminalNode
	COMMAND_CALL() antlr.TerminalNode
	Function_call() IFunction_callContext
	COMMAND_END() antlr.TerminalNode

	// IsCall_statementContext differentiates from other interfaces.
	IsCall_statementContext()
}

type Call_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCall_statementContext() *Call_statementContext {
	var p = new(Call_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_call_statement
	return p
}

func InitEmptyCall_statementContext(p *Call_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_call_statement
}

func (*Call_statementContext) IsCall_statementContext() {}

func NewCall_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Call_statementContext {
	var p = new(Call_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YarnSpinnerParserRULE_call_statement

	return p
}

func (s *Call_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Call_statementContext) COMMAND_START() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_START, 0)
}

func (s *Call_statementContext) COMMAND_CALL() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_CALL, 0)
}

func (s *Call_statementContext) Function_call() IFunction_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_callContext)
}

func (s *Call_statementContext) COMMAND_END() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_END, 0)
}

func (s *Call_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Call_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Call_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterCall_statement(s)
	}
}

func (s *Call_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitCall_statement(s)
	}
}

func (p *YarnSpinnerParser) Call_statement() (localctx ICall_statementContext) {
	localctx = NewCall_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, YarnSpinnerParserRULE_call_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(312)
		p.Match(YarnSpinnerParserCOMMAND_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(313)
		p.Match(YarnSpinnerParserCOMMAND_CALL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(314)
		p.Function_call()
	}
	{
		p.SetState(315)
		p.Match(YarnSpinnerParserCOMMAND_END)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICommand_statementContext is an interface to support dynamic dispatch.
type ICommand_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COMMAND_START() antlr.TerminalNode
	Command_formatted_text() ICommand_formatted_textContext
	COMMAND_TEXT_END() antlr.TerminalNode
	AllHashtag() []IHashtagContext
	Hashtag(i int) IHashtagContext

	// IsCommand_statementContext differentiates from other interfaces.
	IsCommand_statementContext()
}

type Command_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommand_statementContext() *Command_statementContext {
	var p = new(Command_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_command_statement
	return p
}

func InitEmptyCommand_statementContext(p *Command_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_command_statement
}

func (*Command_statementContext) IsCommand_statementContext() {}

func NewCommand_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Command_statementContext {
	var p = new(Command_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YarnSpinnerParserRULE_command_statement

	return p
}

func (s *Command_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Command_statementContext) COMMAND_START() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_START, 0)
}

func (s *Command_statementContext) Command_formatted_text() ICommand_formatted_textContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommand_formatted_textContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommand_formatted_textContext)
}

func (s *Command_statementContext) COMMAND_TEXT_END() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_TEXT_END, 0)
}

func (s *Command_statementContext) AllHashtag() []IHashtagContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IHashtagContext); ok {
			len++
		}
	}

	tst := make([]IHashtagContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IHashtagContext); ok {
			tst[i] = t.(IHashtagContext)
			i++
		}
	}

	return tst
}

func (s *Command_statementContext) Hashtag(i int) IHashtagContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHashtagContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHashtagContext)
}

func (s *Command_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Command_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Command_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterCommand_statement(s)
	}
}

func (s *Command_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitCommand_statement(s)
	}
}

func (p *YarnSpinnerParser) Command_statement() (localctx ICommand_statementContext) {
	localctx = NewCommand_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, YarnSpinnerParserRULE_command_statement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(317)
		p.Match(YarnSpinnerParserCOMMAND_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(318)
		p.Command_formatted_text()
	}
	{
		p.SetState(319)
		p.Match(YarnSpinnerParserCOMMAND_TEXT_END)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	p.SetState(323)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == YarnSpinnerParserHASHTAG {
		{
			p.SetState(320)
			p.Hashtag()
		}

		p.SetState(325)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICommand_formatted_textContext is an interface to support dynamic dispatch.
type ICommand_formatted_textContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllCOMMAND_TEXT() []antlr.TerminalNode
	COMMAND_TEXT(i int) antlr.TerminalNode
	AllCOMMAND_EXPRESSION_START() []antlr.TerminalNode
	COMMAND_EXPRESSION_START(i int) antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllEXPRESSION_END() []antlr.TerminalNode
	EXPRESSION_END(i int) antlr.TerminalNode

	// IsCommand_formatted_textContext differentiates from other interfaces.
	IsCommand_formatted_textContext()
}

type Command_formatted_textContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommand_formatted_textContext() *Command_formatted_textContext {
	var p = new(Command_formatted_textContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_command_formatted_text
	return p
}

func InitEmptyCommand_formatted_textContext(p *Command_formatted_textContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_command_formatted_text
}

func (*Command_formatted_textContext) IsCommand_formatted_textContext() {}

func NewCommand_formatted_textContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Command_formatted_textContext {
	var p = new(Command_formatted_textContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YarnSpinnerParserRULE_command_formatted_text

	return p
}

func (s *Command_formatted_textContext) GetParser() antlr.Parser { return s.parser }

func (s *Command_formatted_textContext) AllCOMMAND_TEXT() []antlr.TerminalNode {
	return s.GetTokens(YarnSpinnerParserCOMMAND_TEXT)
}

func (s *Command_formatted_textContext) COMMAND_TEXT(i int) antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_TEXT, i)
}

func (s *Command_formatted_textContext) AllCOMMAND_EXPRESSION_START() []antlr.TerminalNode {
	return s.GetTokens(YarnSpinnerParserCOMMAND_EXPRESSION_START)
}

func (s *Command_formatted_textContext) COMMAND_EXPRESSION_START(i int) antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_EXPRESSION_START, i)
}

func (s *Command_formatted_textContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *Command_formatted_textContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Command_formatted_textContext) AllEXPRESSION_END() []antlr.TerminalNode {
	return s.GetTokens(YarnSpinnerParserEXPRESSION_END)
}

func (s *Command_formatted_textContext) EXPRESSION_END(i int) antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserEXPRESSION_END, i)
}

func (s *Command_formatted_textContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Command_formatted_textContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Command_formatted_textContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterCommand_formatted_text(s)
	}
}

func (s *Command_formatted_textContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitCommand_formatted_text(s)
	}
}

func (p *YarnSpinnerParser) Command_formatted_text() (localctx ICommand_formatted_textContext) {
	localctx = NewCommand_formatted_textContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, YarnSpinnerParserRULE_command_formatted_text)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(333)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == YarnSpinnerParserCOMMAND_EXPRESSION_START || _la == YarnSpinnerParserCOMMAND_TEXT {
		p.SetState(331)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case YarnSpinnerParserCOMMAND_TEXT:
			{
				p.SetState(326)
				p.Match(YarnSpinnerParserCOMMAND_TEXT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case YarnSpinnerParserCOMMAND_EXPRESSION_START:
			{
				p.SetState(327)
				p.Match(YarnSpinnerParserCOMMAND_EXPRESSION_START)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(328)
				p.expression(0)
			}
			{
				p.SetState(329)
				p.Match(YarnSpinnerParserEXPRESSION_END)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(335)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IShortcut_option_statementContext is an interface to support dynamic dispatch.
type IShortcut_option_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllShortcut_option() []IShortcut_optionContext
	Shortcut_option(i int) IShortcut_optionContext
	BLANK_LINE_FOLLOWING_OPTION() antlr.TerminalNode

	// IsShortcut_option_statementContext differentiates from other interfaces.
	IsShortcut_option_statementContext()
}

type Shortcut_option_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShortcut_option_statementContext() *Shortcut_option_statementContext {
	var p = new(Shortcut_option_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_shortcut_option_statement
	return p
}

func InitEmptyShortcut_option_statementContext(p *Shortcut_option_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_shortcut_option_statement
}

func (*Shortcut_option_statementContext) IsShortcut_option_statementContext() {}

func NewShortcut_option_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Shortcut_option_statementContext {
	var p = new(Shortcut_option_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YarnSpinnerParserRULE_shortcut_option_statement

	return p
}

func (s *Shortcut_option_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Shortcut_option_statementContext) AllShortcut_option() []IShortcut_optionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IShortcut_optionContext); ok {
			len++
		}
	}

	tst := make([]IShortcut_optionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IShortcut_optionContext); ok {
			tst[i] = t.(IShortcut_optionContext)
			i++
		}
	}

	return tst
}

func (s *Shortcut_option_statementContext) Shortcut_option(i int) IShortcut_optionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IShortcut_optionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IShortcut_optionContext)
}

func (s *Shortcut_option_statementContext) BLANK_LINE_FOLLOWING_OPTION() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserBLANK_LINE_FOLLOWING_OPTION, 0)
}

func (s *Shortcut_option_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Shortcut_option_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Shortcut_option_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterShortcut_option_statement(s)
	}
}

func (s *Shortcut_option_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitShortcut_option_statement(s)
	}
}

func (p *YarnSpinnerParser) Shortcut_option_statement() (localctx IShortcut_option_statementContext) {
	localctx = NewShortcut_option_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, YarnSpinnerParserRULE_shortcut_option_statement)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(339)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(336)
				p.Shortcut_option()
			}

		}
		p.SetState(341)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

	{
		p.SetState(342)
		p.Shortcut_option()
	}
	p.SetState(344)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == YarnSpinnerParserBLANK_LINE_FOLLOWING_OPTION {
		{
			p.SetState(343)
			p.Match(YarnSpinnerParserBLANK_LINE_FOLLOWING_OPTION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IShortcut_optionContext is an interface to support dynamic dispatch.
type IShortcut_optionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SHORTCUT_ARROW() antlr.TerminalNode
	Line_statement() ILine_statementContext
	INDENT() antlr.TerminalNode
	DEDENT() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsShortcut_optionContext differentiates from other interfaces.
	IsShortcut_optionContext()
}

type Shortcut_optionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShortcut_optionContext() *Shortcut_optionContext {
	var p = new(Shortcut_optionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_shortcut_option
	return p
}

func InitEmptyShortcut_optionContext(p *Shortcut_optionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_shortcut_option
}

func (*Shortcut_optionContext) IsShortcut_optionContext() {}

func NewShortcut_optionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Shortcut_optionContext {
	var p = new(Shortcut_optionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YarnSpinnerParserRULE_shortcut_option

	return p
}

func (s *Shortcut_optionContext) GetParser() antlr.Parser { return s.parser }

func (s *Shortcut_optionContext) SHORTCUT_ARROW() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserSHORTCUT_ARROW, 0)
}

func (s *Shortcut_optionContext) Line_statement() ILine_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILine_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILine_statementContext)
}

func (s *Shortcut_optionContext) INDENT() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserINDENT, 0)
}

func (s *Shortcut_optionContext) DEDENT() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserDEDENT, 0)
}

func (s *Shortcut_optionContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *Shortcut_optionContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *Shortcut_optionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Shortcut_optionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Shortcut_optionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterShortcut_option(s)
	}
}

func (s *Shortcut_optionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitShortcut_option(s)
	}
}

func (p *YarnSpinnerParser) Shortcut_option() (localctx IShortcut_optionContext) {
	localctx = NewShortcut_optionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, YarnSpinnerParserRULE_shortcut_option)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(346)
		p.Match(YarnSpinnerParserSHORTCUT_ARROW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(347)
		p.Line_statement()
	}
	p.SetState(356)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(348)
			p.Match(YarnSpinnerParserINDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(352)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&18743298) != 0 {
			{
				p.SetState(349)
				p.Statement()
			}

			p.SetState(354)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(355)
			p.Match(YarnSpinnerParserDEDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILine_group_statementContext is an interface to support dynamic dispatch.
type ILine_group_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllLine_group_item() []ILine_group_itemContext
	Line_group_item(i int) ILine_group_itemContext
	BLANK_LINE_FOLLOWING_OPTION() antlr.TerminalNode

	// IsLine_group_statementContext differentiates from other interfaces.
	IsLine_group_statementContext()
}

type Line_group_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLine_group_statementContext() *Line_group_statementContext {
	var p = new(Line_group_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_line_group_statement
	return p
}

func InitEmptyLine_group_statementContext(p *Line_group_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_line_group_statement
}

func (*Line_group_statementContext) IsLine_group_statementContext() {}

func NewLine_group_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Line_group_statementContext {
	var p = new(Line_group_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YarnSpinnerParserRULE_line_group_statement

	return p
}

func (s *Line_group_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Line_group_statementContext) AllLine_group_item() []ILine_group_itemContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILine_group_itemContext); ok {
			len++
		}
	}

	tst := make([]ILine_group_itemContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILine_group_itemContext); ok {
			tst[i] = t.(ILine_group_itemContext)
			i++
		}
	}

	return tst
}

func (s *Line_group_statementContext) Line_group_item(i int) ILine_group_itemContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILine_group_itemContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILine_group_itemContext)
}

func (s *Line_group_statementContext) BLANK_LINE_FOLLOWING_OPTION() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserBLANK_LINE_FOLLOWING_OPTION, 0)
}

func (s *Line_group_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Line_group_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Line_group_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterLine_group_statement(s)
	}
}

func (s *Line_group_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitLine_group_statement(s)
	}
}

func (p *YarnSpinnerParser) Line_group_statement() (localctx ILine_group_statementContext) {
	localctx = NewLine_group_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, YarnSpinnerParserRULE_line_group_statement)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(361)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(358)
				p.Line_group_item()
			}

		}
		p.SetState(363)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

	{
		p.SetState(364)
		p.Line_group_item()
	}
	p.SetState(366)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == YarnSpinnerParserBLANK_LINE_FOLLOWING_OPTION {
		{
			p.SetState(365)
			p.Match(YarnSpinnerParserBLANK_LINE_FOLLOWING_OPTION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILine_group_itemContext is an interface to support dynamic dispatch.
type ILine_group_itemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LINE_GROUP_ARROW() antlr.TerminalNode
	Line_statement() ILine_statementContext
	INDENT() antlr.TerminalNode
	DEDENT() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsLine_group_itemContext differentiates from other interfaces.
	IsLine_group_itemContext()
}

type Line_group_itemContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLine_group_itemContext() *Line_group_itemContext {
	var p = new(Line_group_itemContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_line_group_item
	return p
}

func InitEmptyLine_group_itemContext(p *Line_group_itemContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_line_group_item
}

func (*Line_group_itemContext) IsLine_group_itemContext() {}

func NewLine_group_itemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Line_group_itemContext {
	var p = new(Line_group_itemContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YarnSpinnerParserRULE_line_group_item

	return p
}

func (s *Line_group_itemContext) GetParser() antlr.Parser { return s.parser }

func (s *Line_group_itemContext) LINE_GROUP_ARROW() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserLINE_GROUP_ARROW, 0)
}

func (s *Line_group_itemContext) Line_statement() ILine_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILine_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILine_statementContext)
}

func (s *Line_group_itemContext) INDENT() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserINDENT, 0)
}

func (s *Line_group_itemContext) DEDENT() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserDEDENT, 0)
}

func (s *Line_group_itemContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *Line_group_itemContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *Line_group_itemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Line_group_itemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Line_group_itemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterLine_group_item(s)
	}
}

func (s *Line_group_itemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitLine_group_item(s)
	}
}

func (p *YarnSpinnerParser) Line_group_item() (localctx ILine_group_itemContext) {
	localctx = NewLine_group_itemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, YarnSpinnerParserRULE_line_group_item)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(368)
		p.Match(YarnSpinnerParserLINE_GROUP_ARROW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(369)
		p.Line_statement()
	}
	p.SetState(378)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(370)
			p.Match(YarnSpinnerParserINDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(374)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&18743298) != 0 {
			{
				p.SetState(371)
				p.Statement()
			}

			p.SetState(376)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(377)
			p.Match(YarnSpinnerParserDEDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclare_statementContext is an interface to support dynamic dispatch.
type IDeclare_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetType_ returns the type_ token.
	GetType_() antlr.Token

	// SetType_ sets the type_ token.
	SetType_(antlr.Token)

	// Getter signatures
	COMMAND_START() antlr.TerminalNode
	COMMAND_DECLARE() antlr.TerminalNode
	Variable() IVariableContext
	OPERATOR_ASSIGNMENT() antlr.TerminalNode
	Expression() IExpressionContext
	COMMAND_END() antlr.TerminalNode
	EXPRESSION_AS() antlr.TerminalNode
	FUNC_ID() antlr.TerminalNode

	// IsDeclare_statementContext differentiates from other interfaces.
	IsDeclare_statementContext()
}

type Declare_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	type_  antlr.Token
}

func NewEmptyDeclare_statementContext() *Declare_statementContext {
	var p = new(Declare_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_declare_statement
	return p
}

func InitEmptyDeclare_statementContext(p *Declare_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_declare_statement
}

func (*Declare_statementContext) IsDeclare_statementContext() {}

func NewDeclare_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Declare_statementContext {
	var p = new(Declare_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YarnSpinnerParserRULE_declare_statement

	return p
}

func (s *Declare_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Declare_statementContext) GetType_() antlr.Token { return s.type_ }

func (s *Declare_statementContext) SetType_(v antlr.Token) { s.type_ = v }

func (s *Declare_statementContext) COMMAND_START() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_START, 0)
}

func (s *Declare_statementContext) COMMAND_DECLARE() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_DECLARE, 0)
}

func (s *Declare_statementContext) Variable() IVariableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *Declare_statementContext) OPERATOR_ASSIGNMENT() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserOPERATOR_ASSIGNMENT, 0)
}

func (s *Declare_statementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Declare_statementContext) COMMAND_END() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_END, 0)
}

func (s *Declare_statementContext) EXPRESSION_AS() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserEXPRESSION_AS, 0)
}

func (s *Declare_statementContext) FUNC_ID() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserFUNC_ID, 0)
}

func (s *Declare_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declare_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Declare_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterDeclare_statement(s)
	}
}

func (s *Declare_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitDeclare_statement(s)
	}
}

func (p *YarnSpinnerParser) Declare_statement() (localctx IDeclare_statementContext) {
	localctx = NewDeclare_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, YarnSpinnerParserRULE_declare_statement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(380)
		p.Match(YarnSpinnerParserCOMMAND_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(381)
		p.Match(YarnSpinnerParserCOMMAND_DECLARE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(382)
		p.Variable()
	}
	{
		p.SetState(383)
		p.Match(YarnSpinnerParserOPERATOR_ASSIGNMENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(384)
		p.expression(0)
	}
	p.SetState(387)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == YarnSpinnerParserEXPRESSION_AS {
		{
			p.SetState(385)
			p.Match(YarnSpinnerParserEXPRESSION_AS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(386)

			var _m = p.Match(YarnSpinnerParserFUNC_ID)

			localctx.(*Declare_statementContext).type_ = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(389)
		p.Match(YarnSpinnerParserCOMMAND_END)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEnum_statementContext is an interface to support dynamic dispatch.
type IEnum_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// Getter signatures
	AllCOMMAND_START() []antlr.TerminalNode
	COMMAND_START(i int) antlr.TerminalNode
	COMMAND_ENUM() antlr.TerminalNode
	AllCOMMAND_END() []antlr.TerminalNode
	COMMAND_END(i int) antlr.TerminalNode
	COMMAND_ENDENUM() antlr.TerminalNode
	ID() antlr.TerminalNode
	AllEnum_case_statement() []IEnum_case_statementContext
	Enum_case_statement(i int) IEnum_case_statementContext

	// IsEnum_statementContext differentiates from other interfaces.
	IsEnum_statementContext()
}

type Enum_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
}

func NewEmptyEnum_statementContext() *Enum_statementContext {
	var p = new(Enum_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_enum_statement
	return p
}

func InitEmptyEnum_statementContext(p *Enum_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_enum_statement
}

func (*Enum_statementContext) IsEnum_statementContext() {}

func NewEnum_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Enum_statementContext {
	var p = new(Enum_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YarnSpinnerParserRULE_enum_statement

	return p
}

func (s *Enum_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Enum_statementContext) GetName() antlr.Token { return s.name }

func (s *Enum_statementContext) SetName(v antlr.Token) { s.name = v }

func (s *Enum_statementContext) AllCOMMAND_START() []antlr.TerminalNode {
	return s.GetTokens(YarnSpinnerParserCOMMAND_START)
}

func (s *Enum_statementContext) COMMAND_START(i int) antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_START, i)
}

func (s *Enum_statementContext) COMMAND_ENUM() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_ENUM, 0)
}

func (s *Enum_statementContext) AllCOMMAND_END() []antlr.TerminalNode {
	return s.GetTokens(YarnSpinnerParserCOMMAND_END)
}

func (s *Enum_statementContext) COMMAND_END(i int) antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_END, i)
}

func (s *Enum_statementContext) COMMAND_ENDENUM() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_ENDENUM, 0)
}

func (s *Enum_statementContext) ID() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserID, 0)
}

func (s *Enum_statementContext) AllEnum_case_statement() []IEnum_case_statementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEnum_case_statementContext); ok {
			len++
		}
	}

	tst := make([]IEnum_case_statementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEnum_case_statementContext); ok {
			tst[i] = t.(IEnum_case_statementContext)
			i++
		}
	}

	return tst
}

func (s *Enum_statementContext) Enum_case_statement(i int) IEnum_case_statementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnum_case_statementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnum_case_statementContext)
}

func (s *Enum_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Enum_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Enum_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterEnum_statement(s)
	}
}

func (s *Enum_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitEnum_statement(s)
	}
}

func (p *YarnSpinnerParser) Enum_statement() (localctx IEnum_statementContext) {
	localctx = NewEnum_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, YarnSpinnerParserRULE_enum_statement)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(391)
		p.Match(YarnSpinnerParserCOMMAND_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(392)
		p.Match(YarnSpinnerParserCOMMAND_ENUM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(393)

		var _m = p.Match(YarnSpinnerParserID)

		localctx.(*Enum_statementContext).name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(394)
		p.Match(YarnSpinnerParserCOMMAND_END)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(396)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(395)
				p.Enum_case_statement()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(398)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 41, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(400)
		p.Match(YarnSpinnerParserCOMMAND_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(401)
		p.Match(YarnSpinnerParserCOMMAND_ENDENUM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(402)
		p.Match(YarnSpinnerParserCOMMAND_END)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEnum_case_statementContext is an interface to support dynamic dispatch.
type IEnum_case_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// GetRawValue returns the rawValue rule contexts.
	GetRawValue() IValueContext

	// SetRawValue sets the rawValue rule contexts.
	SetRawValue(IValueContext)

	// Getter signatures
	COMMAND_START() antlr.TerminalNode
	COMMAND_CASE() antlr.TerminalNode
	COMMAND_END() antlr.TerminalNode
	FUNC_ID() antlr.TerminalNode
	INDENT() antlr.TerminalNode
	OPERATOR_ASSIGNMENT() antlr.TerminalNode
	DEDENT() antlr.TerminalNode
	Value() IValueContext

	// IsEnum_case_statementContext differentiates from other interfaces.
	IsEnum_case_statementContext()
}

type Enum_case_statementContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	name     antlr.Token
	rawValue IValueContext
}

func NewEmptyEnum_case_statementContext() *Enum_case_statementContext {
	var p = new(Enum_case_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_enum_case_statement
	return p
}

func InitEmptyEnum_case_statementContext(p *Enum_case_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_enum_case_statement
}

func (*Enum_case_statementContext) IsEnum_case_statementContext() {}

func NewEnum_case_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Enum_case_statementContext {
	var p = new(Enum_case_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YarnSpinnerParserRULE_enum_case_statement

	return p
}

func (s *Enum_case_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Enum_case_statementContext) GetName() antlr.Token { return s.name }

func (s *Enum_case_statementContext) SetName(v antlr.Token) { s.name = v }

func (s *Enum_case_statementContext) GetRawValue() IValueContext { return s.rawValue }

func (s *Enum_case_statementContext) SetRawValue(v IValueContext) { s.rawValue = v }

func (s *Enum_case_statementContext) COMMAND_START() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_START, 0)
}

func (s *Enum_case_statementContext) COMMAND_CASE() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_CASE, 0)
}

func (s *Enum_case_statementContext) COMMAND_END() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_END, 0)
}

func (s *Enum_case_statementContext) FUNC_ID() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserFUNC_ID, 0)
}

func (s *Enum_case_statementContext) INDENT() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserINDENT, 0)
}

func (s *Enum_case_statementContext) OPERATOR_ASSIGNMENT() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserOPERATOR_ASSIGNMENT, 0)
}

func (s *Enum_case_statementContext) DEDENT() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserDEDENT, 0)
}

func (s *Enum_case_statementContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *Enum_case_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Enum_case_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Enum_case_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterEnum_case_statement(s)
	}
}

func (s *Enum_case_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitEnum_case_statement(s)
	}
}

func (p *YarnSpinnerParser) Enum_case_statement() (localctx IEnum_case_statementContext) {
	localctx = NewEnum_case_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, YarnSpinnerParserRULE_enum_case_statement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(405)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == YarnSpinnerParserINDENT {
		{
			p.SetState(404)
			p.Match(YarnSpinnerParserINDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(407)
		p.Match(YarnSpinnerParserCOMMAND_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(408)
		p.Match(YarnSpinnerParserCOMMAND_CASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(409)

		var _m = p.Match(YarnSpinnerParserFUNC_ID)

		localctx.(*Enum_case_statementContext).name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(412)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == YarnSpinnerParserOPERATOR_ASSIGNMENT {
		{
			p.SetState(410)
			p.Match(YarnSpinnerParserOPERATOR_ASSIGNMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(411)

			var _x = p.Value()

			localctx.(*Enum_case_statementContext).rawValue = _x
		}

	}
	{
		p.SetState(414)
		p.Match(YarnSpinnerParserCOMMAND_END)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(416)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == YarnSpinnerParserDEDENT {
		{
			p.SetState(415)
			p.Match(YarnSpinnerParserDEDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IJump_statementContext is an interface to support dynamic dispatch.
type IJump_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsJump_statementContext differentiates from other interfaces.
	IsJump_statementContext()
}

type Jump_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJump_statementContext() *Jump_statementContext {
	var p = new(Jump_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_jump_statement
	return p
}

func InitEmptyJump_statementContext(p *Jump_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_jump_statement
}

func (*Jump_statementContext) IsJump_statementContext() {}

func NewJump_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Jump_statementContext {
	var p = new(Jump_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YarnSpinnerParserRULE_jump_statement

	return p
}

func (s *Jump_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Jump_statementContext) CopyAll(ctx *Jump_statementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Jump_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Jump_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type JumpToNodeNameContext struct {
	Jump_statementContext
	destination antlr.Token
}

func NewJumpToNodeNameContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JumpToNodeNameContext {
	var p = new(JumpToNodeNameContext)

	InitEmptyJump_statementContext(&p.Jump_statementContext)
	p.parser = parser
	p.CopyAll(ctx.(*Jump_statementContext))

	return p
}

func (s *JumpToNodeNameContext) GetDestination() antlr.Token { return s.destination }

func (s *JumpToNodeNameContext) SetDestination(v antlr.Token) { s.destination = v }

func (s *JumpToNodeNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JumpToNodeNameContext) COMMAND_START() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_START, 0)
}

func (s *JumpToNodeNameContext) COMMAND_JUMP() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_JUMP, 0)
}

func (s *JumpToNodeNameContext) COMMAND_END() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_END, 0)
}

func (s *JumpToNodeNameContext) ID() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserID, 0)
}

func (s *JumpToNodeNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterJumpToNodeName(s)
	}
}

func (s *JumpToNodeNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitJumpToNodeName(s)
	}
}

type JumpToExpressionContext struct {
	Jump_statementContext
}

func NewJumpToExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JumpToExpressionContext {
	var p = new(JumpToExpressionContext)

	InitEmptyJump_statementContext(&p.Jump_statementContext)
	p.parser = parser
	p.CopyAll(ctx.(*Jump_statementContext))

	return p
}

func (s *JumpToExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JumpToExpressionContext) COMMAND_START() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_START, 0)
}

func (s *JumpToExpressionContext) COMMAND_JUMP() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_JUMP, 0)
}

func (s *JumpToExpressionContext) EXPRESSION_START() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserEXPRESSION_START, 0)
}

func (s *JumpToExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *JumpToExpressionContext) EXPRESSION_END() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserEXPRESSION_END, 0)
}

func (s *JumpToExpressionContext) COMMAND_END() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_END, 0)
}

func (s *JumpToExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterJumpToExpression(s)
	}
}

func (s *JumpToExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitJumpToExpression(s)
	}
}

type DetourToNodeNameContext struct {
	Jump_statementContext
	destination antlr.Token
}

func NewDetourToNodeNameContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DetourToNodeNameContext {
	var p = new(DetourToNodeNameContext)

	InitEmptyJump_statementContext(&p.Jump_statementContext)
	p.parser = parser
	p.CopyAll(ctx.(*Jump_statementContext))

	return p
}

func (s *DetourToNodeNameContext) GetDestination() antlr.Token { return s.destination }

func (s *DetourToNodeNameContext) SetDestination(v antlr.Token) { s.destination = v }

func (s *DetourToNodeNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DetourToNodeNameContext) COMMAND_START() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_START, 0)
}

func (s *DetourToNodeNameContext) COMMAND_DETOUR() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_DETOUR, 0)
}

func (s *DetourToNodeNameContext) COMMAND_END() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_END, 0)
}

func (s *DetourToNodeNameContext) ID() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserID, 0)
}

func (s *DetourToNodeNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterDetourToNodeName(s)
	}
}

func (s *DetourToNodeNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitDetourToNodeName(s)
	}
}

type DetourToExpressionContext struct {
	Jump_statementContext
}

func NewDetourToExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DetourToExpressionContext {
	var p = new(DetourToExpressionContext)

	InitEmptyJump_statementContext(&p.Jump_statementContext)
	p.parser = parser
	p.CopyAll(ctx.(*Jump_statementContext))

	return p
}

func (s *DetourToExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DetourToExpressionContext) COMMAND_START() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_START, 0)
}

func (s *DetourToExpressionContext) COMMAND_DETOUR() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_DETOUR, 0)
}

func (s *DetourToExpressionContext) EXPRESSION_START() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserEXPRESSION_START, 0)
}

func (s *DetourToExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DetourToExpressionContext) EXPRESSION_END() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserEXPRESSION_END, 0)
}

func (s *DetourToExpressionContext) COMMAND_END() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_END, 0)
}

func (s *DetourToExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterDetourToExpression(s)
	}
}

func (s *DetourToExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitDetourToExpression(s)
	}
}

func (p *YarnSpinnerParser) Jump_statement() (localctx IJump_statementContext) {
	localctx = NewJump_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, YarnSpinnerParserRULE_jump_statement)
	p.SetState(440)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 45, p.GetParserRuleContext()) {
	case 1:
		localctx = NewJumpToNodeNameContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(418)
			p.Match(YarnSpinnerParserCOMMAND_START)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(419)
			p.Match(YarnSpinnerParserCOMMAND_JUMP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(420)

			var _m = p.Match(YarnSpinnerParserID)

			localctx.(*JumpToNodeNameContext).destination = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(421)
			p.Match(YarnSpinnerParserCOMMAND_END)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewJumpToExpressionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(422)
			p.Match(YarnSpinnerParserCOMMAND_START)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(423)
			p.Match(YarnSpinnerParserCOMMAND_JUMP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(424)
			p.Match(YarnSpinnerParserEXPRESSION_START)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(425)
			p.expression(0)
		}
		{
			p.SetState(426)
			p.Match(YarnSpinnerParserEXPRESSION_END)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(427)
			p.Match(YarnSpinnerParserCOMMAND_END)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewDetourToNodeNameContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(429)
			p.Match(YarnSpinnerParserCOMMAND_START)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(430)
			p.Match(YarnSpinnerParserCOMMAND_DETOUR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(431)

			var _m = p.Match(YarnSpinnerParserID)

			localctx.(*DetourToNodeNameContext).destination = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(432)
			p.Match(YarnSpinnerParserCOMMAND_END)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewDetourToExpressionContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(433)
			p.Match(YarnSpinnerParserCOMMAND_START)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(434)
			p.Match(YarnSpinnerParserCOMMAND_DETOUR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(435)
			p.Match(YarnSpinnerParserEXPRESSION_START)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(436)
			p.expression(0)
		}
		{
			p.SetState(437)
			p.Match(YarnSpinnerParserEXPRESSION_END)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(438)
			p.Match(YarnSpinnerParserCOMMAND_END)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReturn_statementContext is an interface to support dynamic dispatch.
type IReturn_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COMMAND_START() antlr.TerminalNode
	COMMAND_RETURN() antlr.TerminalNode
	COMMAND_END() antlr.TerminalNode

	// IsReturn_statementContext differentiates from other interfaces.
	IsReturn_statementContext()
}

type Return_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturn_statementContext() *Return_statementContext {
	var p = new(Return_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_return_statement
	return p
}

func InitEmptyReturn_statementContext(p *Return_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_return_statement
}

func (*Return_statementContext) IsReturn_statementContext() {}

func NewReturn_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Return_statementContext {
	var p = new(Return_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YarnSpinnerParserRULE_return_statement

	return p
}

func (s *Return_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Return_statementContext) COMMAND_START() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_START, 0)
}

func (s *Return_statementContext) COMMAND_RETURN() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_RETURN, 0)
}

func (s *Return_statementContext) COMMAND_END() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_END, 0)
}

func (s *Return_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Return_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Return_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterReturn_statement(s)
	}
}

func (s *Return_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitReturn_statement(s)
	}
}

func (p *YarnSpinnerParser) Return_statement() (localctx IReturn_statementContext) {
	localctx = NewReturn_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, YarnSpinnerParserRULE_return_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(442)
		p.Match(YarnSpinnerParserCOMMAND_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(443)
		p.Match(YarnSpinnerParserCOMMAND_RETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(444)
		p.Match(YarnSpinnerParserCOMMAND_END)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOnce_statementContext is an interface to support dynamic dispatch.
type IOnce_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Once_primary_clause() IOnce_primary_clauseContext
	COMMAND_START() antlr.TerminalNode
	COMMAND_ENDONCE() antlr.TerminalNode
	COMMAND_END() antlr.TerminalNode
	Once_alternate_clause() IOnce_alternate_clauseContext

	// IsOnce_statementContext differentiates from other interfaces.
	IsOnce_statementContext()
}

type Once_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOnce_statementContext() *Once_statementContext {
	var p = new(Once_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_once_statement
	return p
}

func InitEmptyOnce_statementContext(p *Once_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_once_statement
}

func (*Once_statementContext) IsOnce_statementContext() {}

func NewOnce_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Once_statementContext {
	var p = new(Once_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YarnSpinnerParserRULE_once_statement

	return p
}

func (s *Once_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Once_statementContext) Once_primary_clause() IOnce_primary_clauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOnce_primary_clauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOnce_primary_clauseContext)
}

func (s *Once_statementContext) COMMAND_START() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_START, 0)
}

func (s *Once_statementContext) COMMAND_ENDONCE() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_ENDONCE, 0)
}

func (s *Once_statementContext) COMMAND_END() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_END, 0)
}

func (s *Once_statementContext) Once_alternate_clause() IOnce_alternate_clauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOnce_alternate_clauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOnce_alternate_clauseContext)
}

func (s *Once_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Once_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Once_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterOnce_statement(s)
	}
}

func (s *Once_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitOnce_statement(s)
	}
}

func (p *YarnSpinnerParser) Once_statement() (localctx IOnce_statementContext) {
	localctx = NewOnce_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, YarnSpinnerParserRULE_once_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(446)
		p.Once_primary_clause()
	}
	p.SetState(448)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 46, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(447)
			p.Once_alternate_clause()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(450)
		p.Match(YarnSpinnerParserCOMMAND_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(451)
		p.Match(YarnSpinnerParserCOMMAND_ENDONCE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(452)
		p.Match(YarnSpinnerParserCOMMAND_END)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOnce_primary_clauseContext is an interface to support dynamic dispatch.
type IOnce_primary_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COMMAND_START() antlr.TerminalNode
	COMMAND_ONCE() antlr.TerminalNode
	COMMAND_END() antlr.TerminalNode
	COMMAND_IF() antlr.TerminalNode
	Expression() IExpressionContext
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsOnce_primary_clauseContext differentiates from other interfaces.
	IsOnce_primary_clauseContext()
}

type Once_primary_clauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOnce_primary_clauseContext() *Once_primary_clauseContext {
	var p = new(Once_primary_clauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_once_primary_clause
	return p
}

func InitEmptyOnce_primary_clauseContext(p *Once_primary_clauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_once_primary_clause
}

func (*Once_primary_clauseContext) IsOnce_primary_clauseContext() {}

func NewOnce_primary_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Once_primary_clauseContext {
	var p = new(Once_primary_clauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YarnSpinnerParserRULE_once_primary_clause

	return p
}

func (s *Once_primary_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *Once_primary_clauseContext) COMMAND_START() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_START, 0)
}

func (s *Once_primary_clauseContext) COMMAND_ONCE() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_ONCE, 0)
}

func (s *Once_primary_clauseContext) COMMAND_END() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_END, 0)
}

func (s *Once_primary_clauseContext) COMMAND_IF() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_IF, 0)
}

func (s *Once_primary_clauseContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Once_primary_clauseContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *Once_primary_clauseContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *Once_primary_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Once_primary_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Once_primary_clauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterOnce_primary_clause(s)
	}
}

func (s *Once_primary_clauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitOnce_primary_clause(s)
	}
}

func (p *YarnSpinnerParser) Once_primary_clause() (localctx IOnce_primary_clauseContext) {
	localctx = NewOnce_primary_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, YarnSpinnerParserRULE_once_primary_clause)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(454)
		p.Match(YarnSpinnerParserCOMMAND_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(455)
		p.Match(YarnSpinnerParserCOMMAND_ONCE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(458)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == YarnSpinnerParserCOMMAND_IF {
		{
			p.SetState(456)
			p.Match(YarnSpinnerParserCOMMAND_IF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(457)
			p.expression(0)
		}

	}
	{
		p.SetState(460)
		p.Match(YarnSpinnerParserCOMMAND_END)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(464)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 48, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(461)
				p.Statement()
			}

		}
		p.SetState(466)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 48, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOnce_alternate_clauseContext is an interface to support dynamic dispatch.
type IOnce_alternate_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COMMAND_START() antlr.TerminalNode
	COMMAND_ELSE() antlr.TerminalNode
	COMMAND_END() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsOnce_alternate_clauseContext differentiates from other interfaces.
	IsOnce_alternate_clauseContext()
}

type Once_alternate_clauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOnce_alternate_clauseContext() *Once_alternate_clauseContext {
	var p = new(Once_alternate_clauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_once_alternate_clause
	return p
}

func InitEmptyOnce_alternate_clauseContext(p *Once_alternate_clauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_once_alternate_clause
}

func (*Once_alternate_clauseContext) IsOnce_alternate_clauseContext() {}

func NewOnce_alternate_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Once_alternate_clauseContext {
	var p = new(Once_alternate_clauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YarnSpinnerParserRULE_once_alternate_clause

	return p
}

func (s *Once_alternate_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *Once_alternate_clauseContext) COMMAND_START() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_START, 0)
}

func (s *Once_alternate_clauseContext) COMMAND_ELSE() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_ELSE, 0)
}

func (s *Once_alternate_clauseContext) COMMAND_END() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_END, 0)
}

func (s *Once_alternate_clauseContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *Once_alternate_clauseContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *Once_alternate_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Once_alternate_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Once_alternate_clauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterOnce_alternate_clause(s)
	}
}

func (s *Once_alternate_clauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitOnce_alternate_clause(s)
	}
}

func (p *YarnSpinnerParser) Once_alternate_clause() (localctx IOnce_alternate_clauseContext) {
	localctx = NewOnce_alternate_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, YarnSpinnerParserRULE_once_alternate_clause)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(467)
		p.Match(YarnSpinnerParserCOMMAND_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(468)
		p.Match(YarnSpinnerParserCOMMAND_ELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(469)
		p.Match(YarnSpinnerParserCOMMAND_END)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(473)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 49, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(470)
				p.Statement()
			}

		}
		p.SetState(475)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 49, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructured_commandContext is an interface to support dynamic dispatch.
type IStructured_commandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCommand_id returns the command_id token.
	GetCommand_id() antlr.Token

	// SetCommand_id sets the command_id token.
	SetCommand_id(antlr.Token)

	// Getter signatures
	FUNC_ID() antlr.TerminalNode
	AllStructured_command_value() []IStructured_command_valueContext
	Structured_command_value(i int) IStructured_command_valueContext

	// IsStructured_commandContext differentiates from other interfaces.
	IsStructured_commandContext()
}

type Structured_commandContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	command_id antlr.Token
}

func NewEmptyStructured_commandContext() *Structured_commandContext {
	var p = new(Structured_commandContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_structured_command
	return p
}

func InitEmptyStructured_commandContext(p *Structured_commandContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_structured_command
}

func (*Structured_commandContext) IsStructured_commandContext() {}

func NewStructured_commandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Structured_commandContext {
	var p = new(Structured_commandContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YarnSpinnerParserRULE_structured_command

	return p
}

func (s *Structured_commandContext) GetParser() antlr.Parser { return s.parser }

func (s *Structured_commandContext) GetCommand_id() antlr.Token { return s.command_id }

func (s *Structured_commandContext) SetCommand_id(v antlr.Token) { s.command_id = v }

func (s *Structured_commandContext) FUNC_ID() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserFUNC_ID, 0)
}

func (s *Structured_commandContext) AllStructured_command_value() []IStructured_command_valueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStructured_command_valueContext); ok {
			len++
		}
	}

	tst := make([]IStructured_command_valueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStructured_command_valueContext); ok {
			tst[i] = t.(IStructured_command_valueContext)
			i++
		}
	}

	return tst
}

func (s *Structured_commandContext) Structured_command_value(i int) IStructured_command_valueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructured_command_valueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructured_command_valueContext)
}

func (s *Structured_commandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Structured_commandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Structured_commandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterStructured_command(s)
	}
}

func (s *Structured_commandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitStructured_command(s)
	}
}

func (p *YarnSpinnerParser) Structured_command() (localctx IStructured_commandContext) {
	localctx = NewStructured_commandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, YarnSpinnerParserRULE_structured_command)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(476)

		var _m = p.Match(YarnSpinnerParserFUNC_ID)

		localctx.(*Structured_commandContext).command_id = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(480)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-33)) & ^0x3f) == 0 && ((int64(1)<<(_la-33))&15855525891) != 0 {
		{
			p.SetState(477)
			p.Structured_command_value()
		}

		p.SetState(482)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructured_command_valueContext is an interface to support dynamic dispatch.
type IStructured_command_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	FUNC_ID() antlr.TerminalNode

	// IsStructured_command_valueContext differentiates from other interfaces.
	IsStructured_command_valueContext()
}

type Structured_command_valueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructured_command_valueContext() *Structured_command_valueContext {
	var p = new(Structured_command_valueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_structured_command_value
	return p
}

func InitEmptyStructured_command_valueContext(p *Structured_command_valueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = YarnSpinnerParserRULE_structured_command_value
}

func (*Structured_command_valueContext) IsStructured_command_valueContext() {}

func NewStructured_command_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Structured_command_valueContext {
	var p = new(Structured_command_valueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = YarnSpinnerParserRULE_structured_command_value

	return p
}

func (s *Structured_command_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Structured_command_valueContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Structured_command_valueContext) FUNC_ID() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserFUNC_ID, 0)
}

func (s *Structured_command_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Structured_command_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Structured_command_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterStructured_command_value(s)
	}
}

func (s *Structured_command_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitStructured_command_value(s)
	}
}

func (p *YarnSpinnerParser) Structured_command_value() (localctx IStructured_command_valueContext) {
	localctx = NewStructured_command_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, YarnSpinnerParserRULE_structured_command_value)
	p.SetState(485)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 51, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(483)
			p.expression(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(484)
			p.Match(YarnSpinnerParserFUNC_ID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *YarnSpinnerParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 13:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *YarnSpinnerParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
