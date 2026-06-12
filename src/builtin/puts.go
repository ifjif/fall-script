package builtin

import (
	"fmt"

	"zzc/fall-script/src/object"
)

func puts(args ...object.Object) object.Object {
	for _, arg := range args {
		fmt.Println(arg.Inspect())
	}

	return object.NULL
}
