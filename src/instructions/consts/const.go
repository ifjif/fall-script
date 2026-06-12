package consts

import (
	"zzc/fall-script/src/instructions/base"
	"zzc/fall-script/src/vm/rt"
)

type Const struct {
	index int
}

func (c *Const) FetchOperand(br *base.ByteReader) {
	idx := br.ReadUint16()
	c.index = int(idx)
}

func (c *Const) Execute(frame *rt.Frame) {
	data := frame.GetConst(c.index)
	frame.PushStack(data)
}
