package ast

import (
	"bytes"

	"zzc/fall-script/src/token"
)

type WhileStmt struct {
	Token     token.Token
	Condition ExprNode
	Body      *BlockStmt
}

func (ws *WhileStmt) StmtNode() {}

func (ws *WhileStmt) TokenValue() string {
	return ws.Token.Value
}

func (ws *WhileStmt) GetToken() token.Token {
	return ws.Token
}

func (ws *WhileStmt) String() string {
	var buf bytes.Buffer

	buf.WriteString("while(")
	buf.WriteString(ws.Condition.String())
	buf.WriteString(")")
	buf.WriteString(ws.Body.String())

	return buf.String()
}
