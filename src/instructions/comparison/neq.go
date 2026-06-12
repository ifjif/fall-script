package comparison

import (
	"zzc/fall-script/src/instructions/base"
	"zzc/fall-script/src/utils"
	"zzc/fall-script/src/vm/rt"
)

type Neq struct {
	base.NoOperandInstruction
}

func (e *Neq) Execute(frame *rt.Frame) {
	right := frame.PopStack()
	left := frame.PopStack()

	result := utils.Compare("!=", left, right)

	frame.PushStack(result)
}
