package ir

import (
	"bytes"

	"zzc/fall-script/src/token"
)

type AssignExpr struct {
	Token token.Token
	Left  Expr
	Value Expr
}

func (ae *AssignExpr) exprNode() {}

func (ae *AssignExpr) String() string {
	var buf bytes.Buffer

	buf.WriteString(ae.Left.String())
	buf.WriteString("=")
	buf.WriteString(ae.Value.String())

	return buf.String()
}
