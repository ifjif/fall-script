package ast

import (
	"bytes"

	"zzc/fall-script/src/token"
)

// pattern if guard => xxx
type MatchArmExpr struct {
	Token   token.Token
	Pattern PatternNode
	Guard   ExprNode
	Body    StmtNode
}

func (ma *MatchArmExpr) ExprNode() {}

func (ma *MatchArmExpr) TokenValue() string {
	return ma.Token.Value
}

func (ma *MatchArmExpr) GetToken() token.Token {
	return ma.Token
}

func (ma *MatchArmExpr) String() string {
	var buf bytes.Buffer

	buf.WriteString(ma.Pattern.String())

	if ma.Guard != nil {
		buf.WriteString(" if ")
		buf.WriteString(ma.Guard.String())
	}
	buf.WriteString(" => ")

	buf.WriteString(ma.Body.String())

	return buf.String()
}
