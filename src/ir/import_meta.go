package ir

import "zzc/fall-script/src/ast"

type ImportMeta struct {
	Ast        ast.Node
	FieldTotal int      // struct 特有
	Methods    []string // struct 特有
	Fields     []int    // struct 特有
}
