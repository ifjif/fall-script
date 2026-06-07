package ast

import "zzc/fall-script/src/token"

type StrExpr struct {
	Token token.Token
	Value string
}

func (se *StrExpr) ExprNode() {}

func (se *StrExpr) String() string {
	return se.Token.Value
}

func (se *StrExpr) TokenValue() string {
	return se.Token.Value
}

func (se *StrExpr) GetToken() token.Token {
	return se.Token
}
