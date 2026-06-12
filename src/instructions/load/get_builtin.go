package load

import (
	"zzc/fall-script/src/instructions/base"
	"zzc/fall-script/src/vm/rt"
)

type GetBuiltin struct {
	index int
}

func (gb *GetBuiltin) FetchOperand(br *base.ByteReader) {
	index := br.ReadUint8()
	gb.index = int(index)
}

func (gb *GetBuiltin) Execute(frame *rt.Frame) {
	bo := frame.GetBuiltin(gb.index)
	frame.PushStack(bo)
}
