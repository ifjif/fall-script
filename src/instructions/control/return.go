package control

import (
	"zzc/fall-script/src/instructions/base"
	"zzc/fall-script/src/object"
	"zzc/fall-script/src/vm/rt"
)

type Return struct {
	base.NoOperandInstruction
}

func (r *Return) Execute(frame *rt.Frame) {
	lower := frame.PrevFrame()
	if lower != nil {
		lower.PushStack(object.NULL)
	}
	thread := frame.Thread()
	thread.PopFrame()
}

type XReturn struct {
	base.NoOperandInstruction
}

func (xr *XReturn) Execute(frame *rt.Frame) {
	rv := frame.PopStack()

	lower := frame.PrevFrame()
	if lower != nil {
		lower.PushStack(rv)
	}

	frame.Thread().PopFrame()
}
