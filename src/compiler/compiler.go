package compiler

import (
	. "zzc/fall-script/src/ast"
	. "zzc/fall-script/src/code"
	. "zzc/fall-script/src/object"
)

type Compiler struct {
	Constants    []Object
	Instructions Instructions
	program      Node
	curNode      Node
}

func NewCompiler(program Node) *Compiler {
	return &Compiler{
		Constants:    make([]Object, 0),
		Instructions: Instructions{},
		program:      program,
	}
}

func (c *Compiler) Compile() {
	c.doCompile(c.program)
}

func (c *Compiler) doCompile(node Node) {
	c.curNode = node

	switch node := node.(type) {
	case *IntExpr:
		c.compileIntExpr(node)
	}
}

func (c *Compiler) compileIntExpr(expr *IntExpr) {
	value := &Integer{Value: expr.Value}
	c.emit(CONST, c.addConstant(value))
}
