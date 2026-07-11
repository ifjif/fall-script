package ir

import (
	"bytes"

	"zzc/fall-script/src/token"
)

type InfixExpr struct {
	Token token.Token
	Left  Expr
	Op    string
	Right Expr
}

func (ie *InfixExpr) exprNode() {}

func (ie *InfixExpr) String() string {
	var buf bytes.Buffer

	buf.WriteString(ie.Left.String())
	buf.WriteString(" ")
	buf.WriteString(ie.Op)
	buf.WriteString(" ")
	buf.WriteString(ie.Right.String())

	return buf.String()
}
