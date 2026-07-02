package store

import (
	"zzc/fall-script/src/instructions/base"
	"zzc/fall-script/src/object"
	"zzc/fall-script/src/vm/rt"
)

type SetFree struct {
	index int
}

func (sf *SetFree) FetchOperand(br *base.ByteReader) {
	i := br.ReadUint8()
	sf.index = int(i)
}

func (sf *SetFree) Execute(frame *rt.Frame) {
	v := frame.PopStack()
	o := frame.GetFree(sf.index)
	if o.Type() != object.BOX_OBJ {
		// todo
		panic("expected type box, got xxx")
	}
	box := o.(*object.Box)
	box.Value = v
}
