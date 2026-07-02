package utils

import . "zzc/fall-script/src/object"

func CalcInfix(op string, left, right Object) Object {
	leftType := left.Type()
	rightType := right.Type()

	switch {
	case leftType == INTEGER_OBJ && rightType == INTEGER_OBJ:
		return calcIntegerInfix(op, left, right)
	case leftType == STRING_OBJ && rightType == STRING_OBJ:
		return calcStringInfix(op, left, right)
	case op == "||" || op == "&&":
		return calcLogicInfix(op, left, right)
	case op == "==":
		return boolToBoolObject(left == right)
	case op == "!=":
		return boolToBoolObject(left != right)
	}

	return UnknownBinaryOperationErr(left, right, op)
}

func calcIntegerInfix(op string, left, right Object) Object {
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
			return DivideByZeroErr(left, right, op)
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

	return UnknownBinaryOperationErr(left, right, op)
}

func calcStringInfix(op string, left, right Object) Object {
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

	return UnknownBinaryOperationErr(left, right, op)
}

func calcLogicInfix(op string, left, right Object) Object {
	leftBool := ObjectToBool(left)
	rightBool := ObjectToBool(right)

	switch op {
	case "||":
		return boolToBoolObject(leftBool || rightBool)
	case "&&":
		return boolToBoolObject(leftBool && rightBool)
	}

	return UnknownBinaryOperationErr(left, right, op)
}
