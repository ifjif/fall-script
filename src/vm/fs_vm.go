package vm

import (
	"zzc/fall-script/src/compiler"
	"zzc/fall-script/src/vm/rt"
	"zzc/fall-script/src/vm/rt/heap"
)

type FsVM struct {
	thread *rt.Thread
}

func NewFsVM(c *compiler.Compiler) *FsVM {
	cp := heap.NewConstPool(c.Constants)
	p := heap.NewPrototype(c.Instructions, cp)
	thread := rt.NewThread()
	frame := rt.NewFrame(p.Bytecode)
	thread.Stack.PushFrame(frame)
	return &FsVM{
		thread: thread,
	}
}

func (fv *FsVM) Run() {
	interpreter(fv.thread)
}
