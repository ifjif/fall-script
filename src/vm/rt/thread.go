package rt

import (
	"zzc/fall-script/src/builtin"
	"zzc/fall-script/src/object"
)

type Thread struct {
	builtins []*builtin.BuiltinDef
	stack    *Stack
	pc       int
}

func NewThread(maxFrams int, builtins []*builtin.BuiltinDef) *Thread {
	return &Thread{
		stack:    NewStack(maxFrams),
		builtins: builtins,
	}
}

func (t *Thread) GetPc() int {
	return t.pc
}

func (t *Thread) SetPc(pc int) {
	t.pc = pc
}

func (t *Thread) NewFrame(closure *object.Closure) *Frame {
	return NewFrame(t, closure)
}

func (t *Thread) PushFrame(frame *Frame) {
	t.stack.PushFrame(frame)
}

func (t *Thread) PopFrame() *Frame {
	frame := t.stack.PopFrame()
	return frame
}

func (t *Thread) IsEmpty() bool {
	return t.stack.IsEmpty()
}

func (t *Thread) CurrentFrame() *Frame {
	return t.stack.top
}

func (t *Thread) GetBuiltin(idx int) object.Object {
	return t.builtins[idx].Fn
}
