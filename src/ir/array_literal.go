package ir

import (
	"bytes"
	"strings"

	"zzc/fall-script/src/token"
)

type ArrayLiteral struct {
	Token    token.Token
	Elements []Expr
}

func (al *ArrayLiteral) exprNode() {}

func (al *ArrayLiteral) String() string {
	var buf bytes.Buffer

	buf.WriteString("[")
	elems := make([]string, len(al.Elements))

	for i, elem := range al.Elements {
		elems[i] = elem.String()
	}

	buf.WriteString(strings.Join(elems, ", "))
	buf.WriteString("]")

	return buf.String()
}
