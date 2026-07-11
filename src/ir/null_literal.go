package ir

import "zzc/fall-script/src/token"

type NullLiteral struct {
	Token token.Token
}

func (nl *NullLiteral) exprNode() {}

func (nl *NullLiteral) String() string {
	return "null"
}
