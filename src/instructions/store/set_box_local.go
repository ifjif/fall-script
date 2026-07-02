package store

import (
	"zzc/fall-script/src/instructions/base"
	"zzc/fall-script/src/object"
	"zzc/fall-script/src/vm/rt"
)

type SetBoxLocal struct {
	index int
}

func (sb *SetBoxLocal) FetchOperand(br *base.ByteReader) {
	i := br.ReadUint8()
	sb.index = int(i)
}

func (sb *SetBoxLocal) Execute(frame *rt.Frame) {
	v := frame.PopStack()
	o := frame.GetLocal(sb.index)
	if o.Type() != object.BOX_OBJ {
		// todo
		panic("expected type box, got xxx")
	}
	box := o.(*object.Box)
	box.Value = v
}
