package funct

import (
	"zzc/fall-script/src/instructions/base"
	"zzc/fall-script/src/object"
	"zzc/fall-script/src/vm/rt"
)

type Call struct {
	args int
}

func (c *Call) FetchOperand(br *base.ByteReader) {
	args := br.ReadUint16()
	c.args = int(args)
}

func (c *Call) Execute(frame *rt.Frame) {
	args := frame.PopStacks(c.args)

	callee := frame.PopStack()

	switch callee := callee.(type) {
	case *object.Closure:
		thread := frame.Thread()
		newFrame := thread.NewFrame(callee)

		for i, arg := range args {
			newFrame.SetLocal(i, arg)
		}

		thread.PushFrame(newFrame)
	case *object.Builtin:
		result := callee.Fn(args...)
		frame.PushStack(result)
	default:
		panic("Error: not function!")
	}
}
