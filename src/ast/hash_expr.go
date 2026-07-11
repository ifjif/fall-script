package ast

import (
	"bytes"
	"strings"

	"zzc/fall-script/src/token"
)

/*
*
let a : i8 =

	struct {
		abc
		a:i8

b:bool
}

	{
		a:12,
		b:23,
	}

*
*/

type Pair struct {
	Key   ExprNode
	Value ExprNode
}

type HashExpr struct {
	Token token.Token
	Pairs []*Pair
}

func (he *HashExpr) ExprNode() {}

func (he *HashExpr) GetToken() token.Token {
	return he.Token
}

func (he *HashExpr) TokenValue() string {
	return he.Token.Value
}

func (he *HashExpr) String() string {
	var buf bytes.Buffer

	buf.WriteString("{")

	pairs := make([]string, len(he.Pairs))
	for i, p := range he.Pairs {
		pairs[i] = p.Key.String() + ":" + p.Value.String()
	}

	buf.WriteString(strings.Join(pairs, ", "))
	buf.WriteString("}")

	return buf.String()
}
