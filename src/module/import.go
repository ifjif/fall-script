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

type ImportMetas []*ImportMeta

func ResolveReExport(l *Loader, source string, ep *ExportMeta, er ExportMetasRegister) *ExportMeta {
	// 解决re-export
	if ep.Origin == OriginReExport && ep.Ast == nil {
		reSource := ResolveImportPath(source, ep.Source)
		reps, ok := er[reSource]
		if !ok {
			reps = resolveExports(l, reSource, er)
		}

		nep := LookupExport(reps, ep.Imported, reSource)
		ResolveReExport(l, reSource, nep, er)
		ep.Ast = nep.Ast
		return nep
	}

	return ep
}

func LookupExport(eps ExportMetas, name, source string) *ExportMeta {
	ep, ok := eps[name]
	if !ok {
		msg := fmt.Sprintf("no exported named %q in file %s", name, source)
		panic(msg)
	}

	return ep
}

func ResolveImports(l *Loader, file string, imports []*ast.ImportStmt, er ExportMetasRegister) ImportMetas {
	import2s := ImportMetas{}

	for i, imp := range imports {
		source := ResolveImportPath(file, imp.Source)
		eps := resolveExports(l, source, er)
		for k, spe := range imp.Specifiers {
			ep := LookupExport(eps, spe.Imported, source)
			ResolveReExport(l, source, ep, er)
			ip := NewImport2(spe.Local, ep.Ast, i, k)
			import2s = append(import2s, ip)
		}
	}

	//	fmt.Println("Imports: ")
	//	for _, ip := range import2s {
	//		fmt.Printf("  name: %s\n", ip.Name)
	//		fmt.Printf("  import_idx: %d\n", ip.ImportIdx)
	//		fmt.Printf("  name_idx: %d\n", ip.NameIdx)
	//		fmt.Printf("  node: %q\n", ip.Ast.String())
	//		fmt.Println("")
	//	}
	//
	//	fmt.Println()
	//
	//	for name, ep := range er {
	//		fmt.Println(name)
	//		fmt.Println("Exports:")
	//		for n, ee := range ep {
	//			fmt.Printf("  name: %s\n", n)
	//			fmt.Printf("  export_idx: %d\n", ee.ExportIdx)
	//			fmt.Printf("  origin: %d\n", ee.Origin)
	//			fmt.Printf("  source: %s\n", ee.Source)
	//			fmt.Printf("  imported: %s\n", ee.Imported)
	//			ast := ""
	//			if ee.Ast != nil {
	//				ast = ee.Ast.String()
	//			}
	//			fmt.Printf("  node: %s\n", ast)
	//			fmt.Println()
	//		}
	//		fmt.Println()
	//	}

	return import2s
}
