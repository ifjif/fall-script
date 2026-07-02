package evaluator

import (
	"zzc/fall-script/src/object"
	. "zzc/fall-script/src/object"
	"zzc/fall-script/src/utils"
)

func (e *Evaluator) calculateIndexExpression(left, index Object) Object {
	result := utils.CalcIndex(left, index)

	if err, ok := result.(*object.ErrorObj); ok {
		return e.appendLineAndCol(err)
	}

	return result
}
