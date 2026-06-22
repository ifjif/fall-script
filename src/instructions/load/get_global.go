package load

import (
	"zzc/fall-script/src/instructions/base"
	"zzc/fall-script/src/object"
	"zzc/fall-script/src/vm/rt"
)

type GetGlobal struct {
	index int
}

func (gg *GetGlobal) FetchOperand(br *base.ByteReader) {
	idx := br.ReadUint16()
	gg.index = int(idx)
}

func (gg *GetGlobal) Execute(frame *rt.Frame) {
	value := frame.GetGlobal(gg.index)
	globalRef, ok := value.(*object.GlobalRef)
	if ok {
		targetModule := globalRef.TargetModule
		if targetModule.Status == object.Uninitialized {
			closure := targetModule.Closure()
			nframe := frame.Thread().NewFrame(closure)
			frame.Thread().PushFrame(nframe)
			targetModule.Status = object.Initializing
			// 恢复pc
			frame.SetNextPc(frame.Thread().GetPc())
			return
		}

		value = targetModule.Globals[globalRef.TargetIdx]
	}

	frame.PushStack(value)
}
