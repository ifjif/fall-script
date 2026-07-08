package compiler

import (
	"fmt"

	"zzc/fall-script/src/ast"
	. "zzc/fall-script/src/ast"
	. "zzc/fall-script/src/code"
	"zzc/fall-script/src/object"
	"zzc/fall-script/src/utils"
)

type Compiler struct {
	SymbolTable *SymbolTable
	program     Node
	curNode     Node
	errors      []string
	scopes      []*Scope
	scopeIndex  int
	imports     []*object.ImportRef
	exports     []*object.ExportRef
	structAsts  map[string]*ast.StructDeclStmt
	structs     map[string]*object.StructMeta
	methods     map[string][]*ast.MethodDeclExpr
	exportNames []string
}

func NewCompiler(program Node, imports []*ImportStmt, exports []*ExportStmt, pst *SymbolTable) *Compiler {
	st := NewEnclosedSymbolTable(pst)
	mainScope := NewScope()

	c := &Compiler{
		program:     program,
		SymbolTable: st,
		scopes:      []*Scope{mainScope},
		scopeIndex:  0,
	}
	// 定义 import
	c.defineImports(imports)

	// 定义exports
	c.defineExports(exports)

	return c
}

func (c *Compiler) SetStructAsts(structs map[string]*ast.StructDeclStmt) {
	c.structAsts = structs
}

func (c *Compiler) SetMethods(methods map[string][]*ast.MethodDeclExpr) {
	c.methods = methods
}

func (c *Compiler) defineExport(name string) *object.ExportRef {
	idx := c.addStrConstant(name)
	exp := &object.ExportRef{Name: idx, GlobalId: -1}
	return exp
}

func (c *Compiler) Compile() {
	// struct
	c.resolveStructs()

	c.doCompile(c.program)
	// 完善exports
	c.resolveExports()
}

func (c *Compiler) MainModule() *object.Module {
	cf := c.MainFn()
	mo := &object.Module{
		Name:      "",
		Imports:   c.imports,
		Exports:   c.exports,
		Cf:        cf,
		GlobalNum: c.SymbolTable.GlobalNum(),
	}

	return mo
}

func (c *Compiler) MainFn() *object.CompiledFunction {
	cf := &object.CompiledFunction{
		StackDepth:   c.currentScope().MaxStackDepth,
		LocalsNum:    c.SymbolTable.maxNum,
		Constants:    c.CurrentConstant(),
		Instructions: c.CurrentInstructions(),
	}

	return cf
}

func (c *Compiler) doCompile(node Node) {
	c.curNode = node

	switch node := node.(type) {
	case *NullExpr:
		c.compileNullExpr(node)
	case *Program:
		c.compileProgram(node.Stmts)
	case *BlockStmt:
		c.compileBlockStmt(node)
	case *ExprStmt:
		c.compileExprStmt(node)
	case *LetStmt:
		c.compileLetStmt(node)
	case *ForStmt:
		c.compileForStmt(node)
	case *WhileStmt:
		c.compileWhileStmt(node)
	case *DoWhileStmt:
		c.compileDoWhileStmt(node)
	case *ReturnStmt:
		c.compileReturnStmt(node)
	case *StructDeclStmt:
		c.compileStructDeclStmt(node)
	case *StructLiteralExpr:
		c.compileStructLieralExpr(node)
	case *AssignExpr:
		c.compileAssignExpr(node)
	case *IntExpr:
		c.compileIntExpr(node)
	case *StrExpr:
		c.compileStrExpr(node)
	case *BoolExpr:
		c.compileBoolExpr(node)
	case *ArrExpr:
		c.compileArrExpr(node)
	case *HashExpr:
		c.compileHashExpr(node)
	case *IdentExpr:
		c.compileIdentExpr(node)
	case *InfixExpr:
		c.compileInfixExpr(node)
	case *PrefixExpr:
		c.compilePrefixExpr(node)
	case *IndexExpr:
		c.compileIndexExpr(node)
	case *CallExpr:
		c.compileCallExpr(node)
	case *IfExpr:
		c.compileIfExpr(node)
	case *FnExpr:
		c.compileFnExpr(node)
	case *MethodDeclExpr:
		c.compileMethodDeclExpr(node)
	case *SliceExpr:
		c.compileSliceExpr(node)
	case *MemberExpr:
		c.compileMemberExpr(node)
	}
}

func (c *Compiler) compileNullExpr(expr *NullExpr) {
	c.emit(Null_)
}

func (c *Compiler) compileProgram(stmts []StmtNode) {
	for _, stmt := range stmts {
		c.doCompile(stmt)
	}

	if !c.lastInstructionIs(Return) && !c.lastInstructionIs(XReturn) {
		c.emit(Return)
	}
}

func (c *Compiler) compileBlockStmt(node *BlockStmt) {
	for _, stmt := range node.Stmts {
		c.doCompile(stmt)
	}
}

func (c *Compiler) compileExprStmt(stmt *ExprStmt) {
	c.doCompile(stmt.Expr)
	c.emit(Pop)
}

func (c *Compiler) compileLetStmt(stmt *LetStmt) {
	symbol := c.SymbolTable.Define(stmt.Name.Value)
	c.doCompile(stmt.Value)
	c.storeSymbol(symbol)
}

func (c *Compiler) compileForStmt(stmt *ForStmt) {
	c.enterBlock()

	if stmt.Start != nil {
		c.doCompile(stmt.Start)
	}

	conditonPos := len(c.CurrentInstructions())

	if stmt.Condition != nil {
		c.doCompile(stmt.Condition)
	} else {
		c.emit(True)
	}

	jifPos := c.emit(JumpIsFalse, 9999)
	c.doCompile(stmt.Body)
	if stmt.Update != nil {
		c.doCompile(stmt.Update)
		c.emit(Pop)
	}
	c.emit(Jump, conditonPos)
	c.changeOperand(jifPos, len(c.CurrentInstructions()))

	c.leaveBlock()
}

func (c *Compiler) compileWhileStmt(stmt *WhileStmt) {
	c.enterBlock()
	conditionPos := len(c.CurrentInstructions())
	c.doCompile(stmt.Condition)
	jifPos := c.emit(JumpIsFalse, 9999)
	c.doCompile(stmt.Body)
	c.emit(Jump, conditionPos)
	c.changeOperand(jifPos, len(c.CurrentInstructions()))
	c.leaveBlock()
}

func (c *Compiler) compileDoWhileStmt(stmt *DoWhileStmt) {
	c.enterBlock()
	startPos := len(c.CurrentInstructions())
	c.doCompile(stmt.Body)
	c.doCompile(stmt.Condition)
	jifPos := c.emit(JumpIsFalse, 9999)
	c.emit(Jump, startPos)
	c.changeOperand(jifPos, len(c.CurrentInstructions()))
	c.leaveBlock()
}

func (c *Compiler) compileReturnStmt(stmt *ReturnStmt) {
	if stmt.Value == nil {
		c.emit(Return)
		return
	}

	c.doCompile(stmt.Value)
	c.emit(XReturn)
}

func (c *Compiler) compileStructDeclStmt(stmt *StructDeclStmt) {
	// skip
}

func (c *Compiler) compileStructLieralExpr(expr *StructLiteralExpr) {
	c.compileIdentExpr(expr.Tag)
	count := len(expr.Elements)
	for _, elem := range expr.Elements {
		c.emit(Const, c.addStrConstant(elem.Key.Value))
		c.doCompile(elem.Value)
	}

	c.emit(InitStruct, count)
}

func (c *Compiler) compileAssignExpr(expr *AssignExpr) {
	left := expr.Left
	switch left := left.(type) {
	case *IdentExpr:
		sym, ok := c.SymbolTable.Resolve(left.Value)
		if !ok {
			// todo
			panic(fmt.Sprintf("Error: Identifier not found: %q", left.Value))
		}
		c.doCompile(expr.Value)
		c.storeSymbol(sym)
		c.loadSymbol(sym)
	case *IndexExpr:
		c.doCompile(left.Left)
		c.doCompile(left.Index)
		c.doCompile(expr.Value)
		c.emit(SetIndex)
	case *MemberExpr:
		c.doCompile(left.Visitor)
		member := left.Member.(*ast.IdentExpr)
		c.emit(Const, c.addStrConstant(member.Value))
		c.doCompile(expr.Value)
		c.emit(SetField)
	}
}

func (c *Compiler) compileIntExpr(expr *IntExpr) {
	c.emit(Const, c.addIntConstant(expr.Value))
}

func (c *Compiler) compileStrExpr(expr *StrExpr) {
	c.emit(Const, c.addStrConstant(expr.Value))
}

func (c *Compiler) compileBoolExpr(expr *BoolExpr) {
	if expr.Value {
		c.emit(True)
	} else {
		c.emit(False)
	}
}

func (c *Compiler) compileArrExpr(expr *ArrExpr) {
	for _, elem := range expr.Elements {
		c.doCompile(elem)
	}
	c.emit(Array_, len(expr.Elements))
}

func (c *Compiler) compileHashExpr(expr *HashExpr) {
	for _, pair := range expr.Pairs {
		c.doCompile(pair.Key)
		c.doCompile(pair.Value)
	}

	c.emit(Hash_, len(expr.Pairs)*2)
}

func (c *Compiler) compileIdentExpr(expr *IdentExpr) {
	name := expr.Value
	sym, ok := c.SymbolTable.Resolve(name)

	if !ok {
		// todo
		err := utils.IdentifierNotFoundErr(name)
		panic(err.Msg)
	}

	c.loadSymbol(sym)
}

func (c *Compiler) compileInfixExpr(expr *InfixExpr) {
	if expr.Op == "<" || expr.Op == "<=" {
		c.doCompile(expr.Right)
		c.doCompile(expr.Left)
		switch expr.Op {
		case "<":
			c.emit(Gt)
		case "<=":
			c.emit(Ge)
		}
		return
	}

	if expr.Op == "||" || expr.Op == "&&" {
		c.doCompile(expr.Left)
		if expr.Op == "||" {
			c.emit(Not)
		}
		jif := c.emit(JumpIsFalse, 9999)
		c.doCompile(expr.Right)
		if expr.Op == "&&" {
			c.emit(Not)
		}
		jif2 := c.emit(JumpIsFalse, 9999)
		c.changeOperand(jif, len(c.CurrentInstructions()))
		if expr.Op == "||" {
			c.emit(True)
		} else {
			c.emit(False)
		}
		ji := c.emit(Jump, 9999)
		c.changeOperand(jif2, len(c.CurrentInstructions()))
		if expr.Op == "||" {
			c.emit(False)
		} else {
			c.emit(True)
		}
		c.changeOperand(ji, len(c.CurrentInstructions()))
		return
	}

	c.doCompile(expr.Left)
	c.doCompile(expr.Right)
	switch expr.Op {
	case ">":
		c.emit(Gt)
	case ">=":
		c.emit(Ge)
	case "+":
		c.emit(Add)
	case "-":
		c.emit(Sub)
	case "*":
		c.emit(Mul)
	case "/":
		c.emit(Div)
	case "==":
		c.emit(Eq)
	case "!=":
		c.emit(Neq)
	}
}

func (c *Compiler) compilePrefixExpr(expr *PrefixExpr) {
	c.doCompile(expr.Right)

	switch expr.Op {
	case "!":
		c.emit(Not)
	case "-":
		c.emit(Neg)
	}
}

func (c *Compiler) compileIndexExpr(expr *IndexExpr) {
	c.doCompile(expr.Left)
	index := expr.Index
	c.doCompile(index)

	if _, ok := index.(*SliceExpr); ok {
		c.emit(Slice)
	} else {
		c.emit(Index)
	}
}

// [strct][a]
func (c *Compiler) compileCallExpr(expr *CallExpr) {
	fn := expr.Fn
	opcode := Call
	ma, ok := fn.(*ast.MemberExpr)
	if ok {
		c.doCompile(ma.Visitor)
		me := ma.Member.(*ast.IdentExpr)
		c.emit(Const, c.addStrConstant(me.Value))
		opcode = CallMethod
	} else {
		c.doCompile(fn)
	}
	for _, arg := range expr.Args {
		c.doCompile(arg)
	}

	c.emit(opcode, len(expr.Args))
}

func (c *Compiler) compileIfExpr(expr *IfExpr) {
	c.doCompile(expr.Condition)

	jumpIsFalsePos := c.emit(JumpIsFalse, 9999)

	c.enterBlock()
	c.doCompile(expr.Consequence)
	c.leaveBlock()

	if c.lastInstructionIs(Pop) {
		c.removeLastPopInst()
	} else {
		c.emit(Null_)
	}

	jumpPos := c.emit(Jump, 9999)

	c.changeOperand(jumpIsFalsePos, len(c.CurrentInstructions()))

	if expr.Alternative == nil {
		c.emit(Null_)
	} else {
		c.enterBlock()
		c.doCompile(expr.Alternative)
		c.leaveBlock()

		if c.lastInstructionIs(Pop) {
			c.removeLastPopInst()
		} else {
			c.emit(Null_)
		}
	}

	// if else 只返回一个 值，操作数栈-1
	c.updateScopeStackDepth(-1)

	c.changeOperand(jumpPos, len(c.CurrentInstructions()))
}

func (c *Compiler) compileFnExpr(expr *FnExpr) {
	var mSym Symbol
	if !expr.UnName {
		mSym = c.SymbolTable.Define(expr.Name)
	}

	c.enterScope()
	if expr.Name != "" {
		c.SymbolTable.defineFunction(expr.Name)
	}

	for _, param := range expr.Params {
		c.SymbolTable.Define(param.Value)
	}

	c.doCompile(expr.Body)
	if c.lastInstructionIs(Pop) {
		c.replaceLastPopWithXReturn()
	}
	if !c.lastInstructionIs(XReturn) && !c.lastInstructionIs(Return) {
		c.emit(Return)
	}

	consts := c.CurrentConstant()
	localNums := c.SymbolTable.maxNum
	freeTable := c.SymbolTable.FreeSymbols
	maxStackDepth := c.currentScope().MaxStackDepth
	insts := c.leaveScope()

	cf := &object.CompiledFunction{
		Instructions: insts,
		Constants:    consts,
		ParamsNum:    len(expr.Params),
		LocalsNum:    localNums,
		StackDepth:   maxStackDepth,
	}

	cfIdx := c.addConstant(cf)

	for _, sym := range freeTable {
		switch sym.Scope {
		case LOCAL:
			c.emit(GetLocal, sym.Pos)
			c.emit(NewBoxLocal, sym.Pos)
			c.emit(GetLocal, sym.Pos)
		case FREE:
			c.emit(GetFreeRaw, sym.Pos)
		}
	}

	c.emit(Closure_, cfIdx, len(freeTable))

	if !expr.UnName {
		c.emit(Dup)
		c.storeSymbol(mSym)
	}
}

func (c *Compiler) compileMethodDeclExpr(expr *MethodDeclExpr) {
	// skip
}

func (c *Compiler) compileSliceExpr(sliceExpr *SliceExpr) {
	start := sliceExpr.Start
	end := sliceExpr.End
	step := sliceExpr.Step
	capc := sliceExpr.Cap

	if start == nil {
		c.emit(Null_)
	} else {
		c.doCompile(start)
	}

	if end == nil {
		c.emit(Null_)
	} else {
		c.doCompile(end)
	}

	if step == nil {
		c.emit(Null_)
	} else {
		c.doCompile(step)
	}

	if capc == nil {
		c.emit(Null_)
	} else {
		c.doCompile(capc)
	}
}

func (c *Compiler) compileMemberExpr(expr *MemberExpr) {
	c.doCompile(expr.Visitor)
	member := expr.Member.(*ast.IdentExpr)
	c.emit(Const, c.addStrConstant(member.Value))

	c.emit(GetField)
}
