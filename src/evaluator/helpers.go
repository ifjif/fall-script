package evaluator

import . "zzc/fall-script/src/object"

func isError(obj Object) bool {
	if obj != nil {
		return obj.Type() == ERROR_OBJ
	}

	return false
}

func boolToBoolObject(v bool) Object {
	if v {
		return TRUE
	}

	return FALSE
}

func objectToBool(o Object) bool {
	switch o := o.(type) {
	case *Integer:
		return o.Value != 0
	case *String:
		return o.Value != ""
	case *Boolean:
		return o.Value
	}

	return false
}

func unwrapReturnValue(o Object) Object {
	if result, ok := o.(*Ret); ok {
		return result.Value
	}

	return o
}
