package ast

import (
	"bytes"
	"strings"

	"zzc/fall-script/src/token"
)

type ArrExpr struct {
	Token    token.Token
	Elements []ExprNode
}

func (ae *ArrExpr) ExprNode() {}

func (ae *ArrExpr) GetToken() token.Token {
	return ae.Token
}

func (ae *ArrExpr) TokenValue() string {
	return ae.Token.Value
}

func (ae *ArrExpr) String() string {
	var buf bytes.Buffer

	buf.WriteString("[")

	elements := make([]string, len(ae.Elements))
	for i, e := range ae.Elements {
		elements[i] = e.String()
	}

	buf.WriteString(strings.Join(elements, ", "))
	buf.WriteString("]")

	return buf.String()
}
