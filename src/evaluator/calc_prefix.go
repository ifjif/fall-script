package evaluator

import . "zzc/fall-script/src/object"

func (e *Evaluator) calculatePrefixExpression(op string, right Object) Object {
	switch op {
	case "!":
		return e.calculateBangPrefixExpression(right)
	case "-":
		return e.calculateMinusPrefixExpression(op, right)
	}

	return e.unknownUnaryOperationErr(op, right)
}

func (e *Evaluator) calculateBangPrefixExpression(right Object) Object {
	result := objectToBool(right)
	return boolToBoolObject(!result)
}

func (e *Evaluator) calculateMinusPrefixExpression(op string, right Object) Object {
	switch right := right.(type) {
	case *Integer:
		return &Integer{Value: -right.Value}
	}

	return e.unknownUnaryOperationErr(op, right)
}
