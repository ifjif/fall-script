package object

import "fmt"

type GlobalRef struct {
	TargetModule *CompiledModule
	TargetIdx    int
}

func (gr *GlobalRef) Type() ObjectType {
	return GLOBAL_REF_OBJ
}

func (gr *GlobalRef) Inspect() string {
	return fmt.Sprintf("global_ref[%p]", gr)
}
