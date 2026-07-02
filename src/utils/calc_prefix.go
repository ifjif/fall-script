package utils

import . "zzc/fall-script/src/object"

func CalcPrefix(op string, right Object) Object {
	switch op {
	case "!":
		return calcBangPrefix(right)
	case "-":
		return calcMinusPrefix(op, right)
	}

	return UnknownUnaryOperationErr(op, right)
}

func calcBangPrefix(right Object) Object {
	result := ObjectToBool(right)
	return boolToBoolObject(!result)
}

func calcMinusPrefix(op string, right Object) Object {
	switch right := right.(type) {
	case *Integer:
		return &Integer{Value: -right.Value}
	}

	return UnknownUnaryOperationErr(op, right)
}
