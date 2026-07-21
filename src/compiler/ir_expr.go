package compiler

import (
	"fmt"

	"zzc/fall-script/src/code"
	"zzc/fall-script/src/ir"
	"zzc/fall-script/src/object"
)

func (c *Compiler) compileExpr(expr ir.Expr) {
	switch expr := expr.(type) {
	case *ir.NullLiteral:
		c.compileNullLiteral(expr)
	case *ir.IntegerLiteral:
		c.compileIntegerLiteral(expr)
	case *ir.StringLiteral:
		c.compileStringLiteral(expr)
	case *ir.BoolLiteral:
		c.compileBoolLiteral(expr)
	case *ir.ArrayLiteral:
		c.compileArrayLiteral(expr)
	case *ir.HashLiteral:
		c.compileHashLiteral(expr)
	case *ir.Ident:
		c.compileIdent(expr)
	case *ir.InfixExpr:
		c.compileInfixExprIr(expr)
	case *ir.PrefixExpr:
		c.compilePrefixExprIr(expr)
	case *ir.AssignExpr:
		c.compileAssignExprIr(expr)
	case *ir.IndexExpr:
		c.compileIndexExprIr(expr)
	case *ir.MemberExpr:
		c.compileMemberExprIr(expr)
	case *ir.CallExpr:
		c.compileCallExprIr(expr)
	case *ir.SliceExpr:
		c.compileSliceExprIr(expr)
	case *ir.IfExpr:
		c.compileIfExprIr(expr)
	case *ir.FnExpr:
		c.compileFnExprIr(expr)
	case *ir.StructLiteral:
		c.compileStructLiteral(expr)
	case *ir.MatchExpr:
		c.compileMatchExpr(expr)
	}
}

func (c *Compiler) compileNullLiteral(expr *ir.NullLiteral) {
	c.emit(code.Null_)
}

func (c *Compiler) compileIntegerLiteral(expr *ir.IntegerLiteral) {
	idx := c.addIntConstant(expr.Value)
	c.emit(code.Const, idx)
}

func (c *Compiler) compileStringLiteral(expr *ir.StringLiteral) {
	idx := c.addStrConstant(expr.Value)
	c.emit(code.Const, idx)
}

func (c *Compiler) compileBoolLiteral(expr *ir.BoolLiteral) {
	if expr.Value {
		c.emit(code.True)
		return
	}
	c.emit(code.False)
}

func (c *Compiler) compileArrayLiteral(expr *ir.ArrayLiteral) {
	for _, elem := range expr.Elements {
		c.compileExpr(elem)
	}

	c.emit(code.Array_, len(expr.Elements))
}

func (c *Compiler) compileHashLiteral(expr *ir.HashLiteral) {
	for _, pair := range expr.Pairs {
		c.compileExpr(pair.Key)
		c.compileExpr(pair.Value)
	}

	c.emit(code.Hash_, len(expr.Pairs)*2)
}

func (c *Compiler) compileIdent(expr *ir.Ident) {
	c.loadSymbol(expr.Symbol)
}

func (c *Compiler) compileInfixExprIr(expr *ir.InfixExpr) {
	if expr.Op == "<" || expr.Op == "<=" {
		c.compileExpr(expr.Right)
		c.compileExpr(expr.Left)
		switch expr.Op {
		case "<":
			c.emit(code.Gt)
		case "<=":
			c.emit(code.Ge)
		}
		return
	}

	if expr.Op == "||" || expr.Op == "&&" {
		c.compileExpr(expr.Left)
		if expr.Op == "||" {
			c.emit(code.Not)
		}
		jif := c.emit(code.JumpIsFalse, 9999)
		c.compileExpr(expr.Right)
		if expr.Op == "&&" {
			c.emit(code.Not)
		}
		jif2 := c.emit(code.JumpIsFalse, 9999)
		c.changeOperand(jif, len(c.CurrentInstructions()))
		if expr.Op == "||" {
			c.emit(code.True)
		} else {
			c.emit(code.False)
		}
		ji := c.emit(code.Jump, 9999)
		c.changeOperand(jif2, len(c.CurrentInstructions()))
		if expr.Op == "||" {
			c.emit(code.False)
		} else {
			c.emit(code.True)
		}
		c.changeOperand(ji, len(c.CurrentInstructions()))
		return
	}

	c.compileExpr(expr.Left)
	c.compileExpr(expr.Right)
	switch expr.Op {
	case ">":
		c.emit(code.Gt)
	case ">=":
		c.emit(code.Ge)
	case "+":
		c.emit(code.Add)
	case "-":
		c.emit(code.Sub)
	case "*":
		c.emit(code.Mul)
	case "/":
		c.emit(code.Div)
	case "==":
		c.emit(code.Eq)
	case "!=":
		c.emit(code.Neq)
	case "&":
		c.emit(code.Band)
	case "|":
		c.emit(code.Bor)
	}
}

func (c *Compiler) compilePrefixExprIr(expr *ir.PrefixExpr) {
	c.compileExpr(expr.Right)

	switch expr.Op {
	case "!":
		c.emit(code.Not)
	case "-":
		c.emit(code.Neg)
	}
}

func (c *Compiler) compileAssignExprIr(expr *ir.AssignExpr) {
	left := expr.Left
	switch left := left.(type) {
	case *ir.Ident:
		c.compileExpr(expr.Value)
		c.storeSymbol(left.Symbol)
		c.loadSymbol(left.Symbol)
	case *ir.IndexExpr:
		c.compileExpr(left.Left)
		c.compileExpr(left.Index)
		c.compileExpr(expr.Value)
		c.emit(code.SetIndex)
	case *ir.MemberExpr:
		c.compileExpr(left.Visitor)
		c.compileExpr(left.Member)
		c.compileExpr(expr.Value)
		c.emit(code.SetField)
	}
}

func (c *Compiler) compileIndexExprIr(expr *ir.IndexExpr) {
	c.compileExpr(expr.Left)
	index := expr.Index
	c.compileExpr(index)

	if _, ok := index.(*ir.SliceExpr); ok {
		c.emit(code.Slice)
	} else {
		c.emit(code.Index)
	}
}

func (c *Compiler) compileMemberExprIr(expr *ir.MemberExpr) {
	c.compileExpr(expr.Visitor)
	c.compileExpr(expr.Member)
	c.emit(code.GetField)
}

func (c *Compiler) compileCallExprIr(expr *ir.CallExpr) {
	callee := expr.Callee
	opcode := code.Call
	if me, ok := callee.(*ir.MemberExpr); ok {
		opcode = code.CallMethod
		c.compileExpr(me.Visitor)
		c.compileExpr(me.Member)
	} else {
		c.compileExpr(callee)
	}

	for _, arg := range expr.Args {
		c.compileExpr(arg)
	}

	c.emit(opcode, len(expr.Args))
}

func (c *Compiler) compileSliceExprIr(expr *ir.SliceExpr) {
	c.compileExpr(expr.Start)
	c.compileExpr(expr.End)
	c.compileExpr(expr.Step)
	c.compileExpr(expr.Cap)
}

func (c *Compiler) compileIfExprIr(expr *ir.IfExpr) {
	c.compileExpr(expr.Condition)
	jumpIsFalsePos := c.emit(code.JumpIsFalse, 9999)
	c.compileBlockStmtIr(expr.Consequence)
	if c.lastInstructionIs(code.Pop) {
		c.removeLastPopInst()
	} else {
		c.emit(code.Null_)
	}

	jumpPos := c.emit(code.Jump, 9999)
	c.changeOperand(jumpIsFalsePos, len(c.CurrentInstructions()))
	if expr.Alternative == nil {
		c.emit(code.Null_)
	} else {
		c.compileBlockStmtIr(expr.Alternative)
		if c.lastInstructionIs(code.Pop) {
			c.removeLastPopInst()
		} else {
			c.emit(code.Null_)
		}
	}
	c.changeOperand(jumpPos, len(c.CurrentInstructions()))

	// if else 只返回一个 值，操作数栈-1
	c.updateScopeStackDepth(-1)
}

func (c *Compiler) compileFnExprIr(expr *ir.FnExpr) {
	c.enterScope()

	c.compileBlockStmtIr(expr.Body)
	if c.lastInstructionIs(code.Pop) {
		c.replaceLastPopWithXReturn()
	}
	if !c.lastInstructionIs(code.XReturn) && !c.lastInstructionIs(code.Return) {
		c.emit(code.Return)
	}
	constants := c.currentScope().Constants
	maxStack := c.currentScope().MaxStackDepth

	ins := c.leaveScope()

	for _, sym := range expr.Frees {
		switch sym.Scope {
		case ir.LOCAL:
			c.emit(code.GetLocal, sym.Pos)
			c.emit(code.NewBoxLocal, sym.Pos)
			c.emit(code.GetLocal, sym.Pos)
		case ir.FREE:
			c.emit(code.GetFreeRaw, sym.Pos)
		}
		// 设置  captured 为 true
		sym.Captured = true
	}

	f := &object.CompiledFunction{
		LocalsNum:    expr.LocalVars,
		Instructions: ins,
		Constants:    constants,
		ParamsNum:    len(expr.Params),
		StackDepth:   maxStack,
	}

	fidx := c.addConstant(f)

	c.emit(code.Closure_, fidx, len(expr.Frees))

	if !expr.UnName {
		c.emit(code.Dup)
		c.storeSymbol(expr.Name.Symbol)
	}
}

func (c *Compiler) compileStructLiteral(expr *ir.StructLiteral) {
	c.loadSymbol(expr.Tag.Symbol)
	for _, pair := range expr.Pairs {
		c.compileExpr(pair.Key)
		c.compileExpr(pair.Value)
	}
	c.emit(code.InitStruct, len(expr.Pairs))
}

// ===================================编译match
func (c *Compiler) compileMatchExpr(expr *ir.MatchExpr) {
	c.compileExpr(expr.Subject)
	arms := c.compileMatchArmExprs(expr.MatchArms)

	fmt.Println(arms)
	matchLastInstructionPos := len(c.CurrentInstructions())
	for _, arm := range arms {
		c.changeOperand(arm, matchLastInstructionPos)
	}

	// match只返回一个值
	c.updateScopeStackDepth(1 - len(arms))
}

func (c *Compiler) compileMatchArmExprs(arms []*ir.MatchArmExpr) []int {
	pjs := make([]int, 0)
	for _, arm := range arms {
		js := c.compileMatchArmExpr(arm)
		if js >= 0 {
			pjs = append(pjs, js)
		}
	}

	return pjs
}

func (c *Compiler) compileMatchArmExpr(arm *ir.MatchArmExpr) int {
	pattern := arm.Pattern
	_, ok := pattern.(*ir.WildcardPattern)

	pjif := -1
	gjif := -1
	if !ok {
		c.compilePattern(arm.Pattern)
		pjif = c.emit(code.JumpIsFalse, 9999)
		if arm.Guard != nil {
			c.compileExpr(arm.Guard)
			gjif = c.emit(code.JumpIsFalse, 9999)
		}
	}

	// 弹出subject
	c.emit(code.Pop)

	c.compileStmt(arm.Body)
	if c.lastInstructionIs(code.Pop) {
		c.removeLastPopInst()
	} else {
		c.emit(code.Null_)
	}

	if !ok {
		pj := c.emit(code.Jump, 9999)
		c.changeOperand(pjif, len(c.CurrentInstructions()))
		if gjif >= 0 {
			c.changeOperand(gjif, len(c.CurrentInstructions()))
		}

		return pj
	}

	return -1
}
