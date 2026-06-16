package parser

import (
	. "zzc/fall-script/src/ast"
	. "zzc/fall-script/src/token"
)

func (p *Parser) parseAttributes() []*AttributeExpr {
	var attrs []*AttributeExpr
	for p.curType() == SHARP {
		attr := p.parseAttribute()
		if attr == nil {
			return nil
		}

		attrs = append(attrs, attr)
	}

	return attrs
}

func (p *Parser) parseAttribute() *AttributeExpr {
	attr := &AttributeExpr{
		Token: p.curToken,
	}

	if !p.expectPeek(LBRACKET) {
		return nil
	}

	if !p.expectPeek(IDENT) {
		return nil
	}

	name := p.curToken.Value
	attr.Name = name
	p.nextToken()

	if p.curType() == LPAREN {
		exprs := p.parseExprList(RPAREN)
		if exprs == nil {
			return nil
		}

		attr.Args = exprs
	}

	if !p.expectPeek(RBRACKET) {
		return nil
	}

	p.nextToken()

	return attr
}
