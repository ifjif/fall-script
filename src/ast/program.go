package ast

import (
	"bytes"

	"zzc/fall-script/src/token"
)

type Program struct {
	Stmts       []StmtNode
	Structs     map[string]*StructDeclStmt
	Methods     map[string][]*MethodDeclExpr
	Imports     []*ImportStmt
	Exports     []*ExportStmt
	PromotedFns []string
}

func (p *Program) String() string {
	var buf bytes.Buffer

	for _, s := range p.Stmts {
		buf.WriteString(s.String())
		buf.WriteString("\n")
	}

	return buf.String()
}

func (p *Program) TokenValue() string {
	return p.Stmts[0].TokenValue()
}

func (p *Program) GetToken() token.Token {
	return p.Stmts[0].GetToken()
}
