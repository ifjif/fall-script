package token

type TokenType string

const (
	ILLEGAL  = "ILLEGAL"
	EOF      = "EOF"
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"
	BANG     = "!"
	ASSIGN   = "="
	LT       = "<"
	LE       = "<="
	GT       = ">"
	GE       = ">="
	EQ       = "=="
	NEQ      = "!="
	AND      = "&&"
	OR       = "||"
	BAND     = "&"
	BOR      = "|"

	COLON     = ":"
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACKET  = "["
	RBRACKET  = "]"
	LBRACE    = "{"
	RBRACE    = "}"
	COMMA     = ","

	INT      = "INT"
	STRING   = "STRING"
	IDENT    = "IDENT"
	LET      = "LET"
	FUNCTION = "FUNCTION"
	IF       = "IF"
	ELSE     = "ELSE"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	RETURN   = "RETURN"
	MACRO    = "MACRO"
)

var keyword = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"macro":  MACRO,
}

type Token struct {
	Type  TokenType
	Value string
	Line  int
	Col   int
}

func NewToken(t TokenType, v string, line int, col int) Token {
	return Token{
		Type:  t,
		Value: v,
		Line:  line,
		Col:   col,
	}
}

func LookKeyword(value string) TokenType {
	if tt, ok := keyword[value]; ok {
		return tt
	}

	return IDENT
}
