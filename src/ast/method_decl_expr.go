package ast

import (
	"bytes"

	"zzc/fall-script/src/token"
)

type MethodDeclExpr struct {
	Token      token.Token
	StructName *IdentExpr
	Fn         *FnExpr
}

func (md *MethodDeclExpr) ExprNode() {}

func (md *MethodDeclExpr) GetToken() token.Token {
	return md.Token
}

func (md *MethodDeclExpr) TokenValue() string {
	return md.Token.Value
}

func (md *MethodDeclExpr) String() string {
	var buf bytes.Buffer
	buf.WriteString(md.StructName.String())
	buf.WriteString(": ")
	buf.WriteString(md.Fn.String())
	return buf.String()
}
