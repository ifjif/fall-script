package ast

import (
	"bytes"

	"zzc/fall-script/src/token"
)

type LetStmt struct {
	Token token.Token
	Name  *IdentExpr
	Value ExprNode
}

func (ls *LetStmt) StmtNode() {}

func (ls *LetStmt) TokenValue() string {
	return ls.Token.Value
}

func (ls *LetStmt) GetToken() token.Token {
	return ls.Token
}

func (ls *LetStmt) String() string {
	var buf bytes.Buffer

	buf.WriteString("let ")
	buf.WriteString(ls.Name.String())
	if ls.Value != nil {
		buf.WriteString(" = ")
		buf.WriteString(ls.Value.String())
	}
	buf.WriteString(";")

	return buf.String()
}
