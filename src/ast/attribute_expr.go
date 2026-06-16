package ast

import (
	"bytes"
	"strings"

	"zzc/fall-script/src/token"
)

type AttributeExpr struct {
	Token token.Token
	Name  string
	Args  []ExprNode
}

func (ab *AttributeExpr) ExprNode() {}

func (ab *AttributeExpr) GetToken() token.Token {
	return ab.Token
}

func (ab *AttributeExpr) TokenValue() string {
	return ab.Token.Value
}

func (ab *AttributeExpr) String() string {
	var buf bytes.Buffer

	buf.WriteString("#[")
	buf.WriteString(ab.Name)
	buf.WriteString("(")

	if ab.Args != nil {
		args := make([]string, len(ab.Args))

		for i, arg := range ab.Args {
			args[i] = arg.String()
		}
		buf.WriteString(strings.Join(args, ", "))
	}

	buf.WriteString(")")
	buf.WriteString("]")

	return buf.String()
}
