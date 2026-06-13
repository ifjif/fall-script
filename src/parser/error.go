package parser

import (
	"fmt"

	"zzc/fall-script/src/ast"
	"zzc/fall-script/src/token"
	. "zzc/fall-script/src/token"
)

func (p *Parser) curErr(tt token.TokenType) {
	msg := fmt.Sprintf("Error: expected '%s' token buf found '%s' at line %d, column %d", tt, p.curToken.Value, p.nexToken.Line, p.nexToken.Col)
	p.errors = append(p.errors, msg)
}

func (p *Parser) peekErr(tt token.TokenType) {
	msg := fmt.Sprintf("Error: expected next '%s' token buf found '%s' at line %d, column %d", tt, p.nexToken.Value, p.nexToken.Line, p.nexToken.Col)
	p.errors = append(p.errors, msg)
}

func (p *Parser) noPrefixParseFnErr(tt token.TokenType) {
	msg := fmt.Sprintf("Error: no prefix parse function for %s found", tt)
	p.errors = append(p.errors, msg)
}

func (p *Parser) intParseErr(token Token) {
	msg := p.parseErr(INT, token)
	p.errors = append(p.errors, msg)
}

func (p *Parser) expectedIdentifierErr(token Token) {
	msg := p.parseErr(ASSIGN, token)
	p.errors = append(p.errors, msg)
}

func (p *Parser) expectedLeftSideValue(expr ast.Node) bool {
	switch expr.(type) {
	case *ast.IdentExpr, *ast.IndexExpr:
		return true
	}
	curToken := expr.GetToken()
	msg := fmt.Sprintf("Error: Expected left side for assign is '%s' token , infix '%s' token, buf found '%s' at line %d, column %d",
		token.IDENT, token.LBRACKET, curToken.Type, curToken.Line, curToken.Col)
	p.errors = append(p.errors, msg)

	return false
}

func (p *Parser) parseErr(expectedTokentType TokenType, token Token) string {
	msg := fmt.Sprintf("Error: Expected '%s' token but found '%s' at line %d, column %d", expectedTokentType, token.Value, token.Line, token.Col)
	return msg
}

func (p *Parser) missValueErr(token Token) {
	msg := fmt.Sprintf("Error: Missing value in line %d, column %d", token.Line, token.Col)
	p.errors = append(p.errors, msg)
}
