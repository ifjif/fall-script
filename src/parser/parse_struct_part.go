package parser

import (
	"zzc/fall-script/src/ast"
	"zzc/fall-script/src/token"
)

func (p *Parser) parseFieldDeclareds(end token.TokenType) []*ast.FieldDeclExpr {
	fields := make([]*ast.FieldDeclExpr, 0)

	for !p.peekTypeIs(end) {
		p.nextToken()
		field := p.parseFieldDeclared()
		if field == nil {
			return nil
		}

		fields = append(fields, field)
	}

	if !p.expectPeek(end) {
		return nil
	}

	return fields
}

func (p *Parser) parseFieldDeclared() *ast.FieldDeclExpr {
	if !p.curTypeIs(token.IDENT) {
		p.curErr(token.IDENT)
		return nil
	}

	field := &ast.FieldDeclExpr{Token: p.curToken}
	field.Name = p.parseIdentExpr().(*ast.IdentExpr)

	if !p.peekTypeIs(token.COLON) {
		field.IsEmbed = true
	} else {
		p.nextToken()
	}

	return field
}

func (p *Parser) parseStructPairs(end token.TokenType) []*ast.StructPair {
	pairs := make([]*ast.StructPair, 0)

	if p.peekTypeIs(end) {
		p.nextToken()
		return pairs
	}

	p.nextToken()

	pair := p.parseStructPair()
	if pair == nil {
		return nil
	}
	pairs = append(pairs, pair)

	for p.peekTypeIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		pair := p.parseStructPair()
		if pair == nil {
			return nil
		}
		pairs = append(pairs, pair)
	}

	if !p.expectPeek(end) {
		return nil
	}

	return pairs
}

func (p *Parser) parseStructPair() *ast.StructPair {
	pair := &ast.StructPair{}

	if !p.expectCur(token.IDENT) {
		return nil
	}

	pair.Key = p.parseIdentExpr().(*ast.IdentExpr)

	if !p.expectPeek(token.COLON) {
		return nil
	}

	p.nextToken()

	pair.Value = p.parseExpr(LOWEST)

	return pair
}
