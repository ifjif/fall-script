package compiler

import (
	"zzc/fall-script/src/code"
	"zzc/fall-script/src/object"
)

func (c *Compiler) emit(op code.OpCode, operands ...int) int {
	inst := code.Make(op, operands...)

	newPos := c.addInstruction(inst)
	c.setLastInstruction(op, newPos)

	delta := c.getStackDelta(op, operands)
	c.updateScopeStackDepth(delta)

	return newPos
}

func (c *Compiler) addInstruction(inst []byte) int {
	scope := c.currentScope()
	newPos := len(scope.Instructions)
	newInstructions := append(scope.Instructions, inst...)
	c.setInstructions(newInstructions)
	return newPos
}

func (c *Compiler) setLastInstruction(op code.OpCode, pos int) {
	scope := c.currentScope()
	scope.PrevInst = scope.LastInst
	scope.LastInst = code.EmittedInstruct{OpCode: op, Pos: pos}
}

func (c *Compiler) lastInstructionIs(op code.OpCode) bool {
	scope := c.currentScope()
	insts := scope.Instructions
	if len(insts) == 0 {
		return false
	}

	return scope.LastInst.OpCode == op
}

func (c *Compiler) removeLastPopInst() {
	scope := c.currentScope()
	last := scope.LastInst
	prev := scope.PrevInst

	newInsts := scope.Instructions[:last.Pos]
	scope.LastInst = prev
	scope.Instructions = newInsts
}

func (c *Compiler) changeOperand(pos int, operand int) {
	scope := c.currentScope()
	opcode := scope.Instructions[pos]
	inst := code.Make(code.OpCode(opcode), operand)
	c.replaceInstruction(pos, inst)
}

func (c *Compiler) replaceInstruction(pos int, inst []byte) {
	scope := c.currentScope()
	insts := scope.Instructions

	for i := 0; i < len(inst); i++ {
		insts[pos+i] = inst[i]
	}
}

func (c *Compiler) replaceLastPopWithXReturn() {
	lastInst := c.currentScope().LastInst
	c.replaceInstruction(lastInst.Pos, code.Make(code.XReturn))
	c.currentScope().LastInst.OpCode = code.XReturn
}

func (c *Compiler) addConstant(o object.Object) int {
	scope := c.currentScope()
	scope.Constants = append(scope.Constants, o)
	return len(scope.Constants) - 1
}

func (c *Compiler) currentScope() *Scope {
	return c.scopes[c.scopeIndex]
}

func (c *Compiler) CurrentInstructions() code.Instructions {
	return c.scopes[c.scopeIndex].Instructions
}

func (c *Compiler) setInstructions(insts code.Instructions) {
	c.scopes[c.scopeIndex].Instructions = insts
}

func (c *Compiler) CurrentConstant() []object.Object {
	return c.scopes[c.scopeIndex].Constants
}

func (c *Compiler) SetConstant(conts []object.Object) {
	c.scopes[c.scopeIndex].Constants = conts
}

func (c *Compiler) CurrentStackDepth() int {
	return c.scopes[c.scopeIndex].MaxStackDepth
}

func (c *Compiler) loadSymbol(sym Symbol) {
	switch sym.Scope {
	case GLOBAL:
		c.emit(code.GetGlobal, sym.Pos)
	case LOCAL:
		c.emit(code.GetLocal, sym.Pos)
	case FREE:
		c.emit(code.GetFree, sym.Pos)
	case BUILTIN:
		c.emit(code.GetBuiltin, sym.Pos)
	}
}

func (c *Compiler) storeSymbol(sym Symbol) {
	switch sym.Scope {
	case GLOBAL:
		c.emit(code.SetGlobal, sym.Pos)
	case LOCAL:
		c.emit(code.SetLocal, sym.Pos)
	}
}

func (c *Compiler) enterScope() {
	newScope := NewScope()
	c.scopes = append(c.scopes, newScope)
	c.scopeIndex++
	c.SymbolTable = NewEnclosedSymbolTableForFn(c.SymbolTable)
}

func (c *Compiler) leaveScope() code.Instructions {
	insts := c.CurrentInstructions()
	c.scopes = c.scopes[:c.scopeIndex]
	c.scopeIndex--
	c.SymbolTable = c.SymbolTable.outer

	return insts
}

func (c *Compiler) updateScopeStackDepth(incr int) {
	curScope := c.currentScope()
	curDepth := curScope.StackDepth + incr

	if curDepth > curScope.MaxStackDepth {
		curScope.MaxStackDepth = curDepth
	}
	curScope.StackDepth = curDepth
}

func (c *Compiler) enterBlock() {
	curSym := c.SymbolTable
	newSym := NewEnclosedSymbolTable(curSym)
	newSym.defNum = curSym.defNum
	c.SymbolTable = newSym
}

func (c *Compiler) leaveBlock() {
	curSym := c.SymbolTable
	outer := curSym.outer

	if outer != nil && outer.maxNum < curSym.maxNum {
		outer.maxNum = curSym.maxNum
	}
	c.SymbolTable = outer
}
