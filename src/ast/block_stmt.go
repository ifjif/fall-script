package ast

import (
	"bytes"

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

func (be *BlockStmt) SetAttributes([]*AttributeExpr) {
}

func (be *BlockStmt) GetAttributes() []*AttributeExpr {
	return nil
}

func (be *BlockStmt) String() string {
	var buf bytes.Buffer

	buf.WriteString("{\n")

	for _, stmt := range be.Stmts {
		buf.WriteString(stmt.String())
		buf.WriteString("\n")
	}

	buf.WriteString("\n}")

	return buf.String()
}
