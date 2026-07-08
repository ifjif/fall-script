package ast

import "zzc/fall-script/src/token"

type FieldDeclExpr struct {
	Token   token.Token
	Name    *IdentExpr
	IsEmbed bool
}

func (fd *FieldDeclExpr) ExprNode() {}

func (fd *FieldDeclExpr) TokenValue() string {
	return fd.Token.Value
}

func (fd *FieldDeclExpr) GetToken() token.Token {
	return fd.Token
}

func (fd *FieldDeclExpr) String() string {
	str := fd.Name.String()
	if !fd.IsEmbed {
		str += ":"
	}

	return str
}
