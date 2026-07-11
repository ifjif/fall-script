package ir

import (
	"zzc/fall-script/src/token"
)

type BoolLiteral struct {
	Token token.Token
	Value bool
}

func (bl *BoolLiteral) exprNode() {}

func (bl *BoolLiteral) String() string {
	return bl.Token.Value
}
