package consts

import (
	"zzc/fall-script/src/instructions/base"
	"zzc/fall-script/src/object"
	"zzc/fall-script/src/vm/rt"
)

type Null struct {
	base.NoOperandInstruction
}

func (n *Null) Execute(frame *rt.Frame) {
	frame.PushStack(object.NULL)
}
