package load

import (
	"zzc/fall-script/src/instructions/base"
	"zzc/fall-script/src/object"
	"zzc/fall-script/src/vm/rt"
)

type GetBoxLocal struct {
	index int
}

func (gb *GetBoxLocal) FetchOperand(br *base.ByteReader) {
	idx := br.ReadUint8()
	gb.index = int(idx)
}

func (gb *GetBoxLocal) Execute(frame *rt.Frame) {
	o := frame.GetLocal(gb.index)
	if o.Type() != object.BOX_OBJ {
		// todo
		panic("xxxxx")
	}
	box := o.(*object.Box)
	frame.PushStack(box.Value)
}
