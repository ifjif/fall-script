package ir

import (
	"bytes"

	"zzc/fall-script/src/token"
)

type DoWhileStmt struct {
	Token     token.Token
	Body      *BlockStmt
	Condition Expr
}

func (dws *DoWhileStmt) stmtNode() {}

func (dws *DoWhileStmt) String() string {
	var buf bytes.Buffer

	buf.WriteString("do")
	buf.WriteString(dws.Body.String())
	buf.WriteString("while(")
	buf.WriteString(dws.Condition.String())
	buf.WriteString(")")

	return buf.String()
}
