package macro

import (
	"zzc/fall-script/src/ast"
	"zzc/fall-script/src/object"
)

/*
* 处理imports
* 根据 imports中信息，先去目标exports找
* 如果是标识符，先从ast中找，没有找到，再从imports中找，如果都没有，报错
*
*
 */
//func resolveImports(imports []*object.ImportRef) {
//	for _, imp := range imports {
//		// 加载源
//		// source := imp.From
//		//	for _, exp := range exports {
//		//		// 非标识符和标识符进行拆分
//		//	}
//		// 在ast中找标识符的节点
//		// 还有剩余的，从imports中找
//	}
//}

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
