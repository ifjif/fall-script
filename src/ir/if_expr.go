package ir

import (
	"bytes"

	"zzc/fall-script/src/token"
)

type IfExpr struct {
	Token       token.Token
	Condition   Expr
	Consequence *BlockStmt
	Alternative *BlockStmt
}

func (ie *IfExpr) exprNode() {}

func (ie *IfExpr) String() string {
	var buf bytes.Buffer
	buf.WriteString("if(")
	buf.WriteString(ie.Condition.String())
	buf.WriteString(")")
	buf.WriteString(ie.Consequence.String())
	if ie.Alternative != nil {
		buf.WriteString("else")
		buf.WriteString(ie.Alternative.String())
	}

	return buf.String()
}
