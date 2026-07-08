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

func (es *ExprStmt) SetAttributes(attrs []*AttributeExpr) {
	if fn, ok := es.Expr.(*FnExpr); ok {
		fn.Attrs = attrs
	}

	if method, ok := es.Expr.(*MethodDeclExpr); ok {
		method.Fn.Attrs = attrs
	}
}

func (es *ExprStmt) GetAttributes() []*AttributeExpr {
	if fn, ok := es.Expr.(*FnExpr); ok {
		return fn.Attrs
	}

	if method, ok := es.Expr.(*MethodDeclExpr); ok {
		return method.Fn.Attrs
	}

	return nil
}
