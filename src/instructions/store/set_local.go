package store

import (
	"zzc/fall-script/src/instructions/base"
	"zzc/fall-script/src/vm/rt"
)

type SetLocal struct {
	index int
}

func (sl *SetLocal) FetchOperand(br *base.ByteReader) {
	idx := br.ReadUint8()
	sl.index = int(idx)
}

func (sl *SetLocal) Execute(frame *rt.Frame) {
	value := frame.PopStack()
	frame.SetLocal(sl.index, value)
}
