package ir

import "zzc/fall-script/src/token"

type ExprStmt struct {
	Token token.Token
	Expr  Expr
}

func (es *ExprStmt) stmtNode() {}

func (es *ExprStmt) String() string {
	return es.Expr.String()
}
