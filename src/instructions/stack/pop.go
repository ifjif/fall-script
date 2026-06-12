package stack

import (
	"zzc/fall-script/src/instructions/base"
	"zzc/fall-script/src/vm/rt"
)

type Pop struct {
	base.NoOperandInstruction
}

func (p *Pop) Execute(frame *rt.Frame) {
	frame.PopStack()
}
