package ir

import (
	"bytes"

	"zzc/fall-script/src/token"
)

type MatchArmExpr struct {
	Token   token.Token
	Pattern PatternNode
	Guard   Expr
	Body    Stmt
}

func (ma *MatchArmExpr) exprNode() {}

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
