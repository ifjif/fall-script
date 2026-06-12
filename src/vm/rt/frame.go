package rt

import (
	"zzc/fall-script/src/code"
	"zzc/fall-script/src/object"
)

type Frame struct {
	closure      *object.Closure
	lower        *Frame
	operandStack *OperandStack
	localvars    LocalVars
	nextPc       int
	thread       *Thread
}

func NewFrame(thread *Thread, closure *object.Closure) *Frame {
	ops := NewOperandStack(closure.Fn.StackDepth)
	lvs := make([]object.Object, closure.Fn.LocalsNum)
	f := &Frame{
		closure:      closure,
		operandStack: ops,
		localvars:    lvs,
		thread:       thread,
	}

	return f
}

func (f *Frame) SetNextPc(pc int) {
	f.nextPc = pc
}

func (f *Frame) SetLower(lower *Frame) {
	f.lower = lower
}

func (f *Frame) Code() code.Instructions {
	return f.closure.Fn.Instructions
}

func (f *Frame) NextPc() int {
	return f.nextPc
}

func (f *Frame) GetConst(index int) object.Object {
	return f.closure.Fn.Constants[index]
}

func (f *Frame) PushStack(o object.Object) {
	f.operandStack.Push(o)
}

func (f *Frame) PopStack() object.Object {
	return f.operandStack.Pop()
}

func (f *Frame) PopStacks(num int) []object.Object {
	return f.operandStack.Pops(num)
}

func (f *Frame) GetLocal(idx int) object.Object {
	return f.localvars.GetLocal(idx)
}

func (f *Frame) SetLocal(idx int, value object.Object) {
	f.localvars.SetLocal(idx, value)
}

func (f *Frame) GetFree(idx int) object.Object {
	return f.closure.Free[idx]
}

func (f *Frame) GetGlobal(idx int) object.Object {
	return f.Thread().GetGlobal(idx)
}

func (f *Frame) SetGlobal(idx int, value object.Object) {
	f.Thread().SetGlobal(idx, value)
}

func (f *Frame) GetBuiltin(idx int) object.Object {
	return f.Thread().GetBuiltin(idx)
}

func (f *Frame) PrevFrame() *Frame {
	return f.lower
}

func (f *Frame) Thread() *Thread {
	return f.thread
}
