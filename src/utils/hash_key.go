package utils

import (
	"zzc/fall-script/src/object"
	. "zzc/fall-script/src/object"
)

func IsHashable(key Object) (HashTableKey, *object.ErrorObj) {
	hashKey, ok := key.(HashTableKey)

	if !ok {
		return nil, UnusableAsHashKeyErr(key)
	}

	return hashKey, nil
}
