package ir

import (
	"bytes"
	"strings"

	"zzc/fall-script/src/token"
)

type StructLiteral struct {
	Token token.Token
	Tag   *Ident
	Pairs []*Pair
}

func (sl *StructLiteral) exprNode() {}

func (sl *StructLiteral) String() string {
	var buf bytes.Buffer

	buf.WriteString(sl.Tag.String())
	buf.WriteString("{\n")
	pairs := make([]string, len(sl.Pairs))
	for i, pair := range sl.Pairs {
		pairs[i] = pair.String()
	}
	buf.WriteString(strings.Join(pairs, ",\n"))
	buf.WriteString("}")

	return buf.String()
}
