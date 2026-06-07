package ast

import "zzc/fall-script/src/token"

type Node interface {
	String() string
	TokenValue() string
	GetToken() token.Token
}

type StmtNode interface {
	Node
	StmtNode()
}

type ExprNode interface {
	Node
	ExprNode()
}
