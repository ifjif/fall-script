package control

import (
	"zzc/fall-script/src/instructions/base"
	"zzc/fall-script/src/vm/rt"
)

type Jump struct {
	pc int
}

func (j *Jump) FetchOperand(br *base.ByteReader) {
	pc := br.ReadUint16()
	j.pc = int(pc)
}

func (j *Jump) Execute(frame *rt.Frame) {
	frame.SetNextPc(j.pc)
}
