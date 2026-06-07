package ast

import (
	"bytes"

	"zzc/fall-script/src/token"
)

type PrefixExpr struct {
	Token token.Token
	Op    string
	Right ExprNode
}

func (pe *PrefixExpr) ExprNode() {}

func (pe *PrefixExpr) GetToken() token.Token {
	return pe.Token
}

func (pe *PrefixExpr) TokenValue() string {
	return pe.Op
}

func (p *PrefixExpr) String() string {
	var buf bytes.Buffer

	buf.WriteString("(")
	buf.WriteString(p.Op)
	buf.WriteString(p.Right.String())
	buf.WriteString(")")

	return buf.String()
}
