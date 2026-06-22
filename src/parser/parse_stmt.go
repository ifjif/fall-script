package parser

import (
	. "zzc/fall-script/src/ast"
	"zzc/fall-script/src/token"
	. "zzc/fall-script/src/token"
)

func (p *Parser) parseStmt() StmtNode {
	attrs := p.parseAttributes()
	var stmt StmtNode
	switch p.curType() {
	case IMPORT:
		stmt = p.parseImportStmt()
	case EXPORT:
		stmt = p.parseExportStmt()
	case LET:
		stmt = p.parseLetStmt()
	case FOR:
		stmt = p.parseForStmt()
	case WHILE:
		stmt = p.parseWhileStmt()
	case DO:
		stmt = p.parseDoWhileStmt()
	case RETURN:
		stmt = p.parseReturnStmt()
	default:
		stmt = p.parseExprStmt()
	}

	if stmt != nil {
		stmt.SetAttributes(attrs)
	}
	return stmt
}

func (p *Parser) parseImportStmt() StmtNode {
	stmt := &ImportStmt{Token: p.curToken}
	specifiers := p.parseSpecifiers()
	source := p.parseSource()
	if specifiers == nil || source == "" {
		return nil
	}
	stmt.Specifiers = specifiers
	stmt.Source = source

	if p.peekTypeIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExportStmt() StmtNode {
	stmt := &ExportStmt{Token: p.curToken}
	declaration := p.parseDeclaration()
	if declaration == nil {
		return nil
	}
	stmt.Declaration = declaration
	return stmt
}

func (p *Parser) parseLetStmt() StmtNode {
	stmt := &LetStmt{Token: p.curToken}

	if !p.expectPeek(IDENT) {
		return nil
	}

	stmt.Name = p.parseIdentExpr().(*IdentExpr)

	if p.peekTypeIs(ASSIGN) {
		p.nextToken()
		p.nextToken()
		value := p.parseExpr(LOWEST)
		fn, ok := value.(*FnExpr)
		if ok && fn.UnName {
			fn.Name = stmt.Name.Value
		}
		stmt.Value = value
	}

	if p.peekTypeIs(SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseForStmt() StmtNode {
	stmt := &ForStmt{Token: p.curToken}

	if !p.expectPeek(LPAREN) {
		return nil
	}

	if !p.peekTypeIs(SEMICOLON) {
		p.nextToken()
		stmt.Start = p.parseStmt()
	} else {
		p.nextToken()
	}

	if !p.expectCur(SEMICOLON) {
		return nil
	}

	if !p.peekTypeIs(SEMICOLON) {
		p.nextToken()
		stmt.Condition = p.parseExpr(LOWEST)
	}

	if !p.expectPeek(SEMICOLON) {
		return nil
	}

	if !p.peekTypeIs(RPAREN) {
		p.nextToken()
		stmt.Update = p.parseExpr(LOWEST)
	}

	if !p.expectPeek(RPAREN) || !p.expectPeek(LBRACE) {
		return nil
	}

	stmt.Body = p.parseBlockStmt()

	return stmt
}

func (p *Parser) parseWhileStmt() StmtNode {
	stmt := &WhileStmt{Token: p.curToken}
	if !p.expectPeek(LPAREN) {
		return nil
	}

	p.nextToken()
	stmt.Condition = p.parseExpr(LOWEST)

	if !p.expectPeek(RPAREN) || !p.expectPeek(LBRACE) {
		return nil
	}

	stmt.Body = p.parseBlockStmt()

	return stmt
}

func (p *Parser) parseDoWhileStmt() StmtNode {
	stmt := &DoWhileStmt{Token: p.curToken}
	if !p.expectPeek(LBRACE) {
		return nil
	}

	stmt.Body = p.parseBlockStmt()

	if !p.expectPeek(WHILE) || !p.expectPeek(LPAREN) {
		return nil
	}

	p.nextToken()
	stmt.Condition = p.parseExpr(LOWEST)

	if !p.expectPeek(RPAREN) {
		return nil
	}

	return stmt
}

func (p *Parser) parseReturnStmt() StmtNode {
	stmt := &ReturnStmt{Token: p.curToken}
	if p.peekTypeIs(SEMICOLON) || p.peekTypeIs(RBRACE) || p.peekTypeIs(EOF) {
		if p.peekTypeIs(SEMICOLON) {
			p.nextToken()
		}
		return stmt
	}

	p.nextToken()
	stmt.Value = p.parseExpr(LOWEST)

	if p.peekTypeIs(SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExprStmt() StmtNode {
	stmt := &ExprStmt{Token: p.curToken}
	stmt.Expr = p.parseExpr(LOWEST)

	if p.peekTypeIs(SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseBlockStmt() *BlockStmt {
	bs := &BlockStmt{Token: p.curToken}
	stmts := []StmtNode{}

	p.nextToken()

	for !p.curTypeIs(RBRACE) && !p.curTypeIs(EOF) {
		stmt := p.parseStmt()
		if stmt != nil {
			stmts = append(stmts, stmt)
		}

		p.nextToken()
	}

	bs.Stmts = stmts

	return bs
}
