package logic

import (
	"zzc/fall-script/src/instructions/base"
	"zzc/fall-script/src/object"
	"zzc/fall-script/src/utils"
	"zzc/fall-script/src/vm/rt"
)

type Not struct {
	base.NoOperandInstruction
}

func (n *Not) Execute(frame *rt.Frame) {
	o := frame.PopStack()

	b := utils.ObjectToBool(o)

	if !b {
		frame.PushStack(object.TRUE)
	} else {
		frame.PushStack(object.FALSE)
	}
}
