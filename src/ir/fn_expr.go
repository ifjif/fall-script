package ir

import (
	"bytes"
	"fmt"
	"strings"

	"zzc/fall-script/src/token"
)

type FnExpr struct {
	Token     token.Token
	Name      *Ident
	Params    []*Ident
	Body      *BlockStmt
	UnName    bool
	LocalVars int
	Frees     []*Symbol
}

func (fe *FnExpr) exprNode() {}

func (fe *FnExpr) String() string {
	var buf bytes.Buffer

	buf.WriteString("(")
	buf.WriteString(fmt.Sprintf("Locals:%d, frees:%d", fe.LocalVars, len(fe.Frees)))
	buf.WriteString(")")
	for _, free := range fe.Frees {
		msg := fmt.Sprintf("free:(name: %s pos: %d scope: %s captured: %t)\n", free.Name, free.Pos, free.Scope, free.Captured)
		buf.WriteString(msg)
	}
	buf.WriteString("fn ")
	if !fe.UnName {
		buf.WriteString(fe.Name.String())
	}
	params := make([]string, len(fe.Params))
	for i, param := range fe.Params {
		params[i] = param.String()
	}
	buf.WriteString("(")
	buf.WriteString(strings.Join(params, ", "))
	buf.WriteString(")")
	buf.WriteString(fe.Body.String())

	return buf.String()
}
