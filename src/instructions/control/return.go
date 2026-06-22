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
	curModule := frame.Closure().OwnerModule
	lower := frame.PrevFrame()
	var prevModule *object.CompiledModule

	if lower != nil {
		prevModule = lower.Closure().OwnerModule
		if curModule == prevModule {
			lower.PushStack(object.NULL)
		}
	}
	thread := frame.Thread()
	thread.PopFrame()

	// 模块结束
	if curModule != prevModule {
		curModule.Status = object.Initialized
	}
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
