package ast

import (
	"bytes"
	"strings"

	"zzc/fall-script/src/token"
)

type StructPair struct {
	Key   *IdentExpr
	Value ExprNode
}
type StructLiteralExpr struct {
	Token    token.Token
	Tag      *IdentExpr
	Elements []*StructPair
}

func (sl *StructLiteralExpr) ExprNode() {}

func (sl *StructLiteralExpr) TokenValue() string {
	return sl.Token.Value
}

func (sl *StructLiteralExpr) GetToken() token.Token {
	return sl.Token
}

func (sl *StructLiteralExpr) String() string {
	var buf bytes.Buffer

	buf.WriteString(sl.Tag.Value)
	buf.WriteString("{\n")

	elements := make([]string, len(sl.Elements))

	for i, elem := range sl.Elements {
		key := elem.Key.String()
		value := elem.Value.String()

		pair := key + ": " + value
		elements[i] = pair
	}

	buf.WriteString(strings.Join(elements, ",\n"))
	buf.WriteString("\n}")

	return buf.String()
}
