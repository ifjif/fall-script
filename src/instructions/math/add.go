package math

import (
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

	rv := right.(*object.Integer).Value
	lv := left.(*object.Integer).Value

	nv := &object.Integer{Value: rv + lv}
	frame.PushStack(nv)
}
