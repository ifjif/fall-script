package ir

import (
	"bytes"

	"zzc/fall-script/src/token"
)

type ForStmt struct {
	Token     token.Token
	Start     Stmt
	Condition Expr
	Update    Expr
	Body      *BlockStmt
}

func (fs *ForStmt) stmtNode() {}

func (fs *ForStmt) String() string {
	var buf bytes.Buffer

	buf.WriteString("for(")
	if fs.Start != nil {
		buf.WriteString(fs.Start.String())
	}
	buf.WriteString(";")
	if fs.Condition != nil {
		buf.WriteString(fs.Condition.String())
	}
	buf.WriteString(";")
	if fs.Update != nil {
		buf.WriteString(fs.Update.String())
	}
	buf.WriteString(")")
	buf.WriteString(fs.Body.String())

	return buf.String()
}
