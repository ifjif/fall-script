package ir

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

func (ip *InfixPattern) exprNode()    {}
func (ip *InfixPattern) patternNode() {}

func (ip *InfixPattern) String() string {
	return fmt.Sprintf("%s %s %s", ip.Left, ip.Op, ip.Right)
}
