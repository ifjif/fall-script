package utils

import (
	"zzc/fall-script/src/object"
	. "zzc/fall-script/src/object"
)

func CalcIndex(left, index Object) Object {
	lt := left.Type()
	it := index.Type()
	switch {
	case lt == ARRAY_OBJ && it == INTEGER_OBJ:
		return calcArrayIndex(left, index)
	case lt == HASH_OBJ:
		return calcHashIndex(left, index)
	case lt == STRING_OBJ:
		return calcStringIndex(left, index)
	case lt == QUOTE_OBJ:
		return calcQuoteIndex(left, index)
	}

	return UnsupportedIndexOperationErr(left, index)
}

func calcArrayIndex(left, index Object) Object {
	arr := left.(*Array)
	idx := index.(*Integer)

	length := len(arr.Elems)

	if idx.Value < 0 || int(idx.Value) >= length {
		return IndexOutOfBoundErr(left, index)
	}

	return arr.Elems[idx.Value]
}

func calcHashIndex(left, index Object) Object {
	hashObj := left.(*Hash)

	key, ok := index.(HashTableKey)
	if !ok {
		return UnusableAsHashKeyErr(index)
	}

	pair, ok := hashObj.Pairs[key.HashKey()]

	if !ok {
		return NULL
	}

	return pair.Value
}

// todo 添加 byte object
func calcStringIndex(left, index Object) Object {
	strObj := left.(*String)

	idx := index.(*Integer).Value
	v := strObj.Value
	length := len(v)

	if idx < 0 || idx >= int64(length) {
		return IndexOutOfBoundErr(left, index)
	}

	return &object.Byte{Value: v[idx]}
}

func calcQuoteIndex(left, index Object) Object {
	quoteObj := left.(*Quote)

	return quoteObj.Extract(index)
}
