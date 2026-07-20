package ir

import "zzc/fall-script/src/token"

type LiteralPattern struct {
	Token token.Token
	Value Expr
}

func (lp *LiteralPattern) exprNode()    {}
func (lp *LiteralPattern) patternNode() {}

func (lp *LiteralPattern) String() string {
	return lp.Value.String()
}
