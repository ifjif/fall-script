package math

import (
	"fmt"

	"zzc/fall-script/src/instructions/base"
	"zzc/fall-script/src/object"
	"zzc/fall-script/src/vm/rt"
)

type Add struct {
	base.NoOperandInstruction
}

func (a *Add) Execute(frame *rt.Frame) {
	right := frame.PopStack()
	left := frame.PopStack()

	lt := left.Type()
	rt := left.Type()

	var nv object.Object

	switch {
	case lt == object.INTEGER_OBJ && rt == object.INTEGER_OBJ:
		nv = addForInteger(left, right)
	case lt == object.STRING_OBJ && rt == object.STRING_OBJ:
		nv = addForString(left, right)
	default:
		msg := fmt.Sprintf("Error: unsupported operator %s + %s", lt, rt)
		panic(msg)
	}

	frame.PushStack(nv)
}

func addForInteger(left, right object.Object) object.Object {
	rv := right.(*object.Integer).Value
	lv := left.(*object.Integer).Value

	nv := &object.Integer{Value: rv + lv}
	return nv
}

func addForString(left, right object.Object) object.Object {
	rv := right.(*object.String).Value
	lv := left.(*object.String).Value

	nv := &object.String{Value: lv + rv}

	return nv
}
