package ast

import "zzc/fall-script/src/token"

type NullExpr struct {
	Token token.Token
}

func (ne *NullExpr) ExprNode() {}

func (ne *NullExpr) TokenValue() string {
	return ne.Token.Value
}

func (ne *NullExpr) String() string {
	return "null"
}

func (ne *NullExpr) GetToken() token.Token {
	return ne.Token
}
