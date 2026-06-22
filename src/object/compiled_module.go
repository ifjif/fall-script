package object

import "fmt"

type ModuleState int

const (
	Uninitialized ModuleState = iota
	Initializing
	Initialized
)

type CompiledModule struct {
	Globals []Object
	Exports map[string]int
	Cf      *CompiledFunction
	Status  ModuleState
}

func (cm *CompiledModule) Type() ObjectType {
	return COMPILED_MODULE_OBJ
}

func (cm *CompiledModule) Inspect() string {
	return fmt.Sprintf("compiled_module[%q]", cm)
}

func (cm *CompiledModule) Closure() *Closure {
	closure := &Closure{
		Fn:          cm.Cf,
		Free:        []Object{},
		OwnerModule: cm,
	}

	return closure
}
