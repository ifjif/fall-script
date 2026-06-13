package funct

import (
	"zzc/fall-script/src/instructions/base"
	"zzc/fall-script/src/vm/rt"
)

type CurClosure struct {
	base.NoOperandInstruction
}

func (cc *CurClosure) Execute(frame *rt.Frame) {
	cl := frame.Closure()
	frame.PushStack(cl)
}
