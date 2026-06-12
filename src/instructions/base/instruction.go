package base

import "zzc/fall-script/src/vm/rt"

type Instruction interface {
	FetchOperand(br *ByteReader)
	Execute(frame *rt.Frame)
}

type NoOperandInstruction struct{}

func (noi *NoOperandInstruction) FetchOperand(br *ByteReader) {}
