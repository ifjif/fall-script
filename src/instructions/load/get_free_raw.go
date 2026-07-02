package load

import (
	"zzc/fall-script/src/instructions/base"
	"zzc/fall-script/src/object"
	"zzc/fall-script/src/vm/rt"
)

type GetFreeRaw struct {
	index int
}

func (gfr *GetFreeRaw) FetchOperand(br *base.ByteReader) {
	i := br.ReadUint8()
	gfr.index = int(i)
}

func (gfr *GetFreeRaw) Execute(frame *rt.Frame) {
	box := frame.GetFree(gfr.index)
	if box.Type() != object.BOX_OBJ {
		// todo
		panic("expected box type got xxx")
	}

	frame.PushStack(box)
}
