package evaluator

import (
	. "zzc/fall-script/src/ast"
	"zzc/fall-script/src/object"
	. "zzc/fall-script/src/object"
)

type Evaluator struct {
	program Node
	curNode Node
	env     *Environment
}

func NewEvaluator(program Node, env *Environment) *Evaluator {
	e := &Evaluator{program: program, env: env}
	return e
}

func (e *Evaluator) Evaluate() Object {
	return e.eval(e.program)
}

func (e *Evaluator) eval(node Node) Object {
	e.curNode = node
	switch node := node.(type) {
	case *NullExpr:
		return NULL
	case *Program:
		return e.evalProgram(node.Stmts)
	case *ExprStmt:
		return e.eval(node.Expr)
	case *BlockStmt:
		return e.evalBlockStmt(node)
	case *LetStmt:
		return e.evalLetStmt(node)
	case *ForStmt:
		return e.evalForStmt(node)
	case *WhileStmt:
		return e.evalWhileStmt(node)
	case *DoWhileStmt:
		return e.evalDoWhileStmt(node)
	case *ReturnStmt:
		return e.evalReturnStmt(node)
	case *IdentExpr:
		return e.evalIdentExpr(node)
	case *IntExpr:
		return e.evalInteger(node)
	case *BoolExpr:
		return e.evalBoolean(node)
	case *StrExpr:
		return e.evalString(node)
	case *AssignExpr:
		return e.evalAssignExpr(node)
	case *InfixExpr:
		return e.evalInfixExpr(node)
	case *PrefixExpr:
		return e.evalPrefixExpr(node)
	case *ArrExpr:
		return e.evalArrExpr(node)
	case *HashExpr:
		return e.evalHashExpr(node)
	case *IndexExpr:
		return e.evalIndexExpr(node)
	case *FnExpr:
		return e.evalFnExpre(node)
	case *CallExpr:
		return e.evalCallExpr(node)
	case *IfExpr:
		return e.evalIfExpr(node)
	}

	return nil
}

func (e *Evaluator) evalProgram(stmts []StmtNode) Object {
	var result Object

	for _, stmt := range stmts {
		result = e.eval(stmt)

		switch result := result.(type) {
		case *ErrorObj, *Ret:
			return result
		}
	}

	return result
}

func (e *Evaluator) evalBlockStmt(stmt *BlockStmt) Object {
	return e.evalProgram(stmt.Stmts)
}

func (e *Evaluator) evalLetStmt(stmt *LetStmt) Object {
	name := stmt.Name.Value
	_, ok := e.env.ExistsCur(name)
	if ok {
		return e.redeclaredErr(name)
	}

	value := e.eval(stmt.Value)

	if isError(value) {
		return value
	}

	e.env.Set(name, value)

	return nil
}

func (e *Evaluator) evalForStmt(stmt *ForStmt) Object {
	curEnv := e.env
	forEnv := NewEnclosedEnvironment(curEnv)
	e.env = forEnv

	if stmt.Start != nil {
		result := e.eval(stmt.Start)
		if isError(result) {
			return result
		}
	}

	loop := false
	if stmt.Condition == nil {
		loop = true
	} else {
		loop = objectToBool(e.eval(stmt.Condition))
	}

	for loop {
		e.env = NewEnclosedEnvironment(forEnv)
		result := e.evalBlockStmt(stmt.Body)
		switch result.(type) {
		case *ErrorObj, *Ret:
			return result
		}

		e.env = forEnv
		if stmt.Update != nil {
			result := e.eval(stmt.Update)
			if isError(result) {
				return result
			}
		}

		if stmt.Condition == nil {
			loop = true
		} else {
			loop = objectToBool(e.eval(stmt.Condition))
		}
	}

	e.env = curEnv

	return nil
}

func (e *Evaluator) evalWhileStmt(stmt *WhileStmt) Object {
	curEnv := e.env

	for objectToBool(e.eval(stmt.Condition)) {
		newEnv := object.NewEnclosedEnvironment(curEnv)
		e.env = newEnv
		result := e.evalBlockStmt(stmt.Body)

		switch result.(type) {
		case *ErrorObj, *Ret:
			return result
		}

		e.env = newEnv
	}

	return nil
}

func (e *Evaluator) evalDoWhileStmt(stmt *DoWhileStmt) Object {
	curEnv := e.env
	for {
		newEnv := object.NewEnclosedEnvironment(curEnv)
		e.env = newEnv

		result := e.evalBlockStmt(stmt.Body)
		switch result.(type) {
		case *ErrorObj, *Ret:
			return result
		}

		e.env = curEnv

		if !objectToBool(e.eval(stmt.Condition)) {
			return nil
		}
	}
}

func (e *Evaluator) evalReturnStmt(stmt *ReturnStmt) Object {
	result := e.eval(stmt.Value)
	return &Ret{Value: result}
}

func (e *Evaluator) evalIdentExpr(node *IdentExpr) Object {
	name := node.Value
	result, ok := e.env.Get(name)
	if !ok {
		return e.identifierNotFoundErr(name)
	}

	return result
}

func (e *Evaluator) evalInteger(node *IntExpr) Object {
	return &Integer{node.Value}
}

func (e *Evaluator) evalBoolean(node *BoolExpr) Object {
	if node.Value {
		return TRUE
	}
	return FALSE
}

func (e *Evaluator) evalString(node *StrExpr) Object {
	return &String{Value: node.Value}
}

func (e *Evaluator) evalAssignExpr(node *AssignExpr) Object {
	name := node.Name.Value

	_, ok := e.env.Get(name)
	if !ok {
		return e.identifierNotFoundErr(name)
	}

	value := e.eval(node.Value)
	if isError(value) {
		return value
	}

	e.env.ExitsSet(name, value)

	return value
}

func (e *Evaluator) evalInfixExpr(node *InfixExpr) Object {
	curNode := e.curNode
	left := e.eval(node.Left)

	if isError(left) {
		return left
	}

	right := e.eval(node.Right)

	if isError(right) {
		return right
	}

	// 恢复为 中缀节点
	e.curNode = curNode

	return e.calculateInfixExpression(node.Op, left, right)
}

func (e *Evaluator) evalPrefixExpr(node *PrefixExpr) Object {
	curNode := e.curNode

	right := e.eval(node.Right)

	if isError(right) {
		return right
	}

	e.curNode = curNode

	return e.calculatePrefixExpression(node.Op, right)
}

func (e *Evaluator) evalArrExpr(node *ArrExpr) Object {
	elems := e.evalExprList(node.Elements)
	if len(elems) == 1 && isError(elems[0]) {
		return elems[0]
	}
	return &Array{Elems: elems}
}

func (e *Evaluator) evalHashExpr(node *HashExpr) Object {
	pairs := make(map[HashKey]HashPair)

	for _, pair := range node.Pairs {
		keyObj := e.eval(pair.Key)
		if isError(keyObj) {
			return keyObj
		}

		key, ok := keyObj.(HashTableKey)

		if !ok {
			return e.unusableAsHashKeyErr(keyObj)
		}

		valueObj := e.eval(pair.Value)
		if isError(valueObj) {
			return valueObj
		}

		hashPair := HashPair{Key: keyObj, Value: valueObj}

		pairs[key.HashKey()] = hashPair
	}

	return &Hash{Pairs: pairs}
}

func (e *Evaluator) evalIndexExpr(node *IndexExpr) Object {
	curNode := e.curNode
	leftObj := e.eval(node.Left)
	if isError(leftObj) {
		return leftObj
	}

	indexObj := e.eval(node.Index)
	if isError(indexObj) {
		return indexObj
	}
	e.curNode = curNode

	switch {
	case leftObj.Type() == ARRAY_OBJ && indexObj.Type() == INTEGER_OBJ:
		return e.calculateArrayIndexExpression(leftObj, indexObj)
	case leftObj.Type() == HASH_OBJ:
		return e.calculateHashIndexExpression(leftObj, indexObj)
	}

	return e.unsupportedIndexOperationErr(leftObj, indexObj)
}

func (e *Evaluator) evalFnExpre(node *FnExpr) Object {
	fn := &Function{Name: node.Name, Params: node.Params, Body: node.Body, Env: e.env}

	if fn.Name != "" {
		e.env.Set(fn.Name, fn)
	}

	return fn
}

func (e *Evaluator) evalCallExpr(node *CallExpr) Object {
	fn := e.eval(node.Fn)
	if isError(fn) {
		return fn
	}

	fnObj, ok := fn.(*Function)

	if !ok {
		return e.notAFunctionErr(fn)
	}

	args := e.evalExprList(node.Args)
	if len(args) == 1 && isError(args[0]) {
		return args[0]
	}

	return e.applyFunction(fnObj, args)
}

func (e *Evaluator) evalIfExpr(node *IfExpr) Object {
	cond := e.eval(node.Condition)
	if isError(cond) {
		return cond
	}

	if objectToBool(cond) {
		return unwrapReturnValue(e.eval(node.Consequence))
	} else {
		return unwrapReturnValue(e.eval(node.Alternative))
	}
}

func (e *Evaluator) evalExprList(nodes []ExprNode) []Object {
	elems := make([]Object, len(nodes))

	for i, elem := range nodes {
		o := e.eval(elem)
		if isError(o) {
			return []Object{o}
		}
		elems[i] = o
	}

	return elems
}
