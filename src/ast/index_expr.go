package ast

import (
	"bytes"

	"zzc/fall-script/src/token"
)

type IndexExpr struct {
	Token token.Token
	Left  ExprNode
	Index ExprNode
}

func (ie *IndexExpr) ExprNode() {}

func (ie *IndexExpr) GetToken() token.Token {
	return ie.Token
}

func (ie *IndexExpr) TokenValue() string {
	return ie.Token.Value
}

func (ie *IndexExpr) String() string {
	var buf bytes.Buffer

	buf.WriteString(ie.Left.String())
	buf.WriteString("[")
	buf.WriteString(ie.Index.String())
	buf.WriteString("]")

	return buf.String()
}
