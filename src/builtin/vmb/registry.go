package vmb

import (
	"zzc/fall-script/src/builtin"
	"zzc/fall-script/src/object"
)

func init() {
	register("puts", builtin.Puts)
	// len
	// first
	// last
	// rest
	// push
}

func register(name string, fn object.BuiltinFunction) {
	bd := &builtin.BuiltinDef{Name: name, Fn: &object.Builtin{Fn: fn}}
	Builtins = append(Builtins, bd)
}

var Builtins = []*builtin.BuiltinDef{}
