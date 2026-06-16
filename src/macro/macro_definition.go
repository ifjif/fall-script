package macro

import (
	"zzc/fall-script/src/ast"
	"zzc/fall-script/src/object"
)

func DefineMacros(program *ast.Program, env *object.Environment) {
	defines := []int{}

	for i, stmt := range program.Stmts {
		if ok, attr := isMacroDefinition(stmt); ok {
			addMacro(stmt, attr, env)
			defines = append(defines, i)
		}
	}

	deleteMacroDefinition(defines, program)
}

func isMacroDefinition(stmt ast.StmtNode) (bool, *ast.AttributeExpr) {
	exprStmt, ok := stmt.(*ast.ExprStmt)
	if !ok {
		return false, nil
	}

	fn, ok := exprStmt.Expr.(*ast.FnExpr)
	if !ok {
		return false, nil
	}

	for _, attr := range fn.Attrs {
		if attr.Name == "macro" {
			return true, attr
		}
	}

	return false, nil
}

func addMacro(stmt ast.StmtNode, attr *ast.AttributeExpr, env *object.Environment) {
	fn := stmt.(*ast.ExprStmt).Expr.(*ast.FnExpr)

	macro := &object.Macro{
		Name:   fn.Name,
		Kind:   attr.Args[0].String(),
		Params: fn.Params,
		Body:   fn.Body,
		Env:    env,
	}

	env.Set(macro.Name, macro)
}

func deleteMacroDefinition(defines []int, program *ast.Program) {
	for i := len(defines) - 1; i >= 0; i-- {
		idx := defines[i]
		program.Stmts = append(program.Stmts[:idx], program.Stmts[idx+1:]...)
	}
}
