package ir

import (
	"zzc/fall-script/src/token"
)

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) exprNode() {}

func (il *IntegerLiteral) String() string {
	return il.Token.Value
}
