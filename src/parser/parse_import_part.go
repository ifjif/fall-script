package parser

import (
	"zzc/fall-script/src/ast"
	"zzc/fall-script/src/token"
)

func (p *Parser) parseSpecifiers() []*ast.Specifier {
	//{ 或 *
	if !p.peekTypeIs(token.LBRACE) && !p.peekTypeIs(token.ASTERISK) {
		p.expectPeek(token.LBRACE)
		p.expectPeek(token.ASTERISK)
		return nil
	}

	specifiers := []*ast.Specifier{}

	p.nextToken()

	if p.curTypeIs(token.ASTERISK) {
		s := p.parseAllSpecifier()
		if s == nil {
			return nil
		}
		specifiers = append(specifiers, s)
		return specifiers
	}

	p.nextToken()
	s := p.parseSpecifier()
	if s == nil {
		return nil
	}

	specifiers = append(specifiers, s)
	for p.peekTypeIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		s := p.parseSpecifier()
		if s == nil {
			return nil
		}
		specifiers = append(specifiers, s)
	}

	if !p.expectPeek(token.RBRACE) {
		return nil
	}

	return specifiers
}

func (p *Parser) parseAllSpecifier() *ast.Specifier {
	specifier := &ast.Specifier{}
	specifier.Imported = p.curToken.Value
	local := p.parseSpecifierAlias()
	if local == nil {
		return nil
	}

	specifier.Local = local.String()
	return specifier
}

func (p *Parser) parseSpecifier() *ast.Specifier {
	specifier := &ast.Specifier{}
	if !p.expectCur(token.IDENT) {
		return nil
	}
	imported := p.parseIdentExpr()
	specifier.Imported = imported.String()

	if p.peekTypeIs(token.COMMA) || p.peekTypeIs(token.RBRACE) {
		specifier.Local = specifier.Imported
		return specifier
	}

	local := p.parseSpecifierAlias()
	if local == nil {
		return nil
	}
	specifier.Local = local.String()

	return specifier
}

func (p *Parser) parseSpecifierAlias() ast.ExprNode {
	if !p.expectPeek(token.IDENT) && !p.expectCurIdent("as") {
		return nil
	}

	if !p.expectPeek(token.IDENT) {
		return nil
	}
	local := p.parseIdentExpr()

	return local
}

func (p *Parser) parseSource() string {
	if !p.expectPeek(token.IDENT) && !p.expectCurIdent("from") {
		return ""
	}

	if !p.expectPeek(token.STRING) {
		return ""
	}
	str := p.parseStringExpr()

	return str.String()
}
