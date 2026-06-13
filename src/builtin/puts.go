package builtin

import (
	"fmt"

	"zzc/fall-script/src/object"
)

func Puts(args ...object.Object) object.Object {
	for _, arg := range args {
		fmt.Println(arg.Inspect())
	}

	return object.NULL
}
