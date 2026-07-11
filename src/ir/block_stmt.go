package ir

import (
	"bytes"
	"strings"

	"zzc/fall-script/src/token"
)

type BlockStmt struct {
	Token token.Token
	Stmts []Stmt
}

func (bs *BlockStmt) stmtNode() {
}

func (bs *BlockStmt) String() string {
	var buf bytes.Buffer

	buf.WriteString("{\n")
	stmts := make([]string, len(bs.Stmts))

	for i, stmt := range bs.Stmts {
		stmts[i] = stmt.String()
	}
	buf.WriteString(strings.Join(stmts, "\n"))
	buf.WriteString("\n}")

	return buf.String()
}
