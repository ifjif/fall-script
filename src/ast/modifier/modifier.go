package modifier

import (
	"zzc/fall-script/src/ast"
)

type ModifierFunc func(node ast.Node) ast.Node

func Modify(node ast.Node, modifier ModifierFunc) ast.Node {
	switch nnode := node.(type) {
	case *ast.Program:
		for i, stmt := range nnode.Stmts {
			nnode.Stmts[i] = Modify(stmt, modifier).(ast.StmtNode)
		}

	case *ast.ExprStmt:
		result := Modify(nnode.Expr, modifier)
		switch result := result.(type) {
		case ast.ExprNode:
			nnode.Expr = result
		case ast.StmtNode:
			node = result
		}

	case *ast.BlockStmt:
		newStmts := make([]ast.StmtNode, 0)
		for _, stmt := range nnode.Stmts {
			ns := Modify(stmt, modifier)
			if bl, ok := ns.(*ast.BlockStmt); ok {
				newStmts = append(newStmts, bl.Stmts...)
			} else {
				nst := ns.(ast.StmtNode)
				newStmts = append(newStmts, nst)
			}
		}
		nnode.Stmts = newStmts

	case *ast.LetStmt:
		value := Modify(nnode.Value, modifier)
		v := value.(ast.ExprNode)
		nnode.Value = v

	case *ast.CallExpr:
		nnode.Fn = Modify(nnode.Fn, modifier).(ast.ExprNode)
		for i, arg := range nnode.Args {
			nnode.Args[i] = Modify(arg, modifier).(ast.ExprNode)
		}

	case *ast.InfixExpr:
		nnode.Left = Modify(nnode.Left, modifier).(ast.ExprNode)
		nnode.Right = Modify(nnode.Right, modifier).(ast.ExprNode)

	case *ast.FnExpr:
		params := nnode.Params
		for i, param := range params {
			nnode.Params[i] = Modify(param, modifier).(*ast.IdentExpr)
		}
		nnode.Body = Modify(nnode.Body, modifier).(*ast.BlockStmt)
	}

	return modifier(node)
}
