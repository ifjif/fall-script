package rt

import "zzc/fall-script/src/object"

type Frame struct {
	Lower        *Frame
	Code         []byte
	NextPc       int
	operandStack []object.Object
	size         int
}

func NewFrame(code []byte) *Frame {
	return &Frame{
		Code:         code,
		operandStack: []object.Object{},
	}
}

func (f *Frame) Push(o object.Object) {
	f.operandStack = append(f.operandStack, o)
	f.size++
}

func (f *Frame) Pop() object.Object {
	f.size--
	o := f.operandStack[f.size]
	f.operandStack = f.operandStack[:f.size]
	return o
}
