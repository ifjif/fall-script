package builtin

import "zzc/fall-script/src/object"

type BuiltinDef struct {
	Name string
	Fn   *object.Builtin
}
