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
		"", "", "", "", "", "", "", "", "'---'", "", "'#'", "", "", "'==='",
		"'->'", "'<<'", "", "", "'\\'", "", "", "", "", "", "", "", "", "",
		"'true'", "'false'", "'null'", "", "", "", "", "", "", "", "", "", "",
		"", "'+='", "'-='", "'*='", "'%='", "'/='", "'+'", "'-'", "'*'", "'/'",
		"'%'", "'('", "')'", "','", "'as'", "", "", "'}'", "", "'.'", "", "",
		"", "", "", "", "", "'endif'", "", "", "", "", "", "", "", "", "", "",
		"'{'", "", "", "'string'", "'number'", "'bool'",
	}
	staticData.SymbolicNames = []string{
		"", "INDENT", "DEDENT", "BLANK_LINE_FOLLOWING_OPTION", "WS", "COMMENT",
		"NEWLINE", "ID", "BODY_START", "HEADER_DELIMITER", "HASHTAG", "REST_OF_LINE",
		"BODY_WS", "BODY_END", "SHORTCUT_ARROW", "COMMAND_START", "EXPRESSION_START",
		"ESCAPED_ANY", "TEXT_ESCAPE", "TEXT_COMMENT", "TEXT", "UNESCAPABLE_CHARACTER",
		"TEXT_COMMANDHASHTAG_WS", "TEXT_COMMANDHASHTAG_COMMENT", "TEXT_COMMANDHASHTAG_ERROR",
		"HASHTAG_WS", "HASHTAG_TEXT", "EXPR_WS", "KEYWORD_TRUE", "KEYWORD_FALSE",
		"KEYWORD_NULL", "OPERATOR_ASSIGNMENT", "OPERATOR_LOGICAL_LESS_THAN_EQUALS",
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
		"COMMAND_DECLARE", "COMMAND_JUMP", "COMMAND_ENUM", "COMMAND_CASE", "COMMAND_ENDENUM",
		"COMMAND_LOCAL", "COMMAND_END", "COMMAND_TEXT_NEWLINE", "COMMAND_TEXT_END",
		"COMMAND_EXPRESSION_START", "COMMAND_TEXT", "COMMAND_ID_NEWLINE", "TYPE_STRING",
		"TYPE_NUMBER", "TYPE_BOOL",
	}
	staticData.RuleNames = []string{
		"dialogue", "file_hashtag", "node", "header", "body", "statement", "line_statement",
		"line_formatted_text", "hashtag", "line_condition", "expression", "value",
		"variable", "function_call", "if_statement", "if_clause", "else_if_clause",
		"else_clause", "set_statement", "call_statement", "command_statement",
		"command_formatted_text", "shortcut_option_statement", "shortcut_option",
		"declare_statement", "jump_statement",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 84, 315, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 1, 0,
		5, 0, 54, 8, 0, 10, 0, 12, 0, 57, 9, 0, 1, 0, 4, 0, 60, 8, 0, 11, 0, 12,
		0, 61, 1, 1, 1, 1, 1, 1, 1, 2, 4, 2, 68, 8, 2, 11, 2, 12, 2, 69, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 3, 3, 79, 8, 3, 1, 4, 5, 4, 82, 8,
		4, 10, 4, 12, 4, 85, 9, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 5, 5, 97, 8, 5, 10, 5, 12, 5, 100, 9, 5, 1, 5, 3, 5, 103,
		8, 5, 1, 6, 1, 6, 3, 6, 107, 8, 6, 1, 6, 5, 6, 110, 8, 6, 10, 6, 12, 6,
		113, 9, 6, 1, 6, 1, 6, 1, 7, 4, 7, 118, 8, 7, 11, 7, 12, 7, 119, 1, 7,
		1, 7, 1, 7, 1, 7, 4, 7, 126, 8, 7, 11, 7, 12, 7, 127, 1, 8, 1, 8, 1, 8,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 148, 8, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 5, 10, 165, 8, 10, 10, 10, 12, 10, 168, 9, 10, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 177, 8, 11, 1, 12, 1, 12, 1, 13,
		1, 13, 1, 13, 3, 13, 184, 8, 13, 1, 13, 1, 13, 5, 13, 188, 8, 13, 10, 13,
		12, 13, 191, 9, 13, 1, 13, 1, 13, 1, 14, 1, 14, 5, 14, 197, 8, 14, 10,
		14, 12, 14, 200, 9, 14, 1, 14, 3, 14, 203, 8, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 5, 15, 214, 8, 15, 10, 15, 12,
		15, 217, 9, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 5, 16, 224, 8, 16, 10,
		16, 12, 16, 227, 9, 16, 1, 17, 1, 17, 1, 17, 1, 17, 5, 17, 233, 8, 17,
		10, 17, 12, 17, 236, 9, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 5, 20,
		254, 8, 20, 10, 20, 12, 20, 257, 9, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 5, 21, 264, 8, 21, 10, 21, 12, 21, 267, 9, 21, 1, 22, 5, 22, 270, 8,
		22, 10, 22, 12, 22, 273, 9, 22, 1, 22, 1, 22, 3, 22, 277, 8, 22, 1, 23,
		1, 23, 1, 23, 1, 23, 5, 23, 283, 8, 23, 10, 23, 12, 23, 286, 9, 23, 1,
		23, 3, 23, 289, 8, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24,
		3, 24, 298, 8, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 3, 25, 313, 8, 25, 1, 25, 0, 1,
		20, 26, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32,
		34, 36, 38, 40, 42, 44, 46, 48, 50, 0, 6, 1, 0, 49, 51, 1, 0, 47, 48, 2,
		0, 32, 33, 35, 36, 2, 0, 34, 34, 37, 37, 1, 0, 38, 40, 2, 0, 31, 31, 42,
		46, 337, 0, 55, 1, 0, 0, 0, 2, 63, 1, 0, 0, 0, 4, 67, 1, 0, 0, 0, 6, 75,
		1, 0, 0, 0, 8, 83, 1, 0, 0, 0, 10, 102, 1, 0, 0, 0, 12, 104, 1, 0, 0, 0,
		14, 125, 1, 0, 0, 0, 16, 129, 1, 0, 0, 0, 18, 132, 1, 0, 0, 0, 20, 147,
		1, 0, 0, 0, 22, 176, 1, 0, 0, 0, 24, 178, 1, 0, 0, 0, 26, 180, 1, 0, 0,
		0, 28, 194, 1, 0, 0, 0, 30, 208, 1, 0, 0, 0, 32, 218, 1, 0, 0, 0, 34, 228,
		1, 0, 0, 0, 36, 237, 1, 0, 0, 0, 38, 244, 1, 0, 0, 0, 40, 249, 1, 0, 0,
		0, 42, 265, 1, 0, 0, 0, 44, 271, 1, 0, 0, 0, 46, 278, 1, 0, 0, 0, 48, 290,
		1, 0, 0, 0, 50, 312, 1, 0, 0, 0, 52, 54, 3, 2, 1, 0, 53, 52, 1, 0, 0, 0,
		54, 57, 1, 0, 0, 0, 55, 53, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0, 56, 59, 1,
		0, 0, 0, 57, 55, 1, 0, 0, 0, 58, 60, 3, 4, 2, 0, 59, 58, 1, 0, 0, 0, 60,
		61, 1, 0, 0, 0, 61, 59, 1, 0, 0, 0, 61, 62, 1, 0, 0, 0, 62, 1, 1, 0, 0,
		0, 63, 64, 5, 10, 0, 0, 64, 65, 5, 26, 0, 0, 65, 3, 1, 0, 0, 0, 66, 68,
		3, 6, 3, 0, 67, 66, 1, 0, 0, 0, 68, 69, 1, 0, 0, 0, 69, 67, 1, 0, 0, 0,
		69, 70, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 72, 5, 8, 0, 0, 72, 73, 3,
		8, 4, 0, 73, 74, 5, 13, 0, 0, 74, 5, 1, 0, 0, 0, 75, 76, 5, 7, 0, 0, 76,
		78, 5, 9, 0, 0, 77, 79, 5, 11, 0, 0, 78, 77, 1, 0, 0, 0, 78, 79, 1, 0,
		0, 0, 79, 7, 1, 0, 0, 0, 80, 82, 3, 10, 5, 0, 81, 80, 1, 0, 0, 0, 82, 85,
		1, 0, 0, 0, 83, 81, 1, 0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 9, 1, 0, 0, 0,
		85, 83, 1, 0, 0, 0, 86, 103, 3, 12, 6, 0, 87, 103, 3, 28, 14, 0, 88, 103,
		3, 36, 18, 0, 89, 103, 3, 44, 22, 0, 90, 103, 3, 38, 19, 0, 91, 103, 3,
		40, 20, 0, 92, 103, 3, 48, 24, 0, 93, 103, 3, 50, 25, 0, 94, 98, 5, 1,
		0, 0, 95, 97, 3, 10, 5, 0, 96, 95, 1, 0, 0, 0, 97, 100, 1, 0, 0, 0, 98,
		96, 1, 0, 0, 0, 98, 99, 1, 0, 0, 0, 99, 101, 1, 0, 0, 0, 100, 98, 1, 0,
		0, 0, 101, 103, 5, 2, 0, 0, 102, 86, 1, 0, 0, 0, 102, 87, 1, 0, 0, 0, 102,
		88, 1, 0, 0, 0, 102, 89, 1, 0, 0, 0, 102, 90, 1, 0, 0, 0, 102, 91, 1, 0,
		0, 0, 102, 92, 1, 0, 0, 0, 102, 93, 1, 0, 0, 0, 102, 94, 1, 0, 0, 0, 103,
		11, 1, 0, 0, 0, 104, 106, 3, 14, 7, 0, 105, 107, 3, 18, 9, 0, 106, 105,
		1, 0, 0, 0, 106, 107, 1, 0, 0, 0, 107, 111, 1, 0, 0, 0, 108, 110, 3, 16,
		8, 0, 109, 108, 1, 0, 0, 0, 110, 113, 1, 0, 0, 0, 111, 109, 1, 0, 0, 0,
		111, 112, 1, 0, 0, 0, 112, 114, 1, 0, 0, 0, 113, 111, 1, 0, 0, 0, 114,
		115, 5, 6, 0, 0, 115, 13, 1, 0, 0, 0, 116, 118, 5, 20, 0, 0, 117, 116,
		1, 0, 0, 0, 118, 119, 1, 0, 0, 0, 119, 117, 1, 0, 0, 0, 119, 120, 1, 0,
		0, 0, 120, 126, 1, 0, 0, 0, 121, 122, 5, 16, 0, 0, 122, 123, 3, 20, 10,
		0, 123, 124, 5, 58, 0, 0, 124, 126, 1, 0, 0, 0, 125, 117, 1, 0, 0, 0, 125,
		121, 1, 0, 0, 0, 126, 127, 1, 0, 0, 0, 127, 125, 1, 0, 0, 0, 127, 128,
		1, 0, 0, 0, 128, 15, 1, 0, 0, 0, 129, 130, 5, 10, 0, 0, 130, 131, 5, 26,
		0, 0, 131, 17, 1, 0, 0, 0, 132, 133, 5, 15, 0, 0, 133, 134, 5, 64, 0, 0,
		134, 135, 3, 20, 10, 0, 135, 136, 5, 76, 0, 0, 136, 19, 1, 0, 0, 0, 137,
		138, 6, 10, -1, 0, 138, 139, 5, 52, 0, 0, 139, 140, 3, 20, 10, 0, 140,
		141, 5, 53, 0, 0, 141, 148, 1, 0, 0, 0, 142, 143, 5, 48, 0, 0, 143, 148,
		3, 20, 10, 8, 144, 145, 5, 41, 0, 0, 145, 148, 3, 20, 10, 7, 146, 148,
		3, 22, 11, 0, 147, 137, 1, 0, 0, 0, 147, 142, 1, 0, 0, 0, 147, 144, 1,
		0, 0, 0, 147, 146, 1, 0, 0, 0, 148, 166, 1, 0, 0, 0, 149, 150, 10, 6, 0,
		0, 150, 151, 7, 0, 0, 0, 151, 165, 3, 20, 10, 7, 152, 153, 10, 5, 0, 0,
		153, 154, 7, 1, 0, 0, 154, 165, 3, 20, 10, 6, 155, 156, 10, 4, 0, 0, 156,
		157, 7, 2, 0, 0, 157, 165, 3, 20, 10, 5, 158, 159, 10, 3, 0, 0, 159, 160,
		7, 3, 0, 0, 160, 165, 3, 20, 10, 4, 161, 162, 10, 2, 0, 0, 162, 163, 7,
		4, 0, 0, 163, 165, 3, 20, 10, 3, 164, 149, 1, 0, 0, 0, 164, 152, 1, 0,
		0, 0, 164, 155, 1, 0, 0, 0, 164, 158, 1, 0, 0, 0, 164, 161, 1, 0, 0, 0,
		165, 168, 1, 0, 0, 0, 166, 164, 1, 0, 0, 0, 166, 167, 1, 0, 0, 0, 167,
		21, 1, 0, 0, 0, 168, 166, 1, 0, 0, 0, 169, 177, 5, 61, 0, 0, 170, 177,
		5, 28, 0, 0, 171, 177, 5, 29, 0, 0, 172, 177, 3, 24, 12, 0, 173, 177, 5,
		56, 0, 0, 174, 177, 5, 30, 0, 0, 175, 177, 3, 26, 13, 0, 176, 169, 1, 0,
		0, 0, 176, 170, 1, 0, 0, 0, 176, 171, 1, 0, 0, 0, 176, 172, 1, 0, 0, 0,
		176, 173, 1, 0, 0, 0, 176, 174, 1, 0, 0, 0, 176, 175, 1, 0, 0, 0, 177,
		23, 1, 0, 0, 0, 178, 179, 5, 59, 0, 0, 179, 25, 1, 0, 0, 0, 180, 181, 5,
		57, 0, 0, 181, 183, 5, 52, 0, 0, 182, 184, 3, 20, 10, 0, 183, 182, 1, 0,
		0, 0, 183, 184, 1, 0, 0, 0, 184, 189, 1, 0, 0, 0, 185, 186, 5, 54, 0, 0,
		186, 188, 3, 20, 10, 0, 187, 185, 1, 0, 0, 0, 188, 191, 1, 0, 0, 0, 189,
		187, 1, 0, 0, 0, 189, 190, 1, 0, 0, 0, 190, 192, 1, 0, 0, 0, 191, 189,
		1, 0, 0, 0, 192, 193, 5, 53, 0, 0, 193, 27, 1, 0, 0, 0, 194, 198, 3, 30,
		15, 0, 195, 197, 3, 32, 16, 0, 196, 195, 1, 0, 0, 0, 197, 200, 1, 0, 0,
		0, 198, 196, 1, 0, 0, 0, 198, 199, 1, 0, 0, 0, 199, 202, 1, 0, 0, 0, 200,
		198, 1, 0, 0, 0, 201, 203, 3, 34, 17, 0, 202, 201, 1, 0, 0, 0, 202, 203,
		1, 0, 0, 0, 203, 204, 1, 0, 0, 0, 204, 205, 5, 15, 0, 0, 205, 206, 5, 68,
		0, 0, 206, 207, 5, 76, 0, 0, 207, 29, 1, 0, 0, 0, 208, 209, 5, 15, 0, 0,
		209, 210, 5, 64, 0, 0, 210, 211, 3, 20, 10, 0, 211, 215, 5, 76, 0, 0, 212,
		214, 3, 10, 5, 0, 213, 212, 1, 0, 0, 0, 214, 217, 1, 0, 0, 0, 215, 213,
		1, 0, 0, 0, 215, 216, 1, 0, 0, 0, 216, 31, 1, 0, 0, 0, 217, 215, 1, 0,
		0, 0, 218, 219, 5, 15, 0, 0, 219, 220, 5, 65, 0, 0, 220, 221, 3, 20, 10,
		0, 221, 225, 5, 76, 0, 0, 222, 224, 3, 10, 5, 0, 223, 222, 1, 0, 0, 0,
		224, 227, 1, 0, 0, 0, 225, 223, 1, 0, 0, 0, 225, 226, 1, 0, 0, 0, 226,
		33, 1, 0, 0, 0, 227, 225, 1, 0, 0, 0, 228, 229, 5, 15, 0, 0, 229, 230,
		5, 66, 0, 0, 230, 234, 5, 76, 0, 0, 231, 233, 3, 10, 5, 0, 232, 231, 1,
		0, 0, 0, 233, 236, 1, 0, 0, 0, 234, 232, 1, 0, 0, 0, 234, 235, 1, 0, 0,
		0, 235, 35, 1, 0, 0, 0, 236, 234, 1, 0, 0, 0, 237, 238, 5, 15, 0, 0, 238,
		239, 5, 67, 0, 0, 239, 240, 3, 24, 12, 0, 240, 241, 7, 5, 0, 0, 241, 242,
		3, 20, 10, 0, 242, 243, 5, 76, 0, 0, 243, 37, 1, 0, 0, 0, 244, 245, 5,
		15, 0, 0, 245, 246, 5, 69, 0, 0, 246, 247, 3, 26, 13, 0, 247, 248, 5, 76,
		0, 0, 248, 39, 1, 0, 0, 0, 249, 250, 5, 15, 0, 0, 250, 251, 3, 42, 21,
		0, 251, 255, 5, 78, 0, 0, 252, 254, 3, 16, 8, 0, 253, 252, 1, 0, 0, 0,
		254, 257, 1, 0, 0, 0, 255, 253, 1, 0, 0, 0, 255, 256, 1, 0, 0, 0, 256,
		41, 1, 0, 0, 0, 257, 255, 1, 0, 0, 0, 258, 264, 5, 80, 0, 0, 259, 260,
		5, 79, 0, 0, 260, 261, 3, 20, 10, 0, 261, 262, 5, 58, 0, 0, 262, 264, 1,
		0, 0, 0, 263, 258, 1, 0, 0, 0, 263, 259, 1, 0, 0, 0, 264, 267, 1, 0, 0,
		0, 265, 263, 1, 0, 0, 0, 265, 266, 1, 0, 0, 0, 266, 43, 1, 0, 0, 0, 267,
		265, 1, 0, 0, 0, 268, 270, 3, 46, 23, 0, 269, 268, 1, 0, 0, 0, 270, 273,
		1, 0, 0, 0, 271, 269, 1, 0, 0, 0, 271, 272, 1, 0, 0, 0, 272, 274, 1, 0,
		0, 0, 273, 271, 1, 0, 0, 0, 274, 276, 3, 46, 23, 0, 275, 277, 5, 3, 0,
		0, 276, 275, 1, 0, 0, 0, 276, 277, 1, 0, 0, 0, 277, 45, 1, 0, 0, 0, 278,
		279, 5, 14, 0, 0, 279, 288, 3, 12, 6, 0, 280, 284, 5, 1, 0, 0, 281, 283,
		3, 10, 5, 0, 282, 281, 1, 0, 0, 0, 283, 286, 1, 0, 0, 0, 284, 282, 1, 0,
		0, 0, 284, 285, 1, 0, 0, 0, 285, 287, 1, 0, 0, 0, 286, 284, 1, 0, 0, 0,
		287, 289, 5, 2, 0, 0, 288, 280, 1, 0, 0, 0, 288, 289, 1, 0, 0, 0, 289,
		47, 1, 0, 0, 0, 290, 291, 5, 15, 0, 0, 291, 292, 5, 70, 0, 0, 292, 293,
		3, 24, 12, 0, 293, 294, 5, 31, 0, 0, 294, 297, 3, 22, 11, 0, 295, 296,
		5, 55, 0, 0, 296, 298, 5, 57, 0, 0, 297, 295, 1, 0, 0, 0, 297, 298, 1,
		0, 0, 0, 298, 299, 1, 0, 0, 0, 299, 300, 5, 76, 0, 0, 300, 49, 1, 0, 0,
		0, 301, 302, 5, 15, 0, 0, 302, 303, 5, 71, 0, 0, 303, 304, 5, 7, 0, 0,
		304, 313, 5, 76, 0, 0, 305, 306, 5, 15, 0, 0, 306, 307, 5, 71, 0, 0, 307,
		308, 5, 16, 0, 0, 308, 309, 3, 20, 10, 0, 309, 310, 5, 58, 0, 0, 310, 311,
		5, 76, 0, 0, 311, 313, 1, 0, 0, 0, 312, 301, 1, 0, 0, 0, 312, 305, 1, 0,
		0, 0, 313, 51, 1, 0, 0, 0, 32, 55, 61, 69, 78, 83, 98, 102, 106, 111, 119,
		125, 127, 147, 164, 166, 176, 183, 189, 198, 202, 215, 225, 234, 255, 263,
		265, 271, 276, 284, 288, 297, 312,
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
	YarnSpinnerParserID                                   = 7
	YarnSpinnerParserBODY_START                           = 8
	YarnSpinnerParserHEADER_DELIMITER                     = 9
	YarnSpinnerParserHASHTAG                              = 10
	YarnSpinnerParserREST_OF_LINE                         = 11
	YarnSpinnerParserBODY_WS                              = 12
	YarnSpinnerParserBODY_END                             = 13
	YarnSpinnerParserSHORTCUT_ARROW                       = 14
	YarnSpinnerParserCOMMAND_START                        = 15
	YarnSpinnerParserEXPRESSION_START                     = 16
	YarnSpinnerParserESCAPED_ANY                          = 17
	YarnSpinnerParserTEXT_ESCAPE                          = 18
	YarnSpinnerParserTEXT_COMMENT                         = 19
	YarnSpinnerParserTEXT                                 = 20
	YarnSpinnerParserUNESCAPABLE_CHARACTER                = 21
	YarnSpinnerParserTEXT_COMMANDHASHTAG_WS               = 22
	YarnSpinnerParserTEXT_COMMANDHASHTAG_COMMENT          = 23
	YarnSpinnerParserTEXT_COMMANDHASHTAG_ERROR            = 24
	YarnSpinnerParserHASHTAG_WS                           = 25
	YarnSpinnerParserHASHTAG_TEXT                         = 26
	YarnSpinnerParserEXPR_WS                              = 27
	YarnSpinnerParserKEYWORD_TRUE                         = 28
	YarnSpinnerParserKEYWORD_FALSE                        = 29
	YarnSpinnerParserKEYWORD_NULL                         = 30
	YarnSpinnerParserOPERATOR_ASSIGNMENT                  = 31
	YarnSpinnerParserOPERATOR_LOGICAL_LESS_THAN_EQUALS    = 32
	YarnSpinnerParserOPERATOR_LOGICAL_GREATER_THAN_EQUALS = 33
	YarnSpinnerParserOPERATOR_LOGICAL_EQUALS              = 34
	YarnSpinnerParserOPERATOR_LOGICAL_LESS                = 35
	YarnSpinnerParserOPERATOR_LOGICAL_GREATER             = 36
	YarnSpinnerParserOPERATOR_LOGICAL_NOT_EQUALS          = 37
	YarnSpinnerParserOPERATOR_LOGICAL_AND                 = 38
	YarnSpinnerParserOPERATOR_LOGICAL_OR                  = 39
	YarnSpinnerParserOPERATOR_LOGICAL_XOR                 = 40
	YarnSpinnerParserOPERATOR_LOGICAL_NOT                 = 41
	YarnSpinnerParserOPERATOR_MATHS_ADDITION_EQUALS       = 42
	YarnSpinnerParserOPERATOR_MATHS_SUBTRACTION_EQUALS    = 43
	YarnSpinnerParserOPERATOR_MATHS_MULTIPLICATION_EQUALS = 44
	YarnSpinnerParserOPERATOR_MATHS_MODULUS_EQUALS        = 45
	YarnSpinnerParserOPERATOR_MATHS_DIVISION_EQUALS       = 46
	YarnSpinnerParserOPERATOR_MATHS_ADDITION              = 47
	YarnSpinnerParserOPERATOR_MATHS_SUBTRACTION           = 48
	YarnSpinnerParserOPERATOR_MATHS_MULTIPLICATION        = 49
	YarnSpinnerParserOPERATOR_MATHS_DIVISION              = 50
	YarnSpinnerParserOPERATOR_MATHS_MODULUS               = 51
	YarnSpinnerParserLPAREN                               = 52
	YarnSpinnerParserRPAREN                               = 53
	YarnSpinnerParserCOMMA                                = 54
	YarnSpinnerParserEXPRESSION_AS                        = 55
	YarnSpinnerParserSTRING                               = 56
	YarnSpinnerParserFUNC_ID                              = 57
	YarnSpinnerParserEXPRESSION_END                       = 58
	YarnSpinnerParserVAR_ID                               = 59
	YarnSpinnerParserDOT                                  = 60
	YarnSpinnerParserNUMBER                               = 61
	YarnSpinnerParserCOMMAND_NEWLINE                      = 62
	YarnSpinnerParserCOMMAND_WS                           = 63
	YarnSpinnerParserCOMMAND_IF                           = 64
	YarnSpinnerParserCOMMAND_ELSEIF                       = 65
	YarnSpinnerParserCOMMAND_ELSE                         = 66
	YarnSpinnerParserCOMMAND_SET                          = 67
	YarnSpinnerParserCOMMAND_ENDIF                        = 68
	YarnSpinnerParserCOMMAND_CALL                         = 69
	YarnSpinnerParserCOMMAND_DECLARE                      = 70
	YarnSpinnerParserCOMMAND_JUMP                         = 71
	YarnSpinnerParserCOMMAND_ENUM                         = 72
	YarnSpinnerParserCOMMAND_CASE                         = 73
	YarnSpinnerParserCOMMAND_ENDENUM                      = 74
	YarnSpinnerParserCOMMAND_LOCAL                        = 75
	YarnSpinnerParserCOMMAND_END                          = 76
	YarnSpinnerParserCOMMAND_TEXT_NEWLINE                 = 77
	YarnSpinnerParserCOMMAND_TEXT_END                     = 78
	YarnSpinnerParserCOMMAND_EXPRESSION_START             = 79
	YarnSpinnerParserCOMMAND_TEXT                         = 80
	YarnSpinnerParserCOMMAND_ID_NEWLINE                   = 81
	YarnSpinnerParserTYPE_STRING                          = 82
	YarnSpinnerParserTYPE_NUMBER                          = 83
	YarnSpinnerParserTYPE_BOOL                            = 84
)

// YarnSpinnerParser rules.
const (
	YarnSpinnerParserRULE_dialogue                  = 0
	YarnSpinnerParserRULE_file_hashtag              = 1
	YarnSpinnerParserRULE_node                      = 2
	YarnSpinnerParserRULE_header                    = 3
	YarnSpinnerParserRULE_body                      = 4
	YarnSpinnerParserRULE_statement                 = 5
	YarnSpinnerParserRULE_line_statement            = 6
	YarnSpinnerParserRULE_line_formatted_text       = 7
	YarnSpinnerParserRULE_hashtag                   = 8
	YarnSpinnerParserRULE_line_condition            = 9
	YarnSpinnerParserRULE_expression                = 10
	YarnSpinnerParserRULE_value                     = 11
	YarnSpinnerParserRULE_variable                  = 12
	YarnSpinnerParserRULE_function_call             = 13
	YarnSpinnerParserRULE_if_statement              = 14
	YarnSpinnerParserRULE_if_clause                 = 15
	YarnSpinnerParserRULE_else_if_clause            = 16
	YarnSpinnerParserRULE_else_clause               = 17
	YarnSpinnerParserRULE_set_statement             = 18
	YarnSpinnerParserRULE_call_statement            = 19
	YarnSpinnerParserRULE_command_statement         = 20
	YarnSpinnerParserRULE_command_formatted_text    = 21
	YarnSpinnerParserRULE_shortcut_option_statement = 22
	YarnSpinnerParserRULE_shortcut_option           = 23
	YarnSpinnerParserRULE_declare_statement         = 24
	YarnSpinnerParserRULE_jump_statement            = 25
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
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == YarnSpinnerParserHASHTAG {
		{
			p.SetState(52)
			p.File_hashtag()
		}

		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == YarnSpinnerParserID {
		{
			p.SetState(58)
			p.Node()
		}

		p.SetState(61)
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
		p.SetState(63)
		p.Match(YarnSpinnerParserHASHTAG)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(64)

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
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == YarnSpinnerParserID {
		{
			p.SetState(66)
			p.Header()
		}

		p.SetState(69)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(71)
		p.Match(YarnSpinnerParserBODY_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(72)
		p.Body()
	}
	{
		p.SetState(73)
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
	p.EnterRule(localctx, 6, YarnSpinnerParserRULE_header)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(75)

		var _m = p.Match(YarnSpinnerParserID)

		localctx.(*HeaderContext).header_key = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(76)
		p.Match(YarnSpinnerParserHEADER_DELIMITER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == YarnSpinnerParserREST_OF_LINE {
		{
			p.SetState(77)

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
	p.EnterRule(localctx, 8, YarnSpinnerParserRULE_body)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1163266) != 0 {
		{
			p.SetState(80)
			p.Statement()
		}

		p.SetState(85)
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
	Jump_statement() IJump_statementContext
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
	p.EnterRule(localctx, 10, YarnSpinnerParserRULE_statement)
	var _la int

	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(86)
			p.Line_statement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(87)
			p.If_statement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(88)
			p.Set_statement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(89)
			p.Shortcut_option_statement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(90)
			p.Call_statement()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(91)
			p.Command_statement()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(92)
			p.Declare_statement()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(93)
			p.Jump_statement()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(94)
			p.Match(YarnSpinnerParserINDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1163266) != 0 {
			{
				p.SetState(95)
				p.Statement()
			}

			p.SetState(100)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(101)
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
	p.EnterRule(localctx, 12, YarnSpinnerParserRULE_line_statement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(104)
		p.Line_formatted_text()
	}
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == YarnSpinnerParserCOMMAND_START {
		{
			p.SetState(105)
			p.Line_condition()
		}

	}
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == YarnSpinnerParserHASHTAG {
		{
			p.SetState(108)
			p.Hashtag()
		}

		p.SetState(113)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(114)
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
	p.EnterRule(localctx, 14, YarnSpinnerParserRULE_line_formatted_text)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == YarnSpinnerParserEXPRESSION_START || _la == YarnSpinnerParserTEXT {
		p.SetState(125)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case YarnSpinnerParserTEXT:
			p.SetState(117)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = 1
			for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				switch _alt {
				case 1:
					{
						p.SetState(116)
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

				p.SetState(119)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
				if p.HasError() {
					goto errorExit
				}
			}

		case YarnSpinnerParserEXPRESSION_START:
			{
				p.SetState(121)
				p.Match(YarnSpinnerParserEXPRESSION_START)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(122)
				p.expression(0)
			}
			{
				p.SetState(123)
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

		p.SetState(127)
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
	p.EnterRule(localctx, 16, YarnSpinnerParserRULE_hashtag)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(129)
		p.Match(YarnSpinnerParserHASHTAG)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(130)

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

	// Getter signatures
	COMMAND_START() antlr.TerminalNode
	COMMAND_IF() antlr.TerminalNode
	Expression() IExpressionContext
	COMMAND_END() antlr.TerminalNode

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

func (s *Line_conditionContext) COMMAND_START() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_START, 0)
}

func (s *Line_conditionContext) COMMAND_IF() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_IF, 0)
}

func (s *Line_conditionContext) Expression() IExpressionContext {
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

func (s *Line_conditionContext) COMMAND_END() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserCOMMAND_END, 0)
}

func (s *Line_conditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Line_conditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Line_conditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterLine_condition(s)
	}
}

func (s *Line_conditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitLine_condition(s)
	}
}

func (p *YarnSpinnerParser) Line_condition() (localctx ILine_conditionContext) {
	localctx = NewLine_conditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, YarnSpinnerParserRULE_line_condition)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(132)
		p.Match(YarnSpinnerParserCOMMAND_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(133)
		p.Match(YarnSpinnerParserCOMMAND_IF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(134)
		p.expression(0)
	}
	{
		p.SetState(135)
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
	_startState := 20
	p.EnterRecursionRule(localctx, 20, YarnSpinnerParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(147)
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
			p.SetState(138)
			p.Match(YarnSpinnerParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(139)
			p.expression(0)
		}
		{
			p.SetState(140)
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
			p.SetState(142)

			var _m = p.Match(YarnSpinnerParserOPERATOR_MATHS_SUBTRACTION)

			localctx.(*ExpNegativeContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(143)
			p.expression(8)
		}

	case YarnSpinnerParserOPERATOR_LOGICAL_NOT:
		localctx = NewExpNotContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(144)

			var _m = p.Match(YarnSpinnerParserOPERATOR_LOGICAL_NOT)

			localctx.(*ExpNotContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(145)
			p.expression(7)
		}

	case YarnSpinnerParserKEYWORD_TRUE, YarnSpinnerParserKEYWORD_FALSE, YarnSpinnerParserKEYWORD_NULL, YarnSpinnerParserSTRING, YarnSpinnerParserFUNC_ID, YarnSpinnerParserVAR_ID, YarnSpinnerParserNUMBER:
		localctx = NewExpValueContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(146)
			p.Value()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(166)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(164)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpMultDivModContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, YarnSpinnerParserRULE_expression)
				p.SetState(149)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(150)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpMultDivModContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3940649673949184) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpMultDivModContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(151)
					p.expression(7)
				}

			case 2:
				localctx = NewExpAddSubContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, YarnSpinnerParserRULE_expression)
				p.SetState(152)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(153)

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
					p.SetState(154)
					p.expression(6)
				}

			case 3:
				localctx = NewExpComparisonContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, YarnSpinnerParserRULE_expression)
				p.SetState(155)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(156)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpComparisonContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&115964116992) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpComparisonContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(157)
					p.expression(5)
				}

			case 4:
				localctx = NewExpEqualityContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, YarnSpinnerParserRULE_expression)
				p.SetState(158)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(159)

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
					p.SetState(160)
					p.expression(4)
				}

			case 5:
				localctx = NewExpAndOrXorContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, YarnSpinnerParserRULE_expression)
				p.SetState(161)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(162)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpAndOrXorContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1924145348608) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpAndOrXorContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(163)
					p.expression(3)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(168)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
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

type ValueNullContext struct {
	ValueContext
}

func NewValueNullContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValueNullContext {
	var p = new(ValueNullContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *ValueNullContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueNullContext) KEYWORD_NULL() antlr.TerminalNode {
	return s.GetToken(YarnSpinnerParserKEYWORD_NULL, 0)
}

func (s *ValueNullContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.EnterValueNull(s)
	}
}

func (s *ValueNullContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YarnSpinnerParserListener); ok {
		listenerT.ExitValueNull(s)
	}
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

func (p *YarnSpinnerParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, YarnSpinnerParserRULE_value)
	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case YarnSpinnerParserNUMBER:
		localctx = NewValueNumberContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(169)
			p.Match(YarnSpinnerParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case YarnSpinnerParserKEYWORD_TRUE:
		localctx = NewValueTrueContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(170)
			p.Match(YarnSpinnerParserKEYWORD_TRUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case YarnSpinnerParserKEYWORD_FALSE:
		localctx = NewValueFalseContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(171)
			p.Match(YarnSpinnerParserKEYWORD_FALSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case YarnSpinnerParserVAR_ID:
		localctx = NewValueVarContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(172)
			p.Variable()
		}

	case YarnSpinnerParserSTRING:
		localctx = NewValueStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(173)
			p.Match(YarnSpinnerParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case YarnSpinnerParserKEYWORD_NULL:
		localctx = NewValueNullContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(174)
			p.Match(YarnSpinnerParserKEYWORD_NULL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case YarnSpinnerParserFUNC_ID:
		localctx = NewValueFuncContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(175)
			p.Function_call()
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
	p.EnterRule(localctx, 24, YarnSpinnerParserRULE_variable)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(178)
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
	p.EnterRule(localctx, 26, YarnSpinnerParserRULE_function_call)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(180)
		p.Match(YarnSpinnerParserFUNC_ID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(181)
		p.Match(YarnSpinnerParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3103263819137286144) != 0 {
		{
			p.SetState(182)
			p.expression(0)
		}

	}
	p.SetState(189)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == YarnSpinnerParserCOMMA {
		{
			p.SetState(185)
			p.Match(YarnSpinnerParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(186)
			p.expression(0)
		}

		p.SetState(191)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(192)
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
	p.EnterRule(localctx, 28, YarnSpinnerParserRULE_if_statement)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(194)
		p.If_clause()
	}
	p.SetState(198)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(195)
				p.Else_if_clause()
			}

		}
		p.SetState(200)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(202)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(201)
			p.Else_clause()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(204)
		p.Match(YarnSpinnerParserCOMMAND_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(205)
		p.Match(YarnSpinnerParserCOMMAND_ENDIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(206)
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
	p.EnterRule(localctx, 30, YarnSpinnerParserRULE_if_clause)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(208)
		p.Match(YarnSpinnerParserCOMMAND_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(209)
		p.Match(YarnSpinnerParserCOMMAND_IF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(210)
		p.expression(0)
	}
	{
		p.SetState(211)
		p.Match(YarnSpinnerParserCOMMAND_END)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(215)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(212)
				p.Statement()
			}

		}
		p.SetState(217)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 32, YarnSpinnerParserRULE_else_if_clause)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(218)
		p.Match(YarnSpinnerParserCOMMAND_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(219)
		p.Match(YarnSpinnerParserCOMMAND_ELSEIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(220)
		p.expression(0)
	}
	{
		p.SetState(221)
		p.Match(YarnSpinnerParserCOMMAND_END)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(225)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(222)
				p.Statement()
			}

		}
		p.SetState(227)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 34, YarnSpinnerParserRULE_else_clause)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(228)
		p.Match(YarnSpinnerParserCOMMAND_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(229)
		p.Match(YarnSpinnerParserCOMMAND_ELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(230)
		p.Match(YarnSpinnerParserCOMMAND_END)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(234)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(231)
				p.Statement()
			}

		}
		p.SetState(236)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 36, YarnSpinnerParserRULE_set_statement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(237)
		p.Match(YarnSpinnerParserCOMMAND_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(238)
		p.Match(YarnSpinnerParserCOMMAND_SET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(239)
		p.Variable()
	}
	{
		p.SetState(240)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*Set_statementContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&136341589327872) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*Set_statementContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(241)
		p.expression(0)
	}
	{
		p.SetState(242)
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
	p.EnterRule(localctx, 38, YarnSpinnerParserRULE_call_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(244)
		p.Match(YarnSpinnerParserCOMMAND_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(245)
		p.Match(YarnSpinnerParserCOMMAND_CALL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(246)
		p.Function_call()
	}
	{
		p.SetState(247)
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
	p.EnterRule(localctx, 40, YarnSpinnerParserRULE_command_statement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(249)
		p.Match(YarnSpinnerParserCOMMAND_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(250)
		p.Command_formatted_text()
	}
	{
		p.SetState(251)
		p.Match(YarnSpinnerParserCOMMAND_TEXT_END)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	p.SetState(255)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == YarnSpinnerParserHASHTAG {
		{
			p.SetState(252)
			p.Hashtag()
		}

		p.SetState(257)
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
	p.EnterRule(localctx, 42, YarnSpinnerParserRULE_command_formatted_text)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(265)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == YarnSpinnerParserCOMMAND_EXPRESSION_START || _la == YarnSpinnerParserCOMMAND_TEXT {
		p.SetState(263)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case YarnSpinnerParserCOMMAND_TEXT:
			{
				p.SetState(258)
				p.Match(YarnSpinnerParserCOMMAND_TEXT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case YarnSpinnerParserCOMMAND_EXPRESSION_START:
			{
				p.SetState(259)
				p.Match(YarnSpinnerParserCOMMAND_EXPRESSION_START)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(260)
				p.expression(0)
			}
			{
				p.SetState(261)
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

		p.SetState(267)
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
	p.EnterRule(localctx, 44, YarnSpinnerParserRULE_shortcut_option_statement)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(271)
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
				p.SetState(268)
				p.Shortcut_option()
			}

		}
		p.SetState(273)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

	{
		p.SetState(274)
		p.Shortcut_option()
	}
	p.SetState(276)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == YarnSpinnerParserBLANK_LINE_FOLLOWING_OPTION {
		{
			p.SetState(275)
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
	p.EnterRule(localctx, 46, YarnSpinnerParserRULE_shortcut_option)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(278)
		p.Match(YarnSpinnerParserSHORTCUT_ARROW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(279)
		p.Line_statement()
	}
	p.SetState(288)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(280)
			p.Match(YarnSpinnerParserINDENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(284)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1163266) != 0 {
			{
				p.SetState(281)
				p.Statement()
			}

			p.SetState(286)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(287)
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
	Value() IValueContext
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

func (s *Declare_statementContext) Value() IValueContext {
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
	p.EnterRule(localctx, 48, YarnSpinnerParserRULE_declare_statement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(290)
		p.Match(YarnSpinnerParserCOMMAND_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(291)
		p.Match(YarnSpinnerParserCOMMAND_DECLARE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(292)
		p.Variable()
	}
	{
		p.SetState(293)
		p.Match(YarnSpinnerParserOPERATOR_ASSIGNMENT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(294)
		p.Value()
	}
	p.SetState(297)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == YarnSpinnerParserEXPRESSION_AS {
		{
			p.SetState(295)
			p.Match(YarnSpinnerParserEXPRESSION_AS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(296)

			var _m = p.Match(YarnSpinnerParserFUNC_ID)

			localctx.(*Declare_statementContext).type_ = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(299)
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

func (p *YarnSpinnerParser) Jump_statement() (localctx IJump_statementContext) {
	localctx = NewJump_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, YarnSpinnerParserRULE_jump_statement)
	p.SetState(312)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		localctx = NewJumpToNodeNameContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(301)
			p.Match(YarnSpinnerParserCOMMAND_START)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(302)
			p.Match(YarnSpinnerParserCOMMAND_JUMP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(303)

			var _m = p.Match(YarnSpinnerParserID)

			localctx.(*JumpToNodeNameContext).destination = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(304)
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
			p.SetState(305)
			p.Match(YarnSpinnerParserCOMMAND_START)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(306)
			p.Match(YarnSpinnerParserCOMMAND_JUMP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(307)
			p.Match(YarnSpinnerParserEXPRESSION_START)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(308)
			p.expression(0)
		}
		{
			p.SetState(309)
			p.Match(YarnSpinnerParserEXPRESSION_END)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(310)
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

func (p *YarnSpinnerParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 10:
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
