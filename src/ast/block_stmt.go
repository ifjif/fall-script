package ast

import (
	"bytes"
	"strings"

	"zzc/fall-script/src/token"
)

type BlockStmt struct {
	Token token.Token
	Stmts []StmtNode
}

func (bs *BlockStmt) StmtNode() {}

func (bs *BlockStmt) GetToken() token.Token {
	return bs.Token
}

func (be *BlockStmt) TokenValue() string {
	return be.Token.Value
}

func (be *BlockStmt) String() string {
	var buf bytes.Buffer

	buf.WriteString("{\n")

	strs := make([]string, len(be.Stmts))

	for i, stmt := range be.Stmts {
		strs[i] = stmt.String()
	}

	buf.WriteString(strings.Join(strs, "\n"))
	buf.WriteString("\n}")

	return buf.String()
}
