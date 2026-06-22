package utils

import (
	"fmt"
	"strings"

	"zzc/fall-script/src/object"
)

func PrintModule(module *object.Module, indent string) {
	fmt.Printf("%sname: %s\n", indent, module.Name)
	childIndent := indent + "  "
	fmt.Printf("%sglobals     	: %d\n", childIndent, module.GlobalNum)
	fmt.Printf("%simports:\n", childIndent)
	imports := make([]string, len(module.Imports))
	importAndExportIndent := childIndent + "  "
	for i, imp := range module.Imports {
		from := imp.From
		imported := imp.Imported
		local := imp.Local
		imports[i] = fmt.Sprintf("%s%04d from: %d imported: %d local: %d", importAndExportIndent, i, from, imported, local)
	}
	fmt.Println(strings.Join(imports, "\n"))

	fmt.Printf("%sexports:\n", childIndent)
	exports := make([]string, len(module.Exports))
	for i, exp := range module.Exports {
		name := exp.Name
		globalId := exp.GlobalId
		exports[i] = fmt.Sprintf("%s%04d name: %d globalId: %d", importAndExportIndent, i, name, globalId)
	}
	fmt.Println(strings.Join(exports, "\n"))
	printCompiledFunction(module.Cf, childIndent)
}

func printCompiledFunction(cf *object.CompiledFunction, indent string) {
	fmt.Printf("%slocals     	: %d\n", indent, cf.LocalsNum)
	fmt.Printf("%sparams     	: %d\n", indent, cf.ParamsNum)
	fmt.Printf("%sstack depth	: %d\n", indent, cf.StackDepth)
	fmt.Printf("%sconstants:\n", indent)
	constAndInstIndent := indent + "  "
	printConstants(cf.Constants, constAndInstIndent)
	fmt.Printf("%sinstructions:\n", indent)
	fmt.Println(cf.Instructions.StringWithIndent(constAndInstIndent))
}

func printConstants(consts []object.Object, indent string) {
	childIndent := indent + "      "
	for i, ct := range consts {
		fmt.Printf("%s%04d: %s\n", indent, i, ct.Type())
		switch ct := ct.(type) {
		case *object.CompiledFunction:
			printCompiledFunction(ct, childIndent)
		default:
			fmt.Printf("%s%s\n", childIndent, ct.Inspect())
		}
	}
}
