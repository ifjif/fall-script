package macro

import (
	"zzc/fall-script/src/ast"
	"zzc/fall-script/src/ast/modifier"
	"zzc/fall-script/src/evaluator"
	"zzc/fall-script/src/object"
	"zzc/fall-script/src/token"
)

func evalUnquote(eval *evaluator.Evaluator, quoted ast.Node) ast.Node {
	return modifier.Modify(quoted, func(node ast.Node) ast.Node {
		unquote, ok := isUnquoteCall(node)
		if !ok {
			return node
		}

		if len(unquote.Args) != 1 {
			return node
		}
		// todo 记录此ast的 line column file(如果有)
		value := eval.Eval(unquote.Args[0])

		ast := objectToAst(value)
		return ast
	})
}

func isUnquoteCall(node ast.Node) (*ast.CallExpr, bool) {
	callExpr, ok := node.(*ast.CallExpr)

	if !ok {
		return nil, false
	}

	if callExpr.Fn.TokenValue() != "unquote" {
		return nil, false
	}

	return callExpr, true
}

func objectToAst(obj object.Object) ast.Node {
	switch obj := obj.(type) {
	case *object.Integer:
		t := token.Token{
			Type:  token.INT,
			Value: obj.Inspect(),
		}
		return &ast.IntExpr{
			Token: t,
			Value: obj.Value,
		}

	case *object.Boolean:
		var t token.Token
		if obj.Value {
			t = token.Token{
				Type:  token.TRUE,
				Value: "true",
			}
		} else {
			t = token.Token{
				Type:  token.FALSE,
				Value: "false",
			}
		}

		return &ast.BoolExpr{
			Token: t,
			Value: obj.Value,
		}

	case *object.Quote:
		return obj.Node

	default:
		return nil
	}
}
