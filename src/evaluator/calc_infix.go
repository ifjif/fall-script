package evaluator

import (
	"zzc/fall-script/src/object"
	. "zzc/fall-script/src/object"
	"zzc/fall-script/src/utils"
)

func (e *Evaluator) calculateInfixExpression(op string, left, right Object) Object {
	result := utils.CalcInfix(op, left, right)

	if err, ok := result.(*object.ErrorObj); ok {
		return e.appendLineAndCol(err)
	}

	return result
}
