package ir

import (
	"bytes"
	"strings"

	"zzc/fall-script/src/token"
)

type Pair struct {
	Key   Expr
	Value Expr
}

func (p *Pair) String() string {
	var buf bytes.Buffer
	buf.WriteString(p.Key.String())
	buf.WriteString(":")
	buf.WriteString(p.Value.String())
	return buf.String()
}

type HashLiteral struct {
	Token token.Token
	Pairs []*Pair
}

func (hl *HashLiteral) exprNode() {}

func (hl *HashLiteral) String() string {
	var buf bytes.Buffer

	buf.WriteString("{")
	paris := make([]string, len(hl.Pairs))
	for i, pari := range hl.Pairs {
		paris[i] = pari.String()
	}
	buf.WriteString(strings.Join(paris, ", "))
	buf.WriteString("}")

	return buf.String()
}
