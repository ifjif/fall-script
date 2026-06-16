package ast

import (
	"bytes"

	"zzc/fall-script/src/token"
)

type ReturnStmt struct {
	Token token.Token
	Value ExprNode
}

func (rs *ReturnStmt) StmtNode() {}

func (rs *ReturnStmt) TokenValue() string {
	return rs.Token.Value
}

func (rs *ReturnStmt) GetToken() token.Token {
	return rs.Token
}

func (rs *ReturnStmt) SetAttributes(attrs []*AttributeExpr) {
}

func (rs *ReturnStmt) GetAttributes() []*AttributeExpr {
	return nil
}

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
