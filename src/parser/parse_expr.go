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
	if !p.expectedLeftSideValue(left) {
		return nil
	}
	expr := &AssignExpr{Token: p.curToken, Left: left}
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
	if p.peekTypeIs(LBRACE) && p.exprKind != MATCH {
		p.nextToken()
		sl := p.parseStructLiteralExpr(expr)
		return sl
	}
	return expr
}

func (p *Parser) parseStructLiteralExpr(tag *IdentExpr) *StructLiteralExpr {
	expr := &StructLiteralExpr{Token: tag.GetToken(), Tag: tag}

	elements := p.parseStructPairs(RBRACE)

	expr.Elements = elements
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

	if p.curTypeIs(COLON) {
		expr.Index = p.parseSliceExpr(nil)
	} else {
		start := p.parseExpr(LOWEST)
		if p.peekTypeIs(COLON) {
			p.nextToken()
			start = p.parseSliceExpr(start)
		}
		expr.Index = start
	}

	if !p.expectPeek(RBRACKET) {
		return nil
	}

	return expr
}

func (p *Parser) parseFnExpr() ExprNode {
	expr := &FnExpr{
		Token:  p.curToken,
		Params: []*IdentExpr{},
	}
	var method *MethodDeclExpr

	if p.peekTypeIs(LPAREN) {
		p.nextToken()
		p.nextToken()

		// struct
		if p.curTypeIs(IDENT) && p.peekTypeIs(IDENT) {
			method = &MethodDeclExpr{}
			arg0 := p.parseIdentExpr().(*IdentExpr)
			p.nextToken()
			structName := p.parseIdentExpr().(*IdentExpr)

			if !p.expectPeek(RPAREN) {
				return nil
			}

			method.StructName = structName
			expr.Params = append(expr.Params, arg0)
			method.Fn = expr

			if !p.peekTypeIs(IDENT) {
				p.peekErr(IDENT)
				return nil
			}
		} else {
			expr.UnName = true
		}
	}

	if p.peekTypeIs(IDENT) {
		p.nextToken()
		ident := p.parseIdentExpr().(*IdentExpr)
		expr.Name = ident.Value
		expr.Ident = ident
	}

	if !expr.UnName && !p.expectPeek(LPAREN) {
		return nil
	}

	if p.curTypeIs(LPAREN) {
		p.nextToken()
	}

	params := p.parseFnParams(RPAREN)
	expr.Params = append(expr.Params, params...)

	if !p.expectPeek(LBRACE) {
		return nil
	}

	expr.Body = p.parseBlockStmt()

	if method != nil {
		structName := method.StructName.Value
		m, ok := p.methods[structName]
		if !ok {
			m = make([]*MethodDeclExpr, 0)
		}
		m = append(m, method)
		p.methods[structName] = m
		return method
	}
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

	if p.curTypeIs(end) {
		return list
	}

	if p.expectCur(IDENT) {
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

func (p *Parser) parseSliceExpr(start ExprNode) ExprNode {
	token := p.curToken
	node := &SliceExpr{Token: token, Start: start}

	if !p.peekTypeIs(COLON) && !p.peekTypeIs(RBRACKET) { // [x:x
		p.nextToken()
		node.End = p.parseExpr(LOWEST)
	}

	if p.peekTypeIs(COLON) { // [x:x:
		p.nextToken()

		if !p.peekTypeIs(COLON) && !p.peekTypeIs(RBRACKET) { // [x:x:x
			p.nextToken()
			node.Step = p.parseExpr(LOWEST)
		}

		if p.peekTypeIs(COLON) { // [x:x:x:
			p.nextToken()

			if !p.peekTypeIs(RBRACKET) { // [x:x:x:x
				p.nextToken()
				node.Cap = p.parseExpr(LOWEST)
			}
		}
	}

	return node
}

func (p *Parser) parseMemberExpr(left ExprNode) ExprNode {
	expr := &MemberExpr{Token: p.curToken, Visitor: left}
	precedences := p.curPrecedence()
	p.nextToken()
	// todo 可以是 Ident或int
	expr.Member = p.parseExpr(precedences)

	return expr
}

func (p *Parser) parseMatchExpr() ExprNode {
	expr := &MatchExpr{Token: p.curToken}
	p.exprKind = p.curToken.Type
	p.nextToken()

	// todo 带重做更完善(设置表达式作用域)
	expr.Subject = p.parseExpr(LOWEST)
	p.exprKind = ""

	if !p.expectPeek(LBRACE) {
		return nil
	}

	expr.Arms = p.parseMatchArms(RBRACE)
	return expr
}

func (p *Parser) parseMatchArms(end TokenType) []*MatchArmExpr {
	mas := make([]*MatchArmExpr, 0)

	for p.nexType() != end {
		p.nextToken()
		arm := p.parseMatchArm()
		if arm != nil {
			mas = append(mas, arm)
		}
	}

	if !p.expectPeek(end) {
		return nil
	}

	// 判断最后一个是否是 _
	last := mas[len(mas)-1]
	if _, ok := last.Pattern.(*WildcardPattern); !ok {
		p.errors = append(p.errors, "expected last pattern is '_' to match")
		return nil
	}

	return mas
}

// pattern if gurad => body [,]
func (p *Parser) parseMatchArm() *MatchArmExpr {
	expr := &MatchArmExpr{Token: p.curToken}
	pattern := p.parsePatternExpr()
	var guard ExprNode
	var body StmtNode

	if p.nexType() == IF {
		p.nextToken()
		p.nextToken()
		guard = p.parseExpr(LOWEST)
	}

	if !p.expectPeek(FAT_ARROW) {
		return nil
	}

	p.nextToken()
	if p.curType() == LBRACE {
		body = p.parseBlockStmt()
	} else {
		body = p.parseExprStmt()

		if p.nexType() == COMMA {
			p.nextToken()
		}
	}

	expr.Pattern = pattern
	expr.Guard = guard
	expr.Body = body

	return expr
}
