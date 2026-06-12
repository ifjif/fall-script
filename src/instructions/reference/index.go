package reference

import (
	"zzc/fall-script/src/instructions/base"
	"zzc/fall-script/src/object"
	"zzc/fall-script/src/vm/rt"
)

type Index struct {
	base.NoOperandInstruction
}

func (i *Index) Execute(frame *rt.Frame) {
	index := frame.PopStack()
	container := frame.PopStack()

	switch {
	case container.Type() == object.ARRAY_OBJ && index.Type() == object.INTEGER_OBJ:
		arrayIndex(frame, container, index)
	case container.Type() == object.HASH_OBJ:
		hashIndex(frame, container, index)
	default:
		panic("unsupported index operation")
	}
}

func arrayIndex(frame *rt.Frame, arr object.Object, index object.Object) {
	array := arr.(*object.Array)
	idx := index.(*object.Integer)
	length := len(array.Elems)
	iv := idx.Value

	if iv < 0 || iv >= int64(length) {
		panic("index out of bound")
	}

	frame.PushStack(array.Elems[iv])
}

func hashIndex(frame *rt.Frame, arr object.Object, key object.Object) {
	hashKey, ok := key.(object.HashTableKey)

	if !ok {
		panic("Error: unusable hash key!")
	}

	hash := arr.(*object.Hash)
	pair, ok := hash.Pairs[hashKey.HashKey()]

	if !ok {
		frame.PushStack(object.NULL)
	} else {
		frame.PushStack(pair.Value)
	}
}
