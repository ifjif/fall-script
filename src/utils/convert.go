package utils

import (
	. "zzc/fall-script/src/object"
)

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

func ObjectToInteger(o Object) (int64, *ErrorObj) {
	switch o := o.(type) {
	case *Integer:
		return o.Value, nil
	case *Byte:
		return int64(o.Value), nil
	}

	return 0, ConvertToIntegerErr(o)
}

func convertSliceInfos(info []Object) ([]int, Object) {
	ni := make([]int, len(info))

	for i, o := range info {
		if IsNull(o) {
			if i == 2 {
				ni[i] = 1
			} else {
				ni[i] = SliceOmitted
			}
			continue
		}

		v, err := ObjectToInteger(o)
		if err != nil {
			return nil, err
		}
		ni[i] = int(v)
	}

	return ni, nil
}
