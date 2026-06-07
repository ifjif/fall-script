package ast

import "zzc/fall-script/src/token"

type BoolExpr struct {
	Token token.Token
	Value bool
}

func (be *BoolExpr) ExprNode() {}

func (be *BoolExpr) String() string {
	return be.Token.Value
}

func (be *BoolExpr) TokenValue() string {
	return be.Token.Value
}

func (be *BoolExpr) GetToken() token.Token {
	return be.Token
}
