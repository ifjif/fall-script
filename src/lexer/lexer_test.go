package lexer

import (
	"testing"

	. "zzc/fall-script/src/token"
)

func TestLexer(t *testing.T) {
	type testCase struct {
		input    string
		expected []Token
	}

	tests := []testCase{
		{
			`+ - * /
< <= > >=
= ! != ==
; : , { } ( ) [ ]
//fjdlsfkderewre
/*
	jfldsjfjdsf
	fjdslfjsdljf
			jfkdsjfsldjkj
*/ -
1 23 456 7890
"123abc"
let fn abc return true false if else
|| && | &
`,
			[]Token{
				{Type: PLUS, Value: "+", Line: 1, Col: 1},
				{Type: MINUS, Value: "-", Line: 1, Col: 3},
				{Type: ASTERISK, Value: "*", Line: 1, Col: 5},
				{Type: SLASH, Value: "/", Line: 1, Col: 7},
				{Type: LT, Value: "<", Line: 2, Col: 1},
				{Type: LE, Value: "<=", Line: 2, Col: 3},
				{Type: GT, Value: ">", Line: 2, Col: 6},
				{Type: GE, Value: ">=", Line: 2, Col: 8},
				{Type: ASSIGN, Value: "=", Line: 3, Col: 1},
				{Type: BANG, Value: "!", Line: 3, Col: 3},
				{Type: NEQ, Value: "!=", Line: 3, Col: 5},
				{Type: EQ, Value: "==", Line: 3, Col: 8},
				{Type: SEMICOLON, Value: ";", Line: 4, Col: 1},
				{Type: COLON, Value: ":", Line: 4, Col: 3},
				{Type: COMMA, Value: ",", Line: 4, Col: 5},
				{Type: LBRACE, Value: "{", Line: 4, Col: 7},
				{Type: RBRACE, Value: "}", Line: 4, Col: 9},
				{Type: LPAREN, Value: "(", Line: 4, Col: 11},
				{Type: RPAREN, Value: ")", Line: 4, Col: 13},
				{Type: LBRACKET, Value: "[", Line: 4, Col: 15},
				{Type: RBRACKET, Value: "]", Line: 4, Col: 17},
				{Type: MINUS, Value: "-", Line: 10, Col: 4},
				{Type: INT, Value: "1", Line: 11, Col: 1},
				{Type: INT, Value: "23", Line: 11, Col: 3},
				{Type: INT, Value: "456", Line: 11, Col: 6},
				{Type: INT, Value: "7890", Line: 11, Col: 10},
				{Type: STRING, Value: "123abc", Line: 12, Col: 1},
				{Type: LET, Value: "let", Line: 13, Col: 1},
				{Type: FUNCTION, Value: "fn", Line: 13, Col: 5},
				{Type: IDENT, Value: "abc", Line: 13, Col: 8},
				{Type: RETURN, Value: "return", Line: 13, Col: 12},
				{Type: TRUE, Value: "true", Line: 13, Col: 19},
				{Type: FALSE, Value: "false", Line: 13, Col: 24},
				{Type: IF, Value: "if", Line: 13, Col: 30},
				{Type: ELSE, Value: "else", Line: 13, Col: 33},
				{Type: OR, Value: "||", Line: 14, Col: 1},
				{Type: AND, Value: "&&", Line: 14, Col: 4},
				{Type: BOR, Value: "|", Line: 14, Col: 7},
				{Type: BAND, Value: "&", Line: 14, Col: 9},
				{Type: EOF, Value: "", Line: 15, Col: 1},
			},
		},
	}

	for i, tt := range tests {
		l := NewLexer(tt.input)

		for j, token := range tt.expected {
			tok := l.NextToken()
			if tok.Type != token.Type {
				t.Errorf("test [%d] token [%d]: expected token type is %q, got %q\n", i, j, token.Type, tok.Type)
			}

			if tok.Value != token.Value {
				t.Errorf("test [%d] token [%d]: expected token value is %q, got %q\n", i, j, token.Value, tok.Value)
			}

			if tok.Line != token.Line {
				t.Errorf("test [%d] token [%d]: expected token is located on line %d, got %d\n", i, j, token.Line, tok.Line)
			}

			if tok.Col != token.Col {
				t.Errorf("test [%d] token [%d]: expected token is located on column %d, got %d\n", i, j, token.Col, tok.Col)
			}
		}
	}
}
