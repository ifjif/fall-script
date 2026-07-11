package semantic

import (
	"fmt"

	"zzc/fall-script/src/ast"
	"zzc/fall-script/src/ir"
)

func (a *Analyzer) analyzeImports(imports []*ast.ImportStmt) []ir.Stmt {
	ips := []ir.Stmt{}
	for _, imp := range imports {
		source := imp.Source
		for _, specifier := range imp.Specifiers {
			local := specifier.Local
			imported := specifier.Imported
			sym := a.SymbolTable.Define(local)

			importRef := &ir.ImportStmt{
				GlobalId: sym.Pos,
				From:     source,
				Imported: imported,
				Local:    local,
			}

			ips = append(ips, importRef)
		}
	}

	return ips
}

func (a *Analyzer) analyzeExports(exports []*ast.ExportStmt) []ir.Stmt {
	exps := []ir.Stmt{}
	for _, exp := range exports {
		node := exp.Declaration
		name := ""

		switch node := node.(type) {
		case *ast.LetStmt:
			name = node.Name.Value
		case *ast.FnExpr:
			if !node.UnName {
				name = node.Ident.Value
			}
		case *ast.IdentExpr:
			name = node.Value
		}

		if name != "" {
			idx := a.resolveExports(name)
			exp := a.defineExport(name, idx)
			exps = append(exps, exp)
		}
	}

	return exps
}

func (a *Analyzer) resolveExports(name string) int {
	symbol, ok := a.SymbolTable.Resolve(name)
	if !ok {
		panic(fmt.Sprintf("Error: export name '%s' not found", name))
	}
	if symbol.Scope != ir.GLOBAL {
		panic(fmt.Sprintf("Error: export name '%s' is not global variable", name))
	}

	return symbol.Pos
}

func (c *Analyzer) defineExport(name string, globalId int) *ir.ExportStmt {
	exp := &ir.ExportStmt{Name: name, GlobalId: globalId}
	return exp
}
