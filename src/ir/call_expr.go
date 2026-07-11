package ir

import (
	"bytes"
	"strings"

	"zzc/fall-script/src/token"
)

type CallExpr struct {
	Token  token.Token
	Callee Expr
	Args   []Expr
}

func (ce *CallExpr) exprNode() {}

func (ce *CallExpr) String() string {
	var buf bytes.Buffer

	buf.WriteString(ce.Callee.String())
	buf.WriteString("(")
	args := make([]string, len(ce.Args))

	for i, arg := range ce.Args {
		args[i] = arg.String()
	}

	buf.WriteString(strings.Join(args, ", "))
	buf.WriteString(")")

	return buf.String()
}
