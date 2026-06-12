package utils

import "zzc/fall-script/src/object"

func Compare(op string, left, right object.Object) object.Object {
	lv := left.(*object.Integer).Value
	rv := right.(*object.Integer).Value

	result := false
	switch op {
	case "==":
		result = lv == rv
	case "!=":
		result = lv != rv
	case ">":
		result = lv > rv
	case ">=":
		result = lv >= rv
	}

	if result {
		return object.TRUE
	}

	return object.FALSE
}
