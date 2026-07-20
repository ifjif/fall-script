package ast

import "zzc/fall-script/src/token"

// int,string, bool, null
type LiteralPattern struct {
	Token token.Token
	Value ExprNode
}

func (lp *LiteralPattern) ExprNode()    {}
func (lp *LiteralPattern) patternNode() {}

func (lp *LiteralPattern) TokenValue() string {
	return lp.Token.Value
}

func (lp *LiteralPattern) GetToken() token.Token {
	return lp.Token
}

func (lp *LiteralPattern) String() string {
	return lp.Value.String()
}
