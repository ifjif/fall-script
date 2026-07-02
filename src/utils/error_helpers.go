package utils

import (
	"fmt"

	. "zzc/fall-script/src/object"
)

func objectMsg(o Object) string {
	otype := o.Type()
	ovalue := o.Inspect()
	return fmt.Sprintf("%s(%s)", otype, ovalue)
}

func unaryOperatorMsg(op string, right Object) string {
	rightType := right.Type()
	rightValue := right.Inspect()
	return fmt.Sprintf("%s%s(%s)", op, rightType, rightValue)
}

func binaryOperatorMsg(left Object, right Object, op string) string {
	leftType := left.Type()
	leftValue := left.Inspect()
	rightType := right.Type()
	rightValue := right.Inspect()
	return fmt.Sprintf("%s(%s) %s %s(%s)", leftType, leftValue, op, rightType, rightValue)
}
