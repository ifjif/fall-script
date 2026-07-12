package compiler

import (
	"zzc/fall-script/src/code"
	"zzc/fall-script/src/ir"
	"zzc/fall-script/src/object"
)

type Compiler struct {
	scopes     []*Scope
	scopeIndex int
	imports    []*object.ImportRef
	exports    []*object.ExportRef
	structs    []*object.StructMeta
	program    *ir.Program
}

func NewCompiler(program *ir.Program) *Compiler {
	mainScope := NewScope()
	c := &Compiler{
		program: program,
		scopes:  []*Scope{mainScope},
		imports: []*object.ImportRef{},
		exports: []*object.ExportRef{},
		structs: []*object.StructMeta{},
	}

	return c
}

func (c *Compiler) Compiler() {
	for _, stmt := range c.program.Stmts {
		c.compileStmt(stmt)
	}

	if !c.lastInstructionIs(code.Return) && !c.lastInstructionIs(code.XReturn) {
		c.emit(code.Return)
	}
}

func (c *Compiler) MainModule() *object.Module {
	cf := c.MainFn()
	mo := &object.Module{
		Name:      "",
		Imports:   c.imports,
		Exports:   c.exports,
		Structs:   c.structs,
		Cf:        cf,
		GlobalNum: c.program.Globals,
	}

	return mo
}

func (c *Compiler) MainFn() *object.CompiledFunction {
	cf := &object.CompiledFunction{
		StackDepth:   c.currentScope().MaxStackDepth,
		LocalsNum:    c.program.Locals,
		Constants:    c.CurrentConstant(),
		Instructions: c.CurrentInstructions(),
	}

	return cf
}
