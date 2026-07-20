package ast

import "zzc/fall-script/src/token"

type WildcardPattern struct {
	Token token.Token
}

func (wp *WildcardPattern) ExprNode()    {}
func (wp *WildcardPattern) patternNode() {}

func (wp *WildcardPattern) TokenValue() string {
	return wp.Token.Value
}

func (wp *WildcardPattern) GetToken() token.Token {
	return wp.Token
}

func (wp *WildcardPattern) String() string {
	return wp.Token.Value
}
