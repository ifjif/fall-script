package utils

import "zzc/fall-script/src/object"

func ObjectToBool(o object.Object) bool {
	switch o := o.(type) {
	case *object.Null:
		return false
	case *object.Integer:
		return o.Value != 0
	case *object.String:
		return o.Value != ""
	case *object.Boolean:
		return o.Value
	}

	return o != nil
}
