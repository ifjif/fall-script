package ast

import "zzc/fall-script/src/token"

type ExprStmt struct {
	Token token.Token
	Expr  ExprNode
}

func (es *ExprStmt) StmtNode() {}

func (es *ExprStmt) GetToken() token.Token {
	return es.Token
}

func (es *ExprStmt) TokenValue() string {
	return es.Expr.TokenValue()
}

func (es *ExprStmt) String() string {
	return es.Expr.String() + ";"
}
