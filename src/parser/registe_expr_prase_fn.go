package parser

import (
	. "zzc/fall-script/src/ast"
	. "zzc/fall-script/src/token"
)

type (
	prefixParseFn func() ExprNode
	infixParseFn  func(ExprNode) ExprNode
)

func (p *Parser) registeExprFn() {
	p.prefixParseFns = make(map[TokenType]prefixParseFn)
	p.registePrefixFn(INT, p.parseIntExpr)
	p.registePrefixFn(IDENT, p.parseIdentExpr)
	p.registePrefixFn(TRUE, p.parseBoolExpr)
	p.registePrefixFn(FALSE, p.parseBoolExpr)
	p.registePrefixFn(STRING, p.parseStringExpr)
	p.registePrefixFn(BANG, p.parsePrefixExpr)
	p.registePrefixFn(MINUS, p.parsePrefixExpr)
	p.registePrefixFn(LPAREN, p.parseGroupExpr)
	p.registePrefixFn(FUNCTION, p.parseFnExpr)
	p.registePrefixFn(LBRACKET, p.parseArrExpr)
	p.registePrefixFn(LBRACE, p.parseHashExpr)
	p.registePrefixFn(IF, p.parseIfExpr)
	p.registePrefixFn(NULL, p.parseNullExpr)

	p.infixParseFns = make(map[TokenType]infixParseFn)
	p.registeInfixFn(ASSIGN, p.parseAssignExpr)
	p.registeInfixFn(PLUS, p.parseInfixExpr)
	p.registeInfixFn(MINUS, p.parseInfixExpr)
	p.registeInfixFn(ASTERISK, p.parseInfixExpr)
	p.registeInfixFn(SLASH, p.parseInfixExpr)
	p.registeInfixFn(LT, p.parseInfixExpr)
	p.registeInfixFn(LE, p.parseInfixExpr)
	p.registeInfixFn(GT, p.parseInfixExpr)
	p.registeInfixFn(GE, p.parseInfixExpr)
	p.registeInfixFn(EQ, p.parseInfixExpr)
	p.registeInfixFn(NEQ, p.parseInfixExpr)
	p.registeInfixFn(AND, p.parseInfixExpr)
	p.registeInfixFn(OR, p.parseInfixExpr)
	p.registeInfixFn(LPAREN, p.parseCallExpr)
	p.registeInfixFn(LBRACKET, p.parseIndexExpr)
	p.registeInfixFn(DOT, p.parseMemberExpr)
}

func (p *Parser) registePrefixFn(tokenType TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registeInfixFn(tokenType TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}
