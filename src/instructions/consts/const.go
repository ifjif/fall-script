package consts

import (
	"zzc/fall-script/src/instructions/base"
	"zzc/fall-script/src/vm/rt"
)

type Const struct {
	Index uint16
}

func (c *Const) FetchOperand(br *base.ByteReader) {
	c.Index = br.ReadUint16()
}

func (c *Const) Exeucte(frame *rt.Frame) {
}
