package macro

import (
	"fmt"
	"sort"

	"zzc/fall-script/src/ast"
	"zzc/fall-script/src/object"
)

// import宏，export宏，都需要从import和export中删除
func DefineMacros(program *ast.Program, env *object.Environment) {
	// 从import中找,然后删除
	// 从export中找，然后删除
	defines := []int{}

	for i, stmt := range program.Stmts {
		if ok, attr := IsMacroDefinition(stmt); ok {
			AddMacro(stmt, attr, env)
			defines = append(defines, i)
		}
	}

	deleteMacroDefinition(defines, program)
}

func IsMacroDefinition(stmt ast.Node) (bool, *ast.AttributeExpr) {
	exprStmt, ok := stmt.(*ast.ExprStmt)
	var expr ast.ExprNode
	if !ok {
		expr, ok = stmt.(ast.ExprNode)
		if !ok {
			return false, nil
		}
	} else {
		expr = exprStmt.Expr
	}

	fn, ok := expr.(*ast.FnExpr)
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

func AddMacroWithAlias(alias string, stmt ast.Node, attr *ast.AttributeExpr, env *object.Environment) {
	macro := getMacro(stmt, attr, env)
	macro.Name = alias
	env.Set(macro.Name, macro)
}

func AddMacro(stmt ast.Node, attr *ast.AttributeExpr, env *object.Environment) {
	macro := getMacro(stmt, attr, env)
	env.Set(macro.Name, macro)
}

func getMacro(stmt ast.Node, attr *ast.AttributeExpr, env *object.Environment) *object.Macro {
	var expr ast.Node
	if es, ok := stmt.(*ast.ExprStmt); ok {
		expr = es.Expr
	} else {
		expr = stmt
	}
	fn := expr.(*ast.FnExpr)

	macro := &object.Macro{
		Name:   fn.Name,
		Kind:   attr.Args[0].String(),
		Params: fn.Params,
		Body:   fn.Body,
		Env:    env,
	}

	return macro
}

func deleteMacroDefinition(defines []int, program *ast.Program) {
	for i := len(defines) - 1; i >= 0; i-- {
		idx := defines[i]
		program.Stmts = append(program.Stmts[:idx], program.Stmts[idx+1:]...)
	}
}

func DeleteMacroFromImports(defines map[int][]int, imports []*ast.ImportStmt) []*ast.ImportStmt {
	fmt.Println("删除 import中的 宏")
	rootDefines := make([]int, 0)
	for i, def := range defines {
		DeleteMacroFromImport(def, imports[i])
		rootDefines = append(rootDefines, i)
	}

	sort.Ints(rootDefines)

	for i := len(rootDefines) - 1; i >= 0; i-- {
		idx := rootDefines[i]
		if len(imports[idx].Specifiers) == 0 {
			imports = append(imports[:idx], imports[idx+1:]...)
		}
	}

	return imports
}

func DeleteMacroFromImport(defines []int, imports *ast.ImportStmt) {
	for i := len(defines) - 1; i >= 0; i-- {
		idx := defines[i]
		imports.Specifiers = append(imports.Specifiers[:idx], imports.Specifiers[idx+1:]...)
	}
}

func DeleteMacroFromExports(defines []int, exports []*ast.ExportStmt) []*ast.ExportStmt {
	fmt.Println("删除 export中的 宏")
	for i := len(defines) - 1; i >= 0; i-- {
		idx := defines[i]
		exports = append(exports[:idx], exports[idx+1:]...)
	}
	return exports
}
