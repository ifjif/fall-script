package ir

import (
	"bytes"

	"zzc/fall-script/src/token"
)

type ReturnStmt struct {
	Token token.Token
	Value Expr
}

func (rs *ReturnStmt) stmtNode() {}

func (rs *ReturnStmt) String() string {
	var buf bytes.Buffer

	buf.WriteString("return")
	if rs.Value != nil {
		buf.WriteString(" ")
		buf.WriteString(rs.Value.String())
	}
	buf.WriteString(";")

	return buf.String()
}
