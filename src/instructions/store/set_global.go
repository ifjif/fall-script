package store

import (
	"zzc/fall-script/src/instructions/base"
	"zzc/fall-script/src/vm/rt"
)

type SetGlobal struct {
	index int
}

func (sg *SetGlobal) FetchOperand(br *base.ByteReader) {
	idx := br.ReadUint16()
	sg.index = int(idx)
}

func (sg *SetGlobal) Execute(frame *rt.Frame) {
	value := frame.PopStack()
	frame.SetGlobal(sg.index, value)
}
