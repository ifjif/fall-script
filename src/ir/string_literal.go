package ir

import "zzc/fall-script/src/token"

type StringLiteral struct {
	Token token.Token
	Value string
}

func (sl *StringLiteral) exprNode() {}

func (sl *StringLiteral) String() string {
	return sl.Value
}
