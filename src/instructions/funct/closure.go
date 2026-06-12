package funct

import (
	"zzc/fall-script/src/instructions/base"
	"zzc/fall-script/src/object"
	"zzc/fall-script/src/vm/rt"
)

type Closure struct {
	index   int
	freeNum int
}

func (c *Closure) FetchOperand(br *base.ByteReader) {
	index := br.ReadUint16()
	freeNum := br.ReadUint8()
	c.index = int(index)
	c.freeNum = int(freeNum)
}

func (c *Closure) Execute(frame *rt.Frame) {
	cf, ok := frame.GetConst(c.index).(*object.CompiledFunction)
	if !ok {
		panic("Error: not function!")
	}
	frees := frame.PopStacks(c.freeNum)
	closure := &object.Closure{
		Fn:   cf,
		Free: frees,
	}

	frame.PushStack(closure)
}
