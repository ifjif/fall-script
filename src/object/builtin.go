package object

import "fmt"

type Builtin struct {
	Fn BuiltinFunction
}

func (b *Builtin) Type() ObjectType {
	return BUILTIN_FUNCTION_OBJ
}

func (b *Builtin) Inspect() string {
	return fmt.Sprintf("builtin_function[%p]", b.Fn)
}
