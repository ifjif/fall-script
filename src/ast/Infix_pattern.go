package ast

import (
	"fmt"

	"zzc/fall-script/src/token"
)

type InfixPattern struct {
	Token token.Token
	Left  PatternNode
	Op    string
	Right PatternNode
}

func (ip *InfixPattern) ExprNode()    {}
func (ip *InfixPattern) patternNode() {}

func (ip *InfixPattern) TokenValue() string {
	return ip.Token.Value
}

func (ip *InfixPattern) GetToken() token.Token {
	return ip.GetToken()
}

func (ip *InfixPattern) String() string {
	return fmt.Sprintf("%s %s %s", ip.Left.String(), ip.Op, ip.Right.String())
}
