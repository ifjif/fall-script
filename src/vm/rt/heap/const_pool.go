package heap

import "zzc/fall-script/src/object"

type ConstPool struct {
	constants []object.Object
}

func NewConstPool(constants []object.Object) *ConstPool {
	return &ConstPool{constants: constants}
}
