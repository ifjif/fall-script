package rt

import "zzc/fall-script/src/object"

type LocalVars []object.Object

func (lv LocalVars) GetLocal(idx int) object.Object {
	return lv[idx]
}

func (lv LocalVars) SetLocal(idx int, value object.Object) {
	lv[idx] = value
}
