package math

import (
	"zzc/fall-script/src/instructions/base"
	"zzc/fall-script/src/object"
	"zzc/fall-script/src/vm/rt"
)

type Neg struct {
	base.NoOperandInstruction
}

func (n *Neg) Execute(frame *rt.Frame) {
	o := frame.PopStack()

	i, ok := o.(*object.Integer)
	if !ok {
		panic("not integer!")
	}

	i.Value = -i.Value
	frame.PushStack(i)
}
