package lexer

import (
	. "zzc/fall-script/src/token"
)

type Lexer struct {
	input string
	pos   int
	npos  int
	line  int
	col   int
	ch    byte
}

func NewLexer(input string) *Lexer {
	l := &Lexer{
		input: input,
		pos:   0,
		npos:  0,
		line:  1,
		col:   0,
	}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() Token {
	var tok Token
	l.skipWhitespace()

	switch l.ch {
	case 0:
		tok = NewToken(EOF, "", l.line, l.col)
	case '+':
		tok = NewToken(PLUS, string(l.ch), l.line, l.col)
	case '-':
		tok = NewToken(MINUS, string(l.ch), l.line, l.col)
	case '*':
		tok = NewToken(ASTERISK, string(l.ch), l.line, l.col)
	case '/':
		tok = l.tokenizeSlash()
	case '(':
		tok = NewToken(LPAREN, string(l.ch), l.line, l.col)
	case ')':
		tok = NewToken(RPAREN, string(l.ch), l.line, l.col)
	case '{':
		tok = NewToken(LBRACE, string(l.ch), l.line, l.col)
	case '}':
		tok = NewToken(RBRACE, string(l.ch), l.line, l.col)
	case '[':
		tok = NewToken(LBRACKET, string(l.ch), l.line, l.col)
	case ']':
		tok = NewToken(RBRACKET, string(l.ch), l.line, l.col)
	case ':':
		tok = NewToken(COLON, string(l.ch), l.line, l.col)
	case ';':
		tok = NewToken(SEMICOLON, string(l.ch), l.line, l.col)
	case ',':
		tok = NewToken(COMMA, string(l.ch), l.line, l.col)
	case '#':
		tok = NewToken(SHARP, string(l.ch), l.line, l.col)
	case '=':
		tok = l.tokenizeEq()
	case '<':
		tok = l.tokenizeLt()
	case '>':
		tok = l.tokenizeGt()
	case '!':
		tok = l.tokenizeBang()
	case '"':
		tok = l.tokenizeString()
	case '&':
		tok = l.tokenizeBand()
	case '|':
		tok = l.tokenizeBor()
	default:
		if l.isDigit() {
			tok = l.tokenizeNumber()
		} else if l.isLetter() {
			tok = l.tokenizeIdent()
		} else {
			tok = NewToken(ILLEGAL, string(l.ch), l.line, l.col)
		}
	}

	l.readChar()
	return tok
	// == >= <= !=  = > < !
}

func (l *Lexer) tokenizeBand() Token {
	ch := l.ch
	pos := l.pos
	line := l.line
	col := l.col

	if l.peekChar() == '&' {
		l.readChar()
		value := l.input[pos:l.npos]
		return NewToken(AND, value, line, col)
	}

	return NewToken(BAND, string(ch), line, col)
}

func (l *Lexer) tokenizeBor() Token {
	ch := l.ch
	pos := l.pos
	line := l.line
	col := l.col

	if l.peekChar() == '|' {
		l.readChar()
		value := l.input[pos:l.npos]
		return NewToken(OR, value, line, col)
	}

	return NewToken(BOR, string(ch), line, col)
}

func (l *Lexer) tokenizeIdent() Token {
	pos := l.pos
	line := l.line
	col := l.col

	for l.isWord() {
		l.readChar()
	}

	value := l.input[pos:l.pos]
	ttype := LookKeyword(value)
	tok := NewToken(ttype, value, line, col)
	l.reserveChar()

	return tok
}

func (l *Lexer) tokenizeNumber() Token {
	pos := l.pos
	line := l.line
	col := l.col

	for l.isDigit() {
		l.readChar()
	}

	tok := NewToken(INT, l.input[pos:l.pos], line, col)
	l.reserveChar()

	return tok
}

func (l *Lexer) tokenizeString() Token {
	line := l.line
	col := l.col
	l.readChar()
	pos := l.pos

	for l.ch != '"' {
		l.checkNewLine()
		l.readChar()
	}

	tok := NewToken(STRING, l.input[pos:l.pos], line, col)
	return tok
}

func (l *Lexer) tokenizeBang() Token {
	ch := l.ch
	line := l.line
	col := l.col
	pos := l.pos

	if l.peekChar() == '=' {
		l.readChar()
		neq := l.input[pos:l.npos]
		return NewToken(NEQ, string(neq), line, col)
	}

	return NewToken(BANG, string(ch), line, col)
}

/*
* tokenzie
* /
*
* 忽略 // 后的知道碰到 \n
* 忽略 `/ * * / 中的内容
 */
func (l *Lexer) tokenizeSlash() Token {
	ch := l.ch
	line := l.line
	col := l.col

	if l.peekChar() == '/' {
		for l.ch != '\n' && l.ch != 0 {
			l.readChar()
		}
		tok := l.NextToken()
		// 指向当前的token
		l.reserveChar()
		return tok
	}

	if l.peekChar() == '*' {
		l.readChar()
		l.readChar()
		for l.ch != 0 && (l.ch != '*' || l.peekChar() != '/') {
			l.checkNewLine()
			l.readChar()
		}
		l.readChar()
		l.readChar()
		tok := l.NextToken()
		// 指向当前的token
		l.reserveChar()
		return tok
	}

	return NewToken(SLASH, string(ch), line, col)
}

/*
* tokenize
* >
* >=
 */

func (l *Lexer) tokenizeGt() Token {
	ch := l.ch
	line := l.line
	col := l.col
	pos := l.pos

	if l.peekChar() == '=' {
		l.readChar()
		ge := l.input[pos:l.npos]
		return NewToken(GE, string(ge), line, col)
	}

	return NewToken(GT, string(ch), line, col)
}

/*
* tokenize
* <
* <=
 */
func (l *Lexer) tokenizeLt() Token {
	ch := l.ch
	line := l.line
	col := l.col
	pos := l.pos

	if l.peekChar() == '=' {
		l.readChar()
		le := l.input[pos:l.npos]
		return NewToken(LE, string(le), line, col)
	}

	return NewToken(LT, string(ch), line, col)
}

/*
* tokenize
*  =
*  ==
 */
func (l *Lexer) tokenizeEq() Token {
	ch := l.ch
	line := l.line
	col := l.col
	pos := l.pos

	if l.peekChar() == '=' {
		l.readChar()
		eq := l.input[pos:l.npos]
		return NewToken(EQ, string(eq), line, col)
	}

	return NewToken(ASSIGN, string(ch), line, col)
}

func (l *Lexer) readChar() {
	if l.npos > len(l.input) {
		return
	} else if l.npos == len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.npos]
	}
	l.pos = l.npos
	l.npos++
	l.col++
}

func (l *Lexer) reserveChar() {
	l.npos = l.pos
	l.pos--
	l.col--
	l.ch = l.input[l.pos]
}

func (l *Lexer) peekChar() byte {
	if l.npos >= len(l.input) {
		return 0
	}

	return l.input[l.npos]
}

func (l *Lexer) skipWhitespace() {
	for l.isWhitespace() {
		l.checkNewLine()
		l.readChar()
	}
}

func (l *Lexer) isWhitespace() bool {
	return l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r'
}

func (l *Lexer) isDigit() bool {
	return l.ch >= '0' && l.ch <= '9'
}

func (l *Lexer) isLetter() bool {
	return l.ch == '_' || (l.ch >= 'a' && l.ch <= 'z') || (l.ch >= 'A' && l.ch <= 'Z')
}

func (l *Lexer) isWord() bool {
	return l.isLetter() || l.isDigit()
}

func (l *Lexer) checkNewLine() {
	if l.ch == '\n' {
		l.line++
		l.col = 0
	}
}
