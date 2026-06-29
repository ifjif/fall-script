package ast

import (
	"bytes"
	"strings"

	"zzc/fall-script/src/token"
)

type FnExpr struct {
	Token  token.Token
	Name   string
	Ident  *IdentExpr
	Params []*IdentExpr
	Body   *BlockStmt
	UnName bool
	Attrs  []*AttributeExpr
}

func (fe *FnExpr) ExprNode() {}

func (fe *FnExpr) GetToken() token.Token {
	return fe.Token
}

func (fe *FnExpr) TokenValue() string {
	return fe.Token.Value
}

func (fe *FnExpr) String() string {
	var buf bytes.Buffer

	if len(fe.Attrs) > 0 {
		attrs := make([]string, len(fe.Attrs))
		for i, attr := range fe.Attrs {
			attrs[i] = attr.String()
		}
		buf.WriteString(strings.Join(attrs, "\n"))
		buf.WriteString("\n")
	}

	buf.WriteString("fn ")
	buf.WriteString(fe.Name)
	buf.WriteString("(")
	params := make([]string, len(fe.Params))
	for i, param := range fe.Params {
		params[i] = param.String()
	}
	buf.WriteString(strings.Join(params, ", "))
	buf.WriteString(")")
	if fe.Body != nil {
		buf.WriteString(fe.Body.String())
	}

	return buf.String()
}
