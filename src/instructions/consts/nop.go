package consts

import (
	"zzc/fall-script/src/instructions/base"
	"zzc/fall-script/src/vm/rt"
)

type Nop struct {
	base.NoOperandInstruction
}

func (n *Nop) Execute(frame *rt.Frame) {}
