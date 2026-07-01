package evalb

import (
	"zzc/fall-script/src/builtin"
	"zzc/fall-script/src/object"
)

func init() {
	register("puts", builtin.Puts)
	register("len", builtin.Len)
	register("byte", builtin.Byte)
	// first
	// last
	// rest
	// push
}

func register(name string, fn object.BuiltinFunction) {
	Builtins[name] = &object.Builtin{Fn: fn}
}

var Builtins = map[string]*object.Builtin{}
