package ast

import (
	"zzc/fall-script/src/token"
)

type IntExpr struct {
	Token token.Token
	Value int64
}

func (ie *IntExpr) ExprNode() {}

func (ie *IntExpr) GetToken() token.Token {
	return ie.Token
}

func (ie *IntExpr) String() string {
	return ie.Token.Value
}

func (ie *IntExpr) TokenValue() string {
	return ie.Token.Value
}
