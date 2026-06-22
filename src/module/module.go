package module

import (
	"path/filepath"

	"zzc/fall-script/src/ast"
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
			nStmts = append(nStmts, declaration.(ast.StmtNode))
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
