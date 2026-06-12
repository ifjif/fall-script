package consts

import (
	"zzc/fall-script/src/instructions/base"
	"zzc/fall-script/src/object"
	"zzc/fall-script/src/vm/rt"
)

type True struct {
	base.NoOperandInstruction
}

func (t *True) Execute(frame *rt.Frame) {
	frame.PushStack(object.TRUE)
}

type False struct {
	base.NoOperandInstruction
}

func (t *False) Execute(frame *rt.Frame) {
	frame.PushStack(object.FALSE)
}
