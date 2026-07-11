package compiler

import (
	"zzc/fall-script/src/code"
	"zzc/fall-script/src/ir"
)

func (c *Compiler) enterScope() {
	newScope := NewScope()
	c.scopes = append(c.scopes, newScope)
	c.scopeIndex++
}

func (c *Compiler) leaveScope() code.Instructions {
	insts := c.CurrentInstructions()
	c.scopes = c.scopes[:c.scopeIndex]
	c.scopeIndex--

	return insts
}

func (c *Compiler) loadSymbol(sym *ir.Symbol) {
	switch sym.Scope {
	case ir.GLOBAL:
		c.emit(code.GetGlobal, sym.Pos)
	case ir.LOCAL:
		if sym.Captured {
			c.emit(code.GetBoxLocal, sym.Pos)
		} else {
			c.emit(code.GetLocal, sym.Pos)
		}
	case ir.FREE:
		c.emit(code.GetFree, sym.Pos)
	case ir.BUILTIN:
		c.emit(code.GetBuiltin, sym.Pos)
	case ir.FN:
		c.emit(code.CurClosure)
	}
}

func (c *Compiler) storeSymbol(sym *ir.Symbol) {
	switch sym.Scope {
	case ir.GLOBAL:
		c.emit(code.SetGlobal, sym.Pos)
	case ir.LOCAL:
		if sym.Captured {
			c.emit(code.SetBoxLocal, sym.Pos)
		} else {
			c.emit(code.SetLocal, sym.Pos)
		}
	case ir.FREE:
		c.emit(code.SetFree, sym.Pos)
	}
}
