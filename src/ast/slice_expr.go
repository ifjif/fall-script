package ast

import (
	"strings"

	"zzc/fall-script/src/token"
)

type SliceExpr struct {
	Token token.Token
	Start ExprNode
	End   ExprNode
	Step  ExprNode
	Cap   ExprNode
}

func (s *SliceExpr) ExprNode() {}

func (s *SliceExpr) TokenValue() string {
	return s.Token.Value
}

func (s *SliceExpr) GetToken() token.Token {
	return s.Token
}

func (s *SliceExpr) String() string {
	strs := make([]string, 4)

	if s.Start != nil {
		strs[0] = s.Start.String()
	}
	if s.End != nil {
		strs[1] = s.End.String()
	}

	if s.Step != nil {
		strs[2] = s.Step.String()
	}

	if s.Cap != nil {
		strs[3] = s.Cap.String()
	}

	return strings.Join(strs, ":")
}
