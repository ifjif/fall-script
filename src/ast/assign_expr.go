package ast

import (
	"bytes"

	"zzc/fall-script/src/token"
)

type AssignExpr struct {
	Token token.Token
	Name  *IdentExpr
	Value ExprNode
}

func (ae *AssignExpr) ExprNode() {}

func (ae *AssignExpr) TokenValue() string {
	return ae.Token.Value
}

func (ae *AssignExpr) GetToken() token.Token {
	return ae.Token
}

func (ae *AssignExpr) String() string {
	var buf bytes.Buffer

	buf.WriteString(ae.Name.String())
	buf.WriteString(" = ")
	buf.WriteString(ae.Value.String())

	return buf.String()
}
