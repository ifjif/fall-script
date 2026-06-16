package macro

import (
	"zzc/fall-script/src/ast"
	"zzc/fall-script/src/ast/modifier"
	"zzc/fall-script/src/evaluator"
	"zzc/fall-script/src/object"
)

var evalFn = map[string]evaluator.EvalFn{
	"quote": quote,
}

func ExpandMacros(program *ast.Program, env *object.Environment) ast.Node {
	eval := evaluator.NewEvaluator(program, env)
	eval.SetEvalFn(evalFn)

	return modifier.Modify(program, func(node ast.Node) ast.Node {
		nnode := node
		switch node := node.(type) {
		case *ast.FnExpr:
			nnode = handleMacroCall4Fn(eval, node, env)
		case *ast.CallExpr:
			nnode = handleMacroCall4Call(eval, node, env)
		}

		return nnode
	})
}

func replaceFn(nnode ast.Node, node ast.Node) {
	if nnode == node {
		return
	}
	nn, ok := nnode.(*ast.FnExpr)
	if !ok {
		return
	}

	on, ok := node.(*ast.FnExpr)
	if !ok {
		return
	}
	nn.Name = on.Name
	nn.Params = on.Params
	nn.Attrs = on.Attrs
	nn.Token = on.Token
	nn.Ident = on.Ident
	nn.UnName = on.UnName
}

func handleMacroCall4Fn(eval *evaluator.Evaluator, node ast.Node, env *object.Environment) ast.Node {
	fnExpr, ok := node.(*ast.FnExpr)
	if !ok {
		return fnExpr
	}

	// fmt.Println("处理函数定义，找它的宏信息")
	if len(fnExpr.Attrs) <= 0 {
		return fnExpr
	}

	for _, attr := range fnExpr.Attrs {
		o, ok := env.Get(attr.Name)
		if !ok {
			return fnExpr
		}

		macro, ok := o.(*object.Macro)
		if !ok {
			return fnExpr
		}

		// fmt.Printf("找到函数的宏信息：%s\n", macro.Inspect())

		attrs := &ast.ArrExpr{
			Token:    attr.Token,
			Elements: attr.Args,
		}
		args := []ast.ExprNode{attrs, fnExpr}
		qargs := quoteArgs(args)

		nFn := applyMacroCall(eval, macro, qargs).(*ast.FnExpr)
		replaceFn(nFn, fnExpr)
		fnExpr = nFn
	}

	return fnExpr
}

func handleMacroCall4Call(eval *evaluator.Evaluator, node ast.ExprNode, env *object.Environment) ast.Node {
	callExpr := node.(*ast.CallExpr)

	macro, ok := isMacroCall(callExpr, env)
	if !ok {
		return callExpr
	}

	args := quoteArgs(callExpr.Args)

	return applyMacroCall(eval, macro, args)
}

func isMacroCall(callExpr *ast.CallExpr, env *object.Environment) (*object.Macro, bool) {
	left := callExpr.Fn
	ident, ok := left.(*ast.IdentExpr)
	if !ok {
		return nil, ok
	}

	o, ok := env.Get(ident.Value)
	if !ok {
		return nil, ok
	}

	macro, ok := o.(*object.Macro)
	if !ok {
		return nil, ok
	}

	return macro, true
}

func quoteArgs(args []ast.ExprNode) []*object.Quote {
	qargs := make([]*object.Quote, len(args))

	for i, arg := range args {
		qargs[i] = &object.Quote{Node: arg}
	}

	return qargs
}

func extendMacroEnv(macro *object.Macro, args []*object.Quote) *object.Environment {
	env := object.NewEnclosedEnvironment(macro.Env)

	for i, param := range macro.Params {
		env.Set(param.Value, args[i])
	}

	return env
}

func applyMacroCall(eval *evaluator.Evaluator, macro *object.Macro, args []*object.Quote) ast.Node {
	curEnv := eval.GetEnv()
	nenv := extendMacroEnv(macro, args)

	eval.SetEnv(nenv)
	evaluated := eval.Eval(macro.Body)
	eval.SetEnv(curEnv)

	quote, ok := (evaluated).(*object.Quote)
	if !ok {
		panic("Only supported retuning AST-nodes from macro")
	}

	return quote.Node
}
