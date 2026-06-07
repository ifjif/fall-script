package parser

import (
	"strconv"

	. "zzc/fall-script/src/ast"
	. "zzc/fall-script/src/token"
)

func (p *Parser) parseExpr(precedence Precedence) ExprNode {
	prefix := p.prefixParseFns[p.curType()]

	if prefix == nil {
		p.noPrefixParseFnErr(p.curType())
		return nil
	}

	leftExpr := prefix()

	for p.curType() != SEMICOLON && p.peekPrecedence() > precedence {
		infix := p.infixParseFns[p.nexType()]
		if infix == nil {
			return leftExpr
		}

		p.nextToken()
		leftExpr = infix(leftExpr)
	}

	return leftExpr
}

func (p *Parser) parseAssignExpr(left ExprNode) ExprNode {
	ident, ok := left.(*IdentExpr)
	if !ok {
		p.expectedIdentifierErr(left.GetToken())
		return nil
	}
	expr := &AssignExpr{Token: p.curToken, Name: ident}
	p.nextToken()
	expr.Value = p.parseExpr(LOWEST)

	return expr
}

func (p *Parser) parseIntExpr() ExprNode {
	expr := &IntExpr{Token: p.curToken}
	value, err := strconv.ParseInt(p.curToken.Value, 0, 64)
	if err != nil {
		p.intParseErr(p.curToken)
	}
	expr.Value = value
	return expr
}

func (p *Parser) parseBoolExpr() ExprNode {
	expr := &BoolExpr{Token: p.curToken}
	expr.Value = p.curTypeIs(TRUE)
	return expr
}

func (p *Parser) parseStringExpr() ExprNode {
	expr := &StrExpr{Token: p.curToken, Value: p.curToken.Value}
	return expr
}

func (p *Parser) parseInfixExpr(left ExprNode) ExprNode {
	expr := &InfixExpr{Left: left, Token: p.curToken, Op: p.curToken.Value}

	precedence := p.curPrecedence()
	p.nextToken()

	expr.Right = p.parseExpr(precedence)

	return expr
}

func (p *Parser) parsePrefixExpr() ExprNode {
	expr := &PrefixExpr{Token: p.curToken, Op: p.curToken.Value}

	p.nextToken()
	expr.Right = p.parseExpr(PREFIX)

	return expr
}

func (p *Parser) parseGroupExpr() ExprNode {
	p.nextToken()
	expr := p.parseExpr(LOWEST)

	if !p.expectPeek(RPAREN) {
		return nil
	}

	return expr
}

func (p *Parser) parseIdentExpr() ExprNode {
	expr := &IdentExpr{Token: p.curToken, Value: p.curToken.Value}
	return expr
}

func (p *Parser) parseCallExpr(left ExprNode) ExprNode {
	expr := &CallExpr{Token: p.curToken, Fn: left}

	expr.Args = p.parseExprList(RPAREN)

	return expr
}

func (p *Parser) parseIndexExpr(left ExprNode) ExprNode {
	expr := &IndexExpr{Token: p.curToken, Left: left}

	p.nextToken()
	expr.Index = p.parseExpr(LOWEST)

	if !p.expectPeek(RBRACKET) {
		return nil
	}

	return expr
}

func (p *Parser) parseFnExpr() ExprNode {
	expr := &FnExpr{Token: p.curToken}

	if p.peekTypeIs(IDENT) {
		p.nextToken()
		ident := p.parseIdentExpr().(*IdentExpr)
		expr.Name = ident.Value
	}

	if !p.expectPeek(LPAREN) {
		return nil
	}

	expr.Params = p.parseFnParams(RPAREN)

	if !p.expectPeek(LBRACE) {
		return nil
	}

	expr.Body = p.parseBlockStmt()

	return expr
}

func (p *Parser) parseArrExpr() ExprNode {
	expr := &ArrExpr{Token: p.curToken}

	expr.Elements = p.parseExprList(RBRACKET)

	return expr
}

func (p *Parser) parseHashExpr() ExprNode {
	expr := &HashExpr{Token: p.curToken}

	expr.Pairs = p.parseHashList(RBRACE)

	return expr
}

func (p *Parser) parseIfExpr() ExprNode {
	expr := &IfExpr{Token: p.curToken}

	if !p.expectPeek(LPAREN) {
		return nil
	}
	p.nextToken()

	expr.Condition = p.parseExpr(LOWEST)

	if !p.expectPeek(RPAREN) {
		return nil
	}

	if !p.expectPeek(LBRACE) {
		return nil
	}

	expr.Consequence = p.parseBlockStmt()

	if !p.peekTypeIs(ELSE) {
		return expr
	}

	p.nextToken()

	if p.peekTypeIs(IF) {
		stmts := []StmtNode{}
		p.nextToken()
		token := p.curToken
		stmt := p.parseExprStmt()
		stmts = append(stmts, stmt)
		expr.Alternative = &BlockStmt{Token: token, Stmts: stmts}
		return expr
	}

	if !p.expectPeek(LBRACE) {
		return nil
	}

	expr.Alternative = p.parseBlockStmt()

	return expr
}

func (p *Parser) parseNullExpr() ExprNode {
	return &NullExpr{Token: p.curToken}
}

func (p *Parser) parseExprList(end TokenType) []ExprNode {
	list := []ExprNode{}

	if p.peekTypeIs(end) {
		p.nextToken()
		return list
	}

	p.nextToken()
	list = append(list, p.parseExpr(LOWEST))

	for p.peekTypeIs(COMMA) {
		p.nextToken()
		p.nextToken()
		list = append(list, p.parseExpr(LOWEST))
	}

	if !p.expectPeek(end) {
		return nil
	}

	return list
}

func (p *Parser) parseHashList(end TokenType) []*Pair {
	list := []*Pair{}

	if p.peekTypeIs(end) {
		p.nextToken()
		return list
	}

	p.nextToken()

	pair := p.parseHashPair()
	if pair == nil {
		return nil
	}

	list = append(list, pair)

	for p.peekTypeIs(COMMA) {
		p.nextToken()
		p.nextToken()
		pair := p.parseHashPair()
		if pair == nil {
			return nil
		}
		list = append(list, pair)
	}

	if !p.expectPeek(end) {
		return nil
	}

	return list
}

func (p *Parser) parseHashPair() *Pair {
	pair := &Pair{}
	pair.Key = p.parseExpr(LOWEST)
	if !p.expectPeek(COLON) {
		return nil
	}
	p.nextToken()
	pair.Value = p.parseExpr(LOWEST)
	if pair.Value == nil {
		p.missValueErr(p.curToken)
	}

	return pair
}

func (p *Parser) parseFnParams(end TokenType) []*IdentExpr {
	list := []*IdentExpr{}

	if p.peekTypeIs(end) {
		p.nextToken()
		return list
	}

	if p.expectPeek(IDENT) {
		ident := p.parseIdentExpr().(*IdentExpr)
		list = append(list, ident)
	} else {
		return nil
	}

	for p.peekTypeIs(COMMA) {
		p.nextToken()
		if p.expectPeek(IDENT) {
			ident := p.parseIdentExpr().(*IdentExpr)
			list = append(list, ident)
		} else {
			return nil
		}
	}

	if !p.expectPeek(end) {
		return nil
	}

	return list
}
