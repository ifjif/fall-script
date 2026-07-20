package ir

import "zzc/fall-script/src/token"

type WildcardPattern struct {
	Token token.Token
}

func (wp *WildcardPattern) exprNode()    {}
func (wp *WildcardPattern) patternNode() {}

func (wp *WildcardPattern) String() string {
	return wp.Token.Value
}
