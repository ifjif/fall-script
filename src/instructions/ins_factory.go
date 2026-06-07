package instructions

import (
	"zzc/fall-script/src/code"
	"zzc/fall-script/src/instructions/base"
)

var nop = &base.Nop{}

func NewInstruction(op code.OpCode) base.Instruction {
	switch op {
	case code.NOP:
		return nop
	}

	return nil
}
