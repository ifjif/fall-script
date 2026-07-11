package ir

import (
	"bytes"

	"zzc/fall-script/src/token"
)

type LetStmt struct {
	Token token.Token
	Name  *Ident
	Value Expr
}

func (ls *LetStmt) stmtNode() {}

func (ls *LetStmt) String() string {
	var buf bytes.Buffer

	buf.WriteString("let ")
	buf.WriteString(ls.Name.String())
	buf.WriteString(" = ")
	buf.WriteString(ls.Value.String())
	buf.WriteString(";")

	return buf.String()
}
