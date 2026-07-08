package parser

import (
	"zzc/fall-script/src/ast"
	. "zzc/fall-script/src/token"
)

func (p *Parser) getCurIdent() *ast.IdentExpr {
	return &ast.IdentExpr{Token: p.curToken, Value: p.curToken.Value}
}

func (p *Parser) curType() TokenType {
	return p.curToken.Type
}

func (p *Parser) nexType() TokenType {
	return p.nexToken.Type
}

func (p *Parser) curTypeIs(t TokenType) bool {
	return p.curType() == t
}

func (p *Parser) expectCurIdent(v string) bool {
	if p.curToken.Value == v {
		return true
	}

	expected := Token{Type: IDENT, Value: v}
	p.expectedIdentifierErr(expected)

	return false
}

func (p *Parser) expectCur(tt TokenType) bool {
	if p.curTypeIs(tt) {
		return true
	}

	p.curErr(tt)
	return false
}

func (p *Parser) peekTypeIs(tt TokenType) bool {
	return p.nexToken.Type == tt
}

func (p *Parser) expectPeek(tt TokenType) bool {
	if p.peekTypeIs(tt) {
		p.nextToken()
		return true
	}

	p.peekErr(tt)
	return false
}

func (p *Parser) nextToken() {
	p.curToken = p.nexToken
	p.nexToken = p.l.NextToken()
}
