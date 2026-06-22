package ast

import (
	"bytes"

	"zzc/fall-script/src/token"
)

type ExportStmt struct {
	Token       token.Token
	Declaration Node
}

func (es *ExportStmt) StmtNode() {}

func (es *ExportStmt) TokenValue() string {
	return es.Token.Value
}

func (es *ExportStmt) GetToken() token.Token {
	return es.Token
}

func (es *ExportStmt) SetAttributes(attrs []*AttributeExpr) {
}

func (es *ExportStmt) GetAttributes() []*AttributeExpr {
	return nil
}

func (es *ExportStmt) String() string {
	var buf bytes.Buffer

	buf.WriteString(es.TokenValue())
	buf.WriteString(" ")
	buf.WriteString(es.Declaration.String())
	buf.WriteString("\n")

	return buf.String()
}
