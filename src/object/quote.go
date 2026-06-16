package object

import "zzc/fall-script/src/ast"

/*
* 对AST的提取操作
* block
* params
* args
* ident
*
 */
type Quote struct {
	Node ast.Node
}

func (q *Quote) Type() ObjectType {
	return QUOTE_OBJ
}

func (q *Quote) Inspect() string {
	return "QUOTE(" + q.Node.String() + ")"
}

func (q *Quote) Extract(index Object) *Quote {
	switch {
	case index.Type() == STRING_OBJ:
		return q.extractKey(index)
	case index.Type() == INTEGER_OBJ:
		return q.extractIndex(index)
	}

	panic("无法进行AST-node 的提取")
}

func (q *Quote) extractKey(key Object) *Quote {
	str := key.(*String).Value

	switch str {
	case "block":
		return q.extractBlockKey()
	case "params":
		return q.extractParamsKey()
	case "args":
		return q.extractArgsKey()
	case "ident":
		return q.extractIdentKey()
	}

	panic("不存在key值")
}

func (q *Quote) extractBlockKey() *Quote {
	node := q.Node

	switch node := node.(type) {
	case *ast.FnExpr:
		return &Quote{Node: node.Body}
	}

	panic("AST-node 不存在Block")
}

func (q *Quote) extractParamsKey() *Quote {
	// todo
	panic("AST-node 不存在Params")
}

func (q *Quote) extractArgsKey() *Quote {
	// todo
	panic("AST-node 不存在Args")
}

func (q *Quote) extractIdentKey() *Quote {
	node := q.Node

	switch node := node.(type) {
	case *ast.FnExpr:
		if !node.UnName {
			return &Quote{Node: node.Ident}
		}
	}

	panic("AST-node 不存在Identifier")
}

func (q *Quote) extractIndex(index Object) *Quote {
	idx := index.(*Integer).Value
	node := q.Node

	switch node := node.(type) {
	case *ast.BlockStmt:
		stmts := node.Stmts
		if idx > 0 && int(idx) < len(stmts) {
			return &Quote{Node: stmts[idx]}
		}
	}

	panic("AST-node 无法进行索引操作")
}
