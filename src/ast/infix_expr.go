package ast

import (
	"bytes"

	"zzc/fall-script/src/token"
)

type InfixExpr struct {
	Token token.Token
	Left  ExprNode
	Op    string
	Right ExprNode
}

func (ie *InfixExpr) ExprNode() {}

func (ie *InfixExpr) GetToken() token.Token {
	return ie.Token
}

func (ie *InfixExpr) TokenValue() string {
	return ie.Op
}

func (ie *InfixExpr) String() string {
	var buf bytes.Buffer

	buf.WriteString("(")
	buf.WriteString(ie.Left.String())
	buf.WriteString(" " + ie.Op + " ")
	buf.WriteString(ie.Right.String())
	buf.WriteString(")")

	return buf.String()
}
