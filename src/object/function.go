package object

import (
	"bytes"
	"strings"

	"zzc/fall-script/src/ast"
)

type Function struct {
	Name   string
	Params []*ast.IdentExpr
	Body   *ast.BlockStmt
	Env    *Environment
}

func (f *Function) Type() ObjectType {
	return FUNCTION_OBJ
}

func (f *Function) Inspect() string {
	var buf bytes.Buffer

	buf.WriteString("fn ")
	buf.WriteString(f.Name)

	buf.WriteString("(")
	params := make([]string, len(f.Params))
	for i, param := range f.Params {
		params[i] = param.String()
	}
	buf.WriteString(strings.Join(params, ", "))
	buf.WriteString(")")

	buf.WriteString(f.Body.String())

	return buf.String()
}
