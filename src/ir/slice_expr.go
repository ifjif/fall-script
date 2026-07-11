package ir

import (
	"strings"

	"zzc/fall-script/src/token"
)

type SliceExpr struct {
	Token token.Token
	Start Expr
	End   Expr
	Step  Expr
	Cap   Expr
}

func (se *SliceExpr) exprNode() {}

func (se *SliceExpr) String() string {
	strs := make([]string, 4)

	if se.Start != nil {
		strs[0] = se.Start.String()
	}
	if se.End != nil {
		strs[1] = se.End.String()
	}

	if se.Step != nil {
		strs[2] = se.Step.String()
	}

	if se.Cap != nil {
		strs[3] = se.Cap.String()
	}

	return strings.Join(strs, ":")
}
