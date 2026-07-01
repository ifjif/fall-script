package builtin

import (
	"fmt"

	"zzc/fall-script/src/object"
)

func Len(args ...object.Object) object.Object {
	c := len(args)

	if c != 1 {
		msg := fmt.Sprintf("expected arg counts = 1, got = %d", c)
		panic(msg)
	}

	arg1 := args[0]

	t := arg1.Type()

	switch t {
	case object.ARRAY_OBJ:
		arr := arg1.(*object.Array)
		return &object.Integer{Value: int64(len(arr.Elems))}
	case object.HASH_OBJ:
		hash := arg1.(*object.Hash)
		return &object.Integer{Value: int64(len(hash.Pairs))}
	}

	msg := fmt.Sprintf("expected collection type, got %s", t)
	panic(msg)
}
