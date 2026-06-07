package ast

import "zzc/fall-script/src/token"

type IdentExpr struct {
	Token token.Token
	Value string
}

func (ie *IdentExpr) ExprNode() {}

func (ie *IdentExpr) GetToken() token.Token {
	return ie.Token
}

func (ie *IdentExpr) TokenValue() string {
	return ie.Value
}

func (ie *IdentExpr) String() string {
	return ie.Value
}
