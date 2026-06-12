package compiler

import (
	. "zzc/fall-script/src/ast"
	"zzc/fall-script/src/builtin"
	_ "zzc/fall-script/src/builtin"
	"zzc/fall-script/src/code"
	. "zzc/fall-script/src/code"
	"zzc/fall-script/src/object"
	. "zzc/fall-script/src/object"
)

type Compiler struct {
	SymbolTable *SymbolTable
	Builtins    []*builtin.BuiltinDef
	program     Node
	curNode     Node
	errors      []string
	scopes      []*Scope
	scopeIndex  int
}

func NewCompiler(program Node) *Compiler {
	st := NewSymbolTable()
	mainScope := NewScope()

	for i, fn := range builtin.Builtins {
		st.defineBuiltin(i, fn.Name)
	}

	return &Compiler{
		program:     program,
		SymbolTable: st,
		scopes:      []*Scope{mainScope},
		scopeIndex:  0,
		Builtins:    builtin.Builtins,
	}
}

func (c *Compiler) Compile() {
	c.doCompile(c.program)
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
	}
}

func (c *Compiler) compileNullExpr(expr *NullExpr) {
	c.emit(code.Null_)
}

func (c *Compiler) compileProgram(stmts []StmtNode) {
	for _, stmt := range stmts {
		c.doCompile(stmt)
	}

	if !c.lastInstructionIs(code.Return) && !c.lastInstructionIs(code.XReturn) {
		c.emit(code.Return)
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
		c.emit(code.Return)
		return
	}

	c.doCompile(stmt.Value)
	c.emit(code.XReturn)
}

func (c *Compiler) compileAssignExpr(expr *AssignExpr) {
	sym, ok := c.SymbolTable.Resolve(expr.Name.Value)
	if !ok {
		// todo
	}
	c.doCompile(expr.Value)
	c.storeSymbol(sym)
	c.loadSymbol(sym)
}

func (c *Compiler) compileIntExpr(expr *IntExpr) {
	value := &Integer{Value: expr.Value}
	c.emit(Const, c.addConstant(value))
}

func (c *Compiler) compileStrExpr(expr *StrExpr) {
	value := &String{Value: expr.Value}
	c.emit(Const, c.addConstant(value))
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
		// Error
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
	c.doCompile(expr.Index)
	c.emit(Index)
}

func (c *Compiler) compileCallExpr(expr *CallExpr) {
	c.doCompile(expr.Fn)
	for _, arg := range expr.Args {
		c.doCompile(arg)
	}

	c.emit(Call, len(expr.Args))
}

func (c *Compiler) compileIfExpr(expr *IfExpr) {
	c.doCompile(expr.Condition)

	jumpIsFalsePos := c.emit(JumpIsFalse, 9999)

	c.enterBlock()
	c.doCompile(expr.Consequence)
	c.leaveBlock()

	if c.lastInstructionIs(Pop) {
		c.removeLastPopInst()
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
		}
	}

	c.changeOperand(jumpPos, len(c.CurrentInstructions()))
}

func (c *Compiler) compileFnExpr(expr *FnExpr) {
	var mSym Symbol
	if expr.Name != "" {
		mSym = c.SymbolTable.Define(expr.Name)
	}

	c.enterScope()

	for _, param := range expr.Params {
		c.SymbolTable.Define(param.Value)
	}

	c.doCompile(expr.Body)
	if c.lastInstructionIs(code.Pop) {
		c.replaceLastPopWithXReturn()
	}
	if !c.lastInstructionIs(code.XReturn) && !c.lastInstructionIs(code.Return) {
		c.emit(code.Return)
	}

	consts := c.CurrentConstant()
	localNums := c.SymbolTable.maxNum
	freeTable := c.SymbolTable.FreeSymbols
	maxStackDepth := c.currentScope().MaxStackDepth
	insts := c.leaveScope()

	for _, sym := range freeTable {
		c.loadSymbol(sym)
	}

	cf := &object.CompiledFunction{
		Instructions: insts,
		Constants:    consts,
		ParamsNum:    len(expr.Params),
		LocalsNum:    localNums,
		StackDepth:   maxStackDepth,
	}

	cfIdx := c.addConstant(cf)

	c.emit(Closure_, cfIdx, len(freeTable))

	if mSym.Name != "" {
		c.emit(Dup)
		c.storeSymbol(mSym)
	}
}
