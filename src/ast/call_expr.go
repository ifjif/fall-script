package ast

import (
	"bytes"
	"strings"

	"zzc/fall-script/src/token"
)

type CallExpr struct {
	Token token.Token
	Fn    ExprNode
	Args  []ExprNode
}

func (ce *CallExpr) ExprNode() {}

func (ce *CallExpr) GetToken() token.Token {
	return ce.Token
}

func (ce *CallExpr) TokenValue() string {
	return ce.Token.Value
}

func (ce *CallExpr) String() string {
	var buf bytes.Buffer

	buf.WriteString(ce.Fn.String())
	buf.WriteString("(")

	args := []string{}
	for _, arg := range ce.Args {
		args = append(args, arg.String())
	}

	buf.WriteString(strings.Join(args, ", "))
	buf.WriteString(")")

	return buf.String()
}
