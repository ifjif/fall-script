package evaluator

import . "zzc/fall-script/src/object"

func (e *Evaluator) calculateArrayIndexExpression(left, index Object) Object {
	arr := left.(*Array)
	idx := index.(*Integer)

	length := len(arr.Elems)

	if idx.Value < 0 || int(idx.Value) >= length {
		return e.indexOutOfBoundErr(left, index)
	}

	return arr.Elems[idx.Value]
}

func (e *Evaluator) calculateHashIndexExpression(left, index Object) Object {
	hashObj := left.(*Hash)

	key, ok := index.(HashTableKey)
	if !ok {
		return e.unusableAsHashKeyErr(index)
	}

	pair, ok := hashObj.Pairs[key.HashKey()]

	if !ok {
		return NULL
	}

	return pair.Value
}
