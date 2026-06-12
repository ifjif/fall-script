package builtin

import "zzc/fall-script/src/object"

func init() {
	register("puts", puts)
	// len
	// first
	// last
	// rest
	// push
}

type BuiltinDef struct {
	Name string
	Fn   *object.Builtin
}

func register(name string, fn object.BuiltinFunction) {
	bd := &BuiltinDef{Name: name, Fn: &object.Builtin{Fn: fn}}
	Builtins = append(Builtins, bd)
}

var Builtins = []*BuiltinDef{}
