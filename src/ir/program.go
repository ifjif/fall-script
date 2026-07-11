package ir

import "bytes"

type Program struct {
	Stmts   []Stmt
	Locals  int
	Globals int
}

func (p *Program) stmtNode() {}

func (p *Program) String() string {
	var buf bytes.Buffer

	for _, stmt := range p.Stmts {
		buf.WriteString(stmt.String())
		buf.WriteString("\n")
	}

	return buf.String()
}
