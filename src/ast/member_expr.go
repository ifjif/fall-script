package ast

import "zzc/fall-script/src/token"

type MemberExpr struct {
	Token   token.Token
	Visitor ExprNode
	Member  ExprNode
}

func (me MemberExpr) ExprNode() {}

func (me MemberExpr) TokenValue() string {
	return me.Token.Value
}

func (me MemberExpr) GetToken() token.Token {
	return me.Token
}

func (me MemberExpr) String() string {
	return me.Visitor.String() + "." + me.Member.String()
}
