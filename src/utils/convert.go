package utils

import . "zzc/fall-script/src/object"

func ObjectToBool(o Object) bool {
	switch o := o.(type) {
	case *Null:
		return false
	case *Integer:
		return o.Value != 0
	case *String:
		return o.Value != ""
	case *Boolean:
		return o.Value
	}

	return o != nil
}

func boolToBoolObject(v bool) Object {
	if v {
		return TRUE
	}

	return FALSE
}
