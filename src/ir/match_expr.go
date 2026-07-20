package ir

import (
	"bytes"

	"zzc/fall-script/src/token"
)

type MatchExpr struct {
	Token     token.Token
	Subject   Expr
	MatchArms []*MatchArmExpr
}

func (me *MatchExpr) exprNode() {}

func (me *MatchExpr) String() string {
	var buf bytes.Buffer

	buf.WriteString("match ")
	buf.WriteString(me.Subject.String())
	buf.WriteString("{\n")

	for _, arm := range me.MatchArms {
		buf.WriteString(arm.String())
		buf.WriteString("\n")
	}

	buf.WriteString("}")

	return buf.String()
}
