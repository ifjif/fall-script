package ast

import (
	"bytes"

	"zzc/fall-script/src/token"
)

type DoWhileStmt struct {
	Token     token.Token
	Body      *BlockStmt
	Condition ExprNode
}

func (dws *DoWhileStmt) StmtNode() {}

func (dws *DoWhileStmt) TokenValue() string {
	return dws.Token.Value
}

func (dws *DoWhileStmt) GetToken() token.Token {
	return dws.Token
}

func (dws *DoWhileStmt) SetAttributes(attrs []*AttributeExpr) {
}

func (dws *DoWhileStmt) GetAttributes() []*AttributeExpr {
	return nil
}

func (dws *DoWhileStmt) String() string {
	var buf bytes.Buffer

	buf.WriteString("do")
	buf.WriteString(dws.Body.String())
	buf.WriteString("while(")
	buf.WriteString(dws.Condition.String())
	buf.WriteString(")")

	return buf.String()
}
