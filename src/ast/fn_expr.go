package ast

import (
	"bytes"
	"strings"

	"zzc/fall-script/src/token"
)

type FnExpr struct {
	Token  token.Token
	Name   string
	Params []*IdentExpr
	Body   *BlockStmt
	UnName bool
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

	buf.WriteString("fn ")
	buf.WriteString(fe.Name)
	buf.WriteString("(")
	params := make([]string, len(fe.Params))
	for i, param := range fe.Params {
		params[i] = param.String()
	}
	buf.WriteString(strings.Join(params, ", "))
	buf.WriteString(")")
	buf.WriteString(fe.Body.String())

	return buf.String()
}
