package stack

import (
	"zzc/fall-script/src/instructions/base"
	"zzc/fall-script/src/vm/rt"
)

type Dup struct {
	base.NoOperandInstruction
}

func (d *Dup) Execute(frame *rt.Frame) {
	v := frame.PopStack()
	frame.PushStack(v)
	frame.PushStack(v)
}
