package ir

import "zzc/fall-script/src/token"

type PrefixExpr struct {
	Token token.Token
	Op    string
	Right Expr
}

func (pe *PrefixExpr) exprNode() {}

func (pe *PrefixExpr) String() string {
	return pe.Op + pe.Right.String()
}
