package reference

import (
	"zzc/fall-script/src/instructions/base"
	"zzc/fall-script/src/object"
	"zzc/fall-script/src/vm/rt"
)

type Array struct {
	elemNum int
}

func (a *Array) FetchOperand(br *base.ByteReader) {
	elemNum := br.ReadUint16()
	a.elemNum = int(elemNum)
}

func (a *Array) Execute(frame *rt.Frame) {
	elems := frame.PopStacks(a.elemNum)

	arr := &object.Array{Elems: elems}

	frame.PushStack(arr)
}
