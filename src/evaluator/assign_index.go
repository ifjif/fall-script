package evaluator

import (
	"zzc/fall-script/src/object"
	"zzc/fall-script/src/utils"
)

func (e *Evaluator) assign4Index(container, index, value object.Object) object.Object {
	result := utils.Assign4Index(container, index, value)

	if err, ok := result.(*object.ErrorObj); ok {
		return e.appendLineAndCol(err)
	}

	return result
}
