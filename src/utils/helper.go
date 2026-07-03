package utils

import "zzc/fall-script/src/object"

func IsNull(o object.Object) bool {
	return o.Type() == object.NULL_OBJ
}
