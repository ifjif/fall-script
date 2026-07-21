package semantic

import (
	"zzc/fall-script/src/ast"
	"zzc/fall-script/src/ir"
)

func (a *Analyzer) analyzePattern(pattern ast.PatternNode) ir.PatternNode {
	switch pattern := pattern.(type) {
	case *ast.WildcardPattern:
		return &ir.WildcardPattern{Token: pattern.Token}
	case *ast.LiteralPattern:
		value := a.analyzeExpr(pattern.Value)
		return &ir.LiteralPattern{Token: pattern.Token, Value: value}
	case *ast.InfixPattern:
		left := a.analyzePattern(pattern.Left)
		right := a.analyzePattern(pattern.Right)
		return &ir.InfixPattern{Token: pattern.Token, Left: left, Op: pattern.Op, Right: right}
	}
	return nil
}
