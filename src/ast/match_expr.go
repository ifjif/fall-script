package ast

import (
	"bytes"

	"zzc/fall-script/src/token"
)

type MatchExpr struct {
	Token   token.Token
	Subject ExprNode
	Arms    []*MatchArmExpr
}

func (me *MatchExpr) ExprNode() {}

func (me *MatchExpr) TokenValue() string {
	return me.Token.Value
}

func (me *MatchExpr) GetToken() token.Token {
	return me.Token
}

func (me *MatchExpr) String() string {
	var buf bytes.Buffer

	buf.WriteString("match ")
	buf.WriteString(me.Subject.String())
	buf.WriteString("{\n")

	for _, arm := range me.Arms {
		buf.WriteString(arm.String())
		buf.WriteString("\n")
	}

	buf.WriteString("}")

	return buf.String()
}
