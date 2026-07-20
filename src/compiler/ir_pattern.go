package compiler

import (
	"zzc/fall-script/src/code"
	"zzc/fall-script/src/ir"
)

func (c *Compiler) compilePattern(pattern ir.PatternNode) {
	switch pattern := pattern.(type) {
	case *ir.LiteralPattern:
		c.emit(code.Dup)
		c.compileExpr(pattern.Value)
		c.emit(code.Eq)
	}
}
