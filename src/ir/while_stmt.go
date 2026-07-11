package ir

import (
	"bytes"

	"zzc/fall-script/src/token"
)

type WhileStmt struct {
	Token     token.Token
	Condition Expr
	Body      *BlockStmt
}

func (ws *WhileStmt) stmtNode() {}

func (ws *WhileStmt) String() string {
	var buf bytes.Buffer

	buf.WriteString("while(")
	buf.WriteString(ws.Condition.String())
	buf.WriteString(")")
	buf.WriteString(ws.Body.String())

	return buf.String()
}
