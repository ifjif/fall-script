package comparison

import (
	"zzc/fall-script/src/instructions/base"
	"zzc/fall-script/src/utils"
	"zzc/fall-script/src/vm/rt"
)

type Gt struct {
	base.NoOperandInstruction
}

func (g *Gt) Execute(frame *rt.Frame) {
	right := frame.PopStack()
	left := frame.PopStack()

	result := utils.Compare(">", left, right)

	frame.PushStack(result)
}
