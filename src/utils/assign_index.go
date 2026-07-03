package utils

import . "zzc/fall-script/src/object"

func Assign4Index(container, index, value Object) Object {
	ct := container.Type()
	it := index.Type()

	switch {
	case ct == ARRAY_OBJ && it == INTEGER_OBJ:
		arr := container.(*Array)
		idx := index.(*Integer)

		if idx.Value < 0 || idx.Value >= int64(len(arr.Elems)) {
			return IndexOutOfBoundErr(arr, idx)
		}
		arr.Elems[idx.Value] = value
		return value
	case ct == HASH_OBJ:
		hash := container.(*Hash)
		key, err := IsHashable(index)
		if err != nil {
			return err
		}
		pair := HashPair{Key: index, Value: value}
		hash.Pairs[key.HashKey()] = pair
		return value
	}

	return UnsupportedIndexAssignOperationErr(container, index, value)
}
