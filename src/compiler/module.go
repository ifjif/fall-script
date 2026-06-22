package compiler

import (
	"fmt"

	"zzc/fall-script/src/ast"
	"zzc/fall-script/src/object"
)

func (c *Compiler) defineImports(imports []*ast.ImportStmt) {
	ips := []*object.ImportRef{}
	for _, imp := range imports {
		sourceId := c.addStrConstant(imp.Source)
		for _, specifier := range imp.Specifiers {
			local := specifier.Local
			imported := specifier.Imported
			c.SymbolTable.Define(local)
			localId := 0
			importId := 0
			if local == imported {
				idx := c.addStrConstant(local)
				localId, importId = idx, idx
			} else {
				localId = c.addStrConstant(local)
				importId = c.addStrConstant(imported)
			}

			importRef := &object.ImportRef{
				From:     sourceId,
				Imported: importId,
				Local:    localId,
			}

			ips = append(ips, importRef)
		}
	}
	c.imports = ips
}

func (c *Compiler) defineExports(exports []*ast.ExportStmt) {
	exps := []*object.ExportRef{}
	expNames := []string{}
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
			exp := c.defineExport(name)
			exps = append(exps, exp)
			expNames = append(expNames, name)
		}
	}

	c.exports = exps
	c.exportNames = expNames
}

func (c *Compiler) resolveExports() {
	for i, name := range c.exportNames {
		symbol, ok := c.SymbolTable.Resolve(name)
		if !ok {
			panic(fmt.Sprintf("Error: export name '%s' not found", name))
		}
		if symbol.Scope != GLOBAL {
			panic(fmt.Sprintf("Error: export name '%s' is not global variable", name))
		}
		c.exports[i].GlobalId = symbol.Pos
	}
}
