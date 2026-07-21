package compiler

import (
	"zzc/fall-script/src/code"
	"zzc/fall-script/src/ir"
)

type PatternEnv struct {
	OrLeftJp []int
}

func NewPatternEnv() *PatternEnv {
	return &PatternEnv{
		OrLeftJp: []int{},
	}
}

func (c *Compiler) compilePattern(pattern ir.PatternNode) int {
	env := NewPatternEnv()
	jifPos := c.doCompilePattern(pattern, env)
	// 处理 left jump offset
	for _, pos := range env.OrLeftJp {
		c.changeOperand(pos, len(c.CurrentInstructions()))
	}

	return jifPos
}

func (c *Compiler) doCompilePattern(pattern ir.PatternNode, env *PatternEnv) int {
	switch pattern := pattern.(type) {

	case *ir.LiteralPattern:
		c.emit(code.Dup)
		c.compileExpr(pattern.Value)
		c.emit(code.Eq)
		return c.emit(code.JumpIsFalse, 9999)

		// ((1 | 2) | 3) | 4(last)    [bool (jump_is_false next) (jump last) next]
	case *ir.InfixPattern:
		if pattern.Op == "|" {
			return c.compileInfixOrPattern(pattern, env)
		}
	}

	return -1
}

func (c *Compiler) compileInfixOrPattern(infix *ir.InfixPattern, env *PatternEnv) int {
	leftJp := c.compileInfixLeftPattern(infix.Left, env)
	rightJif := c.compileInfixRightPattern(infix.Right, env)

	env.OrLeftJp = append(env.OrLeftJp, leftJp)

	return rightJif
}

func (c *Compiler) compileInfixLeftPattern(left ir.PatternNode, env *PatternEnv) int {
	leftJif := c.doCompilePattern(left, env)
	leftJp := c.emit(code.Jump, 9999)
	c.changeOperand(leftJif, len(c.CurrentInstructions()))

	return leftJp
}

func (c *Compiler) compileInfixRightPattern(right ir.PatternNode, env *PatternEnv) int {
	return c.doCompilePattern(right, env)
}
