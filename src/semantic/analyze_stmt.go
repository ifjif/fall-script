package semantic

import (
	"zzc/fall-script/src/ast"
	"zzc/fall-script/src/ir"
)

func (a *Analyzer) analyzeStmt(stmt ast.StmtNode) ir.Stmt {
	switch stmt := stmt.(type) {
	case *ast.LetStmt:
		return a.analyzeLetStmt(stmt)
	case *ast.BlockStmt:
		return a.analyzeBlockStmt(stmt)
	case *ast.ForStmt:
		return a.analyzeForStmt(stmt)
	case *ast.WhileStmt:
		return a.analyzeWhileStmt(stmt)
	case *ast.DoWhileStmt:
		return a.analyzeDoWhileStmt(stmt)
	case *ast.ReturnStmt:
		return a.analyzeReturnStmt(stmt)
	case *ast.ExprStmt:
		r := a.analyzeExprStmt(stmt)
		if r == nil {
			return nil
		}
		return r
	}

	return nil
}

func (a *Analyzer) analyzeLetStmt(stmt *ast.LetStmt) *ir.LetStmt {
	sym := a.SymbolTable.Define(stmt.Name.Value)
	value := a.analyzeExpr(stmt.Value)
	if value == nil {
		return nil
	}

	name := &ir.Ident{Token: stmt.Token, Symbol: sym}

	return &ir.LetStmt{Token: stmt.Token, Name: name, Value: value}
}

func (a *Analyzer) analyzeBlockStmt(stmt *ast.BlockStmt) *ir.BlockStmt {
	a.enterBlock()
	stmts := []ir.Stmt{}
	for _, stmt := range stmt.Stmts {
		istmt := a.analyzeStmt(stmt)
		if istmt != nil {
			stmts = append(stmts, istmt)
		}
	}
	a.leaveBlock()

	return &ir.BlockStmt{Token: stmt.Token, Stmts: stmts}
}

func (a *Analyzer) analyzeForStmt(stmt *ast.ForStmt) *ir.ForStmt {
	ifs := &ir.ForStmt{Token: stmt.Token}
	a.enterBlock()
	if stmt.Start != nil {
		ifs.Start = a.analyzeStmt(stmt.Start)
	}
	if stmt.Condition != nil {
		ifs.Condition = a.analyzeExpr(stmt.Condition)
	}
	if stmt.Update != nil {
		ifs.Update = a.analyzeExpr(stmt.Update)
	}
	ifs.Body = a.analyzeBlockStmt(stmt.Body)
	a.leaveBlock()
	return ifs
}

func (a *Analyzer) analyzeWhileStmt(stmt *ast.WhileStmt) *ir.WhileStmt {
	condition := a.analyzeExpr(stmt.Condition)
	body := a.analyzeBlockStmt(stmt.Body)

	return &ir.WhileStmt{Token: stmt.Token, Condition: condition, Body: body}
}

func (a *Analyzer) analyzeDoWhileStmt(stmt *ast.DoWhileStmt) *ir.DoWhileStmt {
	body := a.analyzeBlockStmt(stmt.Body)
	condition := a.analyzeExpr(stmt.Condition)

	return &ir.DoWhileStmt{Token: stmt.Token, Body: body, Condition: condition}
}

func (a *Analyzer) analyzeReturnStmt(stmt *ast.ReturnStmt) *ir.ReturnStmt {
	irs := &ir.ReturnStmt{Token: stmt.Token}
	if stmt.Value != nil {
		irs.Value = a.analyzeExpr(stmt.Value)
	}

	return irs
}

func (a *Analyzer) analyzeExprStmt(stmt *ast.ExprStmt) *ir.ExprStmt {
	istmt := &ir.ExprStmt{Token: stmt.Token}
	expr := a.analyzeExpr(stmt.Expr)

	if expr == nil {
		return nil
	}

	istmt.Expr = expr

	return istmt
}
