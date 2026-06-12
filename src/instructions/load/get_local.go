package load

import (
	"zzc/fall-script/src/instructions/base"
	"zzc/fall-script/src/vm/rt"
)

type GetLocal struct {
	index int
}

func (gl *GetLocal) FetchOperand(br *base.ByteReader) {
	idx := br.ReadUint8()
	gl.index = int(idx)
}

func (gl *GetLocal) Execute(frame *rt.Frame) {
	value := frame.GetLocal(gl.index)
	frame.PushStack(value)
}
