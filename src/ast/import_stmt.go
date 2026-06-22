package ast

import (
	"bytes"
	"strings"

	"zzc/fall-script/src/token"
)

type Specifier struct {
	Imported string
	Local    string
}

type ImportStmt struct {
	Token      token.Token
	Specifiers []*Specifier
	Source     string
}

func (is *ImportStmt) StmtNode() {}

func (is *ImportStmt) TokenValue() string {
	return is.Token.Value
}

func (is *ImportStmt) GetToken() token.Token {
	return is.Token
}

func (is *ImportStmt) SetAttributes(attrs []*AttributeExpr) {
}

func (is *ImportStmt) GetAttributes() []*AttributeExpr {
	return nil
}

func (is *ImportStmt) String() string {
	var buf bytes.Buffer
	buf.WriteString(is.Token.Value)

	specifiers := make([]string, len(is.Specifiers))
	isAll := false

	for i, spe := range is.Specifiers {
		name := spe.Imported
		if name == "*" {
			isAll = true
		}
		if spe.Local != "" {
			name += " as " + spe.Local
		}
		specifiers[i] = name
	}

	buf.WriteString(" ")
	if isAll {
		buf.WriteString(strings.Join(specifiers, ", "))
	} else {
		buf.WriteString("{ ")
		buf.WriteString(strings.Join(specifiers, ", "))
		buf.WriteString(" }")
	}

	buf.WriteString(" from ")
	buf.WriteString("\"")
	buf.WriteString(is.Source)
	buf.WriteString("\"\n")

	return buf.String()
}
