package object

import (
	"bytes"
	"strings"

	"zzc/fall-script/src/ast"
)

type Macro struct {
	Name   string
	Params []*ast.IdentExpr
	Body   *ast.BlockStmt
	Kind   string // call attr
	Env    *Environment
}

func (m *Macro) Type() ObjectType {
	return MACRO_OBJ
}

func (m *Macro) Inspect() string {
	var buf bytes.Buffer

	buf.WriteString("#[macro(")
	buf.WriteString(m.Kind)
	buf.WriteString(")]\n")
	buf.WriteString("fn ")
	buf.WriteString(m.Name)
	buf.WriteString("(")

	if len(m.Params) > 0 {
		params := make([]string, len(m.Params))
		for i, param := range m.Params {
			params[i] = param.String()
		}
		buf.WriteString(strings.Join(params, ", "))
	}
	buf.WriteString(")")
	buf.WriteString(m.Body.String())

	return buf.String()
}
