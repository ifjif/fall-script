package ir

import (
	"bytes"

	"zzc/fall-script/src/token"
)

type IndexExpr struct {
	Token token.Token
	Left  Expr
	Index Expr
}

func (ie *IndexExpr) exprNode() {}

func (ie *IndexExpr) String() string {
	var buf bytes.Buffer

	buf.WriteString(ie.Left.String())
	buf.WriteString("[")
	buf.WriteString(ie.Index.String())
	buf.WriteString("]")

	return buf.String()
}
