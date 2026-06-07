package ast

import (
	"bytes"

	"zzc/fall-script/src/token"
)

// startStmt; conditionExpr; updateExpr
//
// 三个条件 任何一个或多个都可以缺失
//
//	for(let i = 1; i <= xx; i = i+1) {
//	 Block;
//	}
//
// for(let i = 1; i <= xx; i=i+1)
// for(let i = 1; i <= xx;)
// for(i = 1; ;i=i+1)
// for(i = 1; ;)
// for(;i<= xx; i=i+1)
// for(;;i=i+1)
// for(;i < xx;)
// for(;i<xx;i=i+1)
// ...
// for(;;)
type ForStmt struct {
	Token     token.Token
	Start     StmtNode
	Condition ExprNode
	Update    ExprNode
	Body      *BlockStmt
}

func (fs *ForStmt) StmtNode() {}

func (fs *ForStmt) TokenValue() string {
	return fs.Token.Value
}

func (fs *ForStmt) GetToken() token.Token {
	return fs.Token
}

func (fs *ForStmt) String() string {
	var buf bytes.Buffer

	buf.WriteString("for(")

	if fs.Start != nil {
		buf.WriteString(fs.Start.String())
	} else {
		buf.WriteString("; ")
	}

	if fs.Condition != nil {
		buf.WriteString(fs.Condition.String())
	}
	buf.WriteString("; ")

	if fs.Update != nil {
		buf.WriteString(fs.Update.String())
	}
	buf.WriteString(")")
	buf.WriteString(fs.Body.String())

	return buf.String()
}
