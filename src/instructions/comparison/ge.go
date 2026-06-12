package comparison

import (
	"zzc/fall-script/src/instructions/base"
	"zzc/fall-script/src/utils"
	"zzc/fall-script/src/vm/rt"
)

type Ge struct {
	base.NoOperandInstruction
}

func (g *Ge) Execute(frame *rt.Frame) {
	right := frame.PopStack()
	left := frame.PopStack()

	result := utils.Compare(">=", left, right)

	frame.PushStack(result)
}
