package parser

import (
	"zzc/fall-script/src/ast"
)

func (p *Parser) parseDeclaration() ast.Node {
	p.nextToken()
	node := p.parseStmt()
	return node
}
