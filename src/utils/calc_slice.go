package utils

import (
	"fmt"

	. "zzc/fall-script/src/object"
)

func CalcSlice(left Object, info []Object) Object {
	ninfo, err := convertSliceInfos(info)
	if err != nil {
		return err
	}

	s, ok := left.(Sliceable)
	if !ok {
		return NotASliceErr(left)
	}

	result := calcSlice(s, ninfo)

	return result
}

func calcSlice(left Sliceable, info []int) (result Object) {
	length := left.Len()
	start := info[0]
	end := info[1]
	step := info[2]
	capc := info[3]

	if start < 0 && start != SliceOmitted {
		start = length + start
	}
	if end < 0 && end != SliceOmitted {
		end = length + end
	}

	if step > 0 {
		if start == SliceOmitted {
			start = 0
		}
		if end == SliceOmitted {
			end = length
		}
	} else if step < 0 {
		if start == SliceOmitted {
			start = length - 1
		}
		if end == SliceOmitted {
			end = -1
		}
	}

	// 共享
	if step == 1 {
		defer func() {
			if r := recover(); r != nil {
				msg := fmt.Sprintf("Error(slice bounds out of range [start = %d, end = %d, maxCap = %d])", start, end, capc)
				result = &ErrorObj{Msg: msg}
			}
		}()
		result = left.Slice(start, end, capc)
		return result
	}

	if step == 0 {
		panic("Error(slice step cannot be zero)")
	}
	if capc != SliceOmitted {
		panic("Error(max Capcity cannot be used with step != 1)")
	}

	if step > 0 && start >= end {
		panic("Error(start index > end index while step > 0)")
	} else if step < 0 && start <= end {
		panic("Error(start index < end index while step < 0)")
	}

	// new
	result = left.SliceCopy(start, end, step)

	return result
}
