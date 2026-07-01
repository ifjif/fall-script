package module

import (
	"path/filepath"
	"sort"

	"zzc/fall-script/src/ast"
	"zzc/fall-script/src/macro"
	"zzc/fall-script/src/object"
)

// 模块名中不能有 . 和 ..
type ModuleRegister map[string]*object.CompiledModule

func CollectImportsAndExports(program *ast.Program) (imports []*ast.ImportStmt, exports []*ast.ExportStmt) {
	nStmts := make([]ast.StmtNode, 0)
	imports = make([]*ast.ImportStmt, 0)
	exports = make([]*ast.ExportStmt, 0)
	for _, stmt := range program.Stmts {
		switch stmt := stmt.(type) {
		case *ast.ImportStmt:
			imports = append(imports, stmt)
		case *ast.ExportStmt:
			declaration := stmt.Declaration
			if es, ok := declaration.(*ast.ExprStmt); ok {
				stmt.Declaration = es.Expr
			}
			exports = append(exports, stmt)
			// ident不需要保存到ast中
			if _, ok := stmt.Declaration.(*ast.IdentExpr); !ok {
				nStmts = append(nStmts, declaration.(ast.StmtNode))
			}
		default:
			nStmts = append(nStmts, stmt)
		}
	}

	program.Stmts = nStmts

	return imports, exports
}

func ResolveImportPath(currentFilePath, importPath string) string {
	dir := filepath.Dir(currentFilePath)

	resolvePath := filepath.Join(dir, importPath)

	return filepath.Clean(resolvePath)
}

func ResolveImportsAndExports(l *Loader, source string, program *ast.Program, imports []*ast.ImportStmt, exports []*ast.ExportStmt) (ImportMetas, ExportMetas) {
	er := ExportMetasRegister{}

	exp := resolveExports2(l, source, program, imports, exports, er)
	imp := ResolveImports(l, source, imports, er)

	return imp, exp
}

func ResolveMacrosFromProgram(l *Loader, p *ast.Program, file string, env *object.Environment) (program *ast.Program, imports []*ast.ImportStmt, exports []*ast.ExportStmt) {
	source := file
	program = p
	imports, exports = CollectImportsAndExports(program)

	imports2, exports2 := ResolveImportsAndExports(l, source, program, imports, exports)

	importsMacroDefine := map[int][]int{}
	for _, imp := range imports2 {
		node := imp.Ast
		if ok, attr := macro.IsMacroDefinition(node); ok {
			macro.AddMacroWithAlias(imp.Name, node, attr, env)
			importIdx := imp.ImportIdx
			importD, ok := importsMacroDefine[importIdx]
			if !ok {
				importD = make([]int, 0)
			}
			importD = append(importD, imp.NameIdx)
			importsMacroDefine[importIdx] = importD
		}
	}

	//	fmt.Println("原始import: ")
	//	for _, imp := range imports {
	//		fmt.Println(imp.String())
	//	}

	//	fmt.Println("要删除的import宏定义:")
	//	fmt.Println(importsMacroDefine)

	imports = macro.DeleteMacroFromImports(importsMacroDefine, imports)

	//	fmt.Println("原始exports: ")
	//	for _, exp := range exports {
	//		fmt.Println(exp.Declaration.String())
	//	}

	exportDefines := []int{}
	for _, exp := range exports2 {
		if ok, _ := macro.IsMacroDefinition(exp.Ast); ok {
			exportDefines = append(exportDefines, exp.ExportIdx)
		}
	}
	sort.Ints(exportDefines)

	//	fmt.Println("要删除的export宏定义: ")
	//	fmt.Println(exportDefines)
	exports = macro.DeleteMacroFromExports(exportDefines, exports)

	macro.DefineMacros(program, env)
	//	fmt.Println("找宏结束===============================")

	return program, imports, exports
}
