package load

import (
	"zzc/fall-script/src/instructions/base"
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
	frame.PushStack(value)
}
