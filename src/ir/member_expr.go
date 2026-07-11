package ir

import (
	"bytes"

	"zzc/fall-script/src/token"
)

type MemberExpr struct {
	Token   token.Token
	Visitor Expr
	Member  Expr
}

func (me *MemberExpr) exprNode() {}

func (me *MemberExpr) String() string {
	var buf bytes.Buffer

	buf.WriteString(me.Visitor.String())
	buf.WriteString(".")
	buf.WriteString(me.Member.String())

	return buf.String()
}
