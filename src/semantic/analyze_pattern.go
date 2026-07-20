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
	}
	return nil
}
