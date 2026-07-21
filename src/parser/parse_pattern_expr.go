package parser

import (
	"fmt"

	. "zzc/fall-script/src/ast"
)

func (p *Parser) parsePatternExpr() PatternNode {
	expr := p.parseExpr(LOWEST)

	return p.resolvePatternExpr(expr)
}

func (p *Parser) resolvePatternExpr(expr ExprNode) PatternNode {
	tok := expr.GetToken()

	if tok.Value == "_" {
		return &WildcardPattern{Token: tok}
	}

	switch expr := expr.(type) {
	case *IntExpr, *StrExpr, *NullExpr, *BoolExpr:
		return &LiteralPattern{Token: tok, Value: expr}
	case *InfixExpr:
		if expr.Op == "|" {
			left := p.resolvePatternExpr(expr.Left)
			right := p.resolvePatternExpr(expr.Right)
			return &InfixPattern{Token: tok, Left: left, Op: expr.Op, Right: right}
		}
	}

	msg := fmt.Sprintf("token type is %q 表达式无法成为match的pattern at line %d, column %d", tok.Type, tok.Line, tok.Col)
	p.errors = append(p.errors, msg)

	return nil
}
