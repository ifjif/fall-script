package module

import (
	"fmt"

	"zzc/fall-script/src/ast"
)

type ImportMeta struct {
	Name      string
	Ast       ast.Node
	ImportIdx int
	NameIdx   int
}

func NewImport2(name string, ast ast.Node, importIdx, nameIdx int) *ImportMeta {
	return &ImportMeta{
		Name:      name,
		Ast:       ast,
		ImportIdx: importIdx,
		NameIdx:   nameIdx,
	}
}

func ResolveImports(l *Loader, file string, imports []*ast.ImportStmt, er ExportMetasRegister) []*ImportMeta {
	import2s := []*ImportMeta{}

	for i, imp := range imports {
		source := ResolveImportPath(file, imp.Source)
		eps := resolveExports(l, source, er)
		for k, spe := range imp.Specifiers {
			ep, ok := eps[spe.Imported]
			if !ok {
				msg := fmt.Sprintf("no exported named %q in file %s", spe.Imported, source)
				panic(msg)
			}
			ip := NewImport2(spe.Local, ep.Ast, i, k)
			import2s = append(import2s, ip)
		}
	}

	fmt.Println("Imports: ")
	for _, ip := range import2s {
		fmt.Printf("  name: %s", ip.Name)
		fmt.Printf("  import_idx: %d", ip.ImportIdx)
		fmt.Printf("  name_idx: %d", ip.NameIdx)
		fmt.Printf("  node: %s\n", ip.Ast)
	}

	for name, ep := range er {
		fmt.Println(name)
		fmt.Println("Exports:")
		for n, ee := range ep {
			fmt.Printf("  name: %s", n)
			fmt.Printf("  export_idx: %d", ee.ExportIdx)
			fmt.Printf("  source: %s", ee.Source)
			fmt.Printf("  node: %s\n", ee.Ast.String())
		}
	}

	return import2s
}
