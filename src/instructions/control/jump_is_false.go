package control

import (
	"zzc/fall-script/src/instructions/base"
	"zzc/fall-script/src/utils"
	"zzc/fall-script/src/vm/rt"
)

type JumpIsFalse struct {
	pc int
}

func (j *JumpIsFalse) FetchOperand(br *base.ByteReader) {
	pc := br.ReadUint16()
	j.pc = int(pc)
}

func (j *JumpIsFalse) Execute(frame *rt.Frame) {
	cond := frame.PopStack()

	result := utils.ObjectToBool(cond)

	if !result {
		frame.SetNextPc(j.pc)
	}
}
