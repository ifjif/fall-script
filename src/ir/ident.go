package ir

import (
	"fmt"

	"zzc/fall-script/src/token"
)

type Ident struct {
	Token  token.Token
	Symbol *Symbol
}

func (i *Ident) exprNode() {}

func (i *Ident) String() string {
	return fmt.Sprintf("(name: %s pos: %d scope: %s captured: %t)", i.Symbol.Name, i.Symbol.Pos, i.Symbol.Scope, i.Symbol.Captured)
}
