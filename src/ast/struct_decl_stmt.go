package ast

import (
	"bytes"
	"strings"

	"zzc/fall-script/src/token"
)

type StructDeclStmt struct {
	Token  token.Token
	Name   *IdentExpr
	Fields []*FieldDeclExpr
}

func (sd *StructDeclStmt) StmtNode() {}

func (sd *StructDeclStmt) TokenValue() string {
	return sd.Token.Value
}

func (sd *StructDeclStmt) GetToken() token.Token {
	return sd.Token
}

func (sd *StructDeclStmt) String() string {
	var buf bytes.Buffer

	buf.WriteString(sd.TokenValue())
	buf.WriteString(" ")
	buf.WriteString(sd.Name.String())
	buf.WriteString("{\n")

	fields := make([]string, len(sd.Fields))

	for i, field := range sd.Fields {
		fields[i] = field.String()
	}

	buf.WriteString(strings.Join(fields, "\n"))
	buf.WriteString("\n}\n")

	return buf.String()
}

func (sd *StructDeclStmt) SetAttributes(attrs []*AttributeExpr) {
}

func (sd *StructDeclStmt) GetAttributes() []*AttributeExpr {
	return nil
}
