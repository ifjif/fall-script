package load

import (
	"zzc/fall-script/src/instructions/base"
	"zzc/fall-script/src/vm/rt"
)

type GetFree struct {
	index int
}

func (gf *GetFree) FetchOperand(br *base.ByteReader) {
	idx := br.ReadUint8()
	gf.index = int(idx)
}

func (gf *GetFree) Execute(frame *rt.Frame) {
	value := frame.GetFree(gf.index)
	frame.PushStack(value)
}
