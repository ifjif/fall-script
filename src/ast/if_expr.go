package ast

import (
	"bytes"

	"zzc/fall-script/src/token"
)

type IfExpr struct {
	Token       token.Token
	Condition   ExprNode
	Consequence *BlockStmt
	Alternative *BlockStmt
}

func (ie *IfExpr) ExprNode() {}

func (ie *IfExpr) GetToken() token.Token {
	return ie.Token
}

func (ie *IfExpr) TokenValue() string {
	return ie.Token.Value
}

func (ie *IfExpr) String() string {
	var buf bytes.Buffer

	buf.WriteString("if")
	buf.WriteString(ie.Condition.String())
	buf.WriteString(ie.Consequence.String())
	if ie.Alternative != nil {
		buf.WriteString("else")
		buf.WriteString(ie.Alternative.String())
	}

	return buf.String()
}
