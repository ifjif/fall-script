package evaluator

import . "zzc/fall-script/src/object"

func (e *Evaluator) calculateInfixExpression(op string, left, right Object) Object {
	leftType := left.Type()
	rightType := right.Type()

	switch {
	case leftType == INTEGER_OBJ && rightType == INTEGER_OBJ:
		return e.calculateIntegerInfixExpression(op, left, right)
	case leftType == STRING_OBJ && rightType == STRING_OBJ:
		return e.calculateStringInfixExpression(op, left, right)
	case op == "||" || op == "&&":
		return e.calculateLogicInfixExpression(op, left, right)
	case op == "==":
		return boolToBoolObject(left == right)
	case op == "!=":
		return boolToBoolObject(left != right)
	}

	return e.unknownBinaryOperationErr(left, right, op)
}

func (e *Evaluator) calculateIntegerInfixExpression(op string, left, right Object) Object {
	leftValue := left.(*Integer).Value
	rightValue := right.(*Integer).Value

	switch op {
	case "+":
		return &Integer{Value: leftValue + rightValue}
	case "-":
		return &Integer{Value: leftValue - rightValue}
	case "*":
		return &Integer{Value: leftValue * rightValue}
	case "/":
		if rightValue == 0 {
			return e.divideByZeroErr(left, right, op)
		}
		return &Integer{Value: leftValue / rightValue}
	case ">":
		return boolToBoolObject(leftValue > rightValue)
	case ">=":
		return boolToBoolObject(leftValue >= rightValue)
	case "<":
		return boolToBoolObject(leftValue < rightValue)
	case "<=":
		return boolToBoolObject(leftValue <= rightValue)
	case "==":
		return boolToBoolObject(leftValue == rightValue)
	case "!=":
		return boolToBoolObject(leftValue != rightValue)
	case "||":
		return boolToBoolObject(leftValue != 0 || rightValue != 0)
	case "&&":
		return boolToBoolObject(leftValue != 0 && rightValue != 0)
	}

	return e.unknownBinaryOperationErr(left, right, op)
}

func (e *Evaluator) calculateStringInfixExpression(op string, left, right Object) Object {
	leftValue := left.(*String).Value
	rightValue := right.(*String).Value

	switch op {
	case "+":
		return &String{Value: leftValue + rightValue}
	case "||":
		return boolToBoolObject(leftValue != "" || rightValue != "")
	case "&&":
		return boolToBoolObject(leftValue != "" && rightValue != "")
	}

	return e.unknownBinaryOperationErr(left, right, op)
}

func (e *Evaluator) calculateLogicInfixExpression(op string, left, right Object) Object {
	leftBool := objectToBool(left)
	rightBool := objectToBool(right)

	switch op {
	case "||":
		return boolToBoolObject(leftBool || rightBool)
	case "&&":
		return boolToBoolObject(leftBool && rightBool)
	}

	return e.unknownBinaryOperationErr(left, right, op)
}
