package reference

import (
	"zzc/fall-script/src/instructions/base"
	"zzc/fall-script/src/object"
	"zzc/fall-script/src/vm/rt"
)

type SetIndex struct {
	base.NoOperandInstruction
}

func (si *SetIndex) Execute(frame *rt.Frame) {
	v := frame.PopStack()
	i := frame.PopStack()
	c := frame.PopStack()

	switch {
	case c.Type() == object.ARRAY_OBJ && i.Type() == object.INTEGER_OBJ:
		setArrayIndex(frame, c, i, v)
	case c.Type() == object.HASH_OBJ:
		setHashIndex(frame, c, i, v)
	}
}

func setArrayIndex(frame *rt.Frame, arr object.Object, index object.Object, value object.Object) {
	array := arr.(*object.Array)
	idx := index.(*object.Integer)

	if idx.Value < 0 || idx.Value >= int64(len(array.Elems)) {
		panic("Error: index out of bound")
	}

	array.Elems[idx.Value] = value
	frame.PushStack(value)
}

func setHashIndex(frame *rt.Frame, hh object.Object, k object.Object, value object.Object) {
	hash := hh.(*object.Hash)
	key, ok := k.(object.HashTableKey)
	if !ok {
		panic("Error: unusable hash key")
	}

	pair := object.HashPair{Key: k, Value: value}
	hash.Pairs[key.HashKey()] = pair

	frame.PushStack(value)
}
