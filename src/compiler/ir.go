package compiler

import (
	"zzc/fall-script/src/code"
	"zzc/fall-script/src/ir"
	"zzc/fall-script/src/object"
)

func (c *Compiler) compileStmt(stmt ir.Stmt) {
	switch stmt := stmt.(type) {
	case *ir.ImportStmt:
		c.compileImportStmt(stmt)
	case *ir.ExportStmt:
		c.compileExportStmt(stmt)
	case *ir.StructDeclStmt:
		c.compileStructDeclStmtIr(stmt)
	case *ir.LetStmt:
		c.compileLetStmtIr(stmt)
	case *ir.BlockStmt:
		c.compileBlockStmtIr(stmt)
	case *ir.ForStmt:
		c.compileForStmtIr(stmt)
	case *ir.WhileStmt:
		c.compileWhileStmtIr(stmt)
	case *ir.DoWhileStmt:
		c.compileDoWhileStmtIr(stmt)
	case *ir.ReturnStmt:
		c.compileReturnStmtIr(stmt)
	case *ir.ExprStmt:
		c.compileExprStmtIr(stmt)
	}
}

func (c *Compiler) compileImportStmt(stmt *ir.ImportStmt) {
	from := c.addStrConstant(stmt.From)
	local := c.addStrConstant(stmt.Local)
	imported := c.addStrConstant(stmt.Imported)

	c.imports = append(c.imports, &object.ImportRef{From: from, Local: local, Imported: imported})
}

func (c *Compiler) compileExportStmt(stmt *ir.ExportStmt) {
	name := c.addStrConstant(stmt.Name)
	c.exports = append(c.exports, &object.ExportRef{Name: name, GlobalId: stmt.GlobalId})
}

func (c *Compiler) compileStructDeclStmtIr(stmt *ir.StructDeclStmt) {
	c.emit(code.Const, c.addConstant(stmt.MetaData))
	c.storeSymbol(stmt.Name)
	c.structs = append(c.structs, stmt.MetaData)
}

func (c *Compiler) compileLetStmtIr(stmt *ir.LetStmt) {
	c.compileExpr(stmt.Value)
	c.storeSymbol(stmt.Name.Symbol)
}

func (c *Compiler) compileBlockStmtIr(stmt *ir.BlockStmt) {
	for _, st := range stmt.Stmts {
		c.compileStmt(st)
	}
}

func (c *Compiler) compileForStmtIr(stmt *ir.ForStmt) {
	if stmt.Start != nil {
		c.compileStmt(stmt.Start)
	}

	conditonPos := len(c.CurrentInstructions())
	if stmt.Condition != nil {
		c.compileExpr(stmt.Condition)
	} else {
		c.emit(code.True)
	}
	jif := c.emit(code.JumpIsFalse, 9999)

	c.compileBlockStmtIr(stmt.Body)

	if stmt.Update != nil {
		c.compileExpr(stmt.Update)
		c.emit(code.Pop)
	}
	c.emit(code.Jump, conditonPos)
	c.changeOperand(jif, len(c.CurrentInstructions()))
}

func (c *Compiler) compileWhileStmtIr(stmt *ir.WhileStmt) {
	conditonPos := len(c.CurrentInstructions())
	c.compileExpr(stmt.Condition)
	jif := c.emit(code.JumpIsFalse, 9999)
	c.compileBlockStmtIr(stmt.Body)
	c.emit(code.Jump, conditonPos)
	c.changeOperand(jif, len(c.CurrentInstructions()))
}

func (c *Compiler) compileDoWhileStmtIr(stmt *ir.DoWhileStmt) {
	blockPos := len(c.CurrentInstructions())
	c.compileBlockStmtIr(stmt.Body)
	c.compileExpr(stmt.Condition)
	jif := c.emit(code.JumpIsFalse, 9999)
	c.emit(code.Jump, blockPos)
	c.changeOperand(jif, len(c.CurrentInstructions()))
}

func (c *Compiler) compileReturnStmtIr(stmt *ir.ReturnStmt) {
	if stmt.Value == nil {
		c.emit(code.Return)
		return
	}

	c.compileExpr(stmt.Value)
	c.emit(code.XReturn)
}

func (c *Compiler) compileExprStmtIr(stmt *ir.ExprStmt) {
	c.compileExpr(stmt.Expr)
	c.emit(code.Pop)
}
