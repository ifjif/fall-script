package store

import (
	"zzc/fall-script/src/instructions/base"
	"zzc/fall-script/src/object"
	"zzc/fall-script/src/vm/rt"
)

type NewBoxLocal struct {
	index int
}

func (nb *NewBoxLocal) FetchOperand(br *base.ByteReader) {
	i := br.ReadUint8()
	nb.index = int(i)
}

func (nb *NewBoxLocal) Execute(frame *rt.Frame) {
	o := frame.PopStack()
	box := &object.Box{Value: o}
	frame.SetLocal(nb.index, box)
}
