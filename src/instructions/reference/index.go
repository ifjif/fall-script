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

	var nv object.Object
	switch {
	case container.Type() == object.ARRAY_OBJ && index.Type() == object.INTEGER_OBJ:
		nv = arrayIndex(container, index)
	case container.Type() == object.HASH_OBJ:
		nv = hashIndex(container, index)
	case container.Type() == object.STRING_OBJ:
		nv = stringIndex(container, index)
	default:
		panic("unsupported index operation")
	}

	frame.PushStack(nv)
}

func arrayIndex(arr object.Object, index object.Object) object.Object {
	array := arr.(*object.Array)
	idx := index.(*object.Integer)
	length := len(array.Elems)
	iv := idx.Value

	if iv < 0 || iv >= int64(length) {
		panic("Error: index out of bounds")
	}

	return array.Elems[iv]
}

func hashIndex(arr object.Object, key object.Object) object.Object {
	hashKey, ok := key.(object.HashTableKey)

	if !ok {
		panic("Error: unusable hash key!")
	}

	hash := arr.(*object.Hash)
	pair, ok := hash.Pairs[hashKey.HashKey()]

	if !ok {
		return object.NULL
	}
	return pair.Value
}

// todo 添加 byte object
func stringIndex(str object.Object, index object.Object) object.Object {
	strO := str.(*object.String)

	idx := index.(*object.Integer).Value
	v := strO.Value
	length := len(v)

	if idx < 0 || idx >= int64(length) {
		panic("Error: index out of bounds")
	}

	return object.NULL
}
