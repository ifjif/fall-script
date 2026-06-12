package rt

import (
	"zzc/fall-script/src/object"
)

type OperandStack struct {
	slots []object.Object
	size  int
}

func NewOperandStack(size int) *OperandStack {
	slots := make([]object.Object, size)
	ops := &OperandStack{
		slots: slots,
	}

	return ops
}

func (os *OperandStack) Push(o object.Object) {
	os.slots[os.size] = o
	os.size++
}

func (os *OperandStack) Pop() object.Object {
	os.size--
	o := os.slots[os.size]

	return o
}

func (os *OperandStack) Pops(num int) []object.Object {
	last := num - 1
	result := make([]object.Object, num)
	for i := 0; i < num; i++ {
		result[last-i] = os.Pop()
	}

	return result
}
