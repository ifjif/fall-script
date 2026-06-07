package ast

import (
	"bytes"

	"zzc/fall-script/src/token"
)

type Program struct {
	Stmts []StmtNode
}

func (p *Program) String() string {
	var buf bytes.Buffer

	for _, s := range p.Stmts {
		buf.WriteString(s.String())
	}

	return buf.String()
}

func (p *Program) TokenValue() string {
	return p.Stmts[0].TokenValue()
}

func (p *Program) GetToken() token.Token {
	return p.Stmts[0].GetToken()
}
