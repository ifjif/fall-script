package object

import (
	"fmt"

	"zzc/fall-script/src/code"
)

type CompiledFunction struct {
	Constants    []Object
	Instructions code.Instructions
	LocalsNum    int
	ParamsNum    int
	StackDepth   int
}

func (cf *CompiledFunction) Type() ObjectType {
	return COMPILED_FUNCTION_OBJ
}

func (cf *CompiledFunction) Inspect() string {
	return fmt.Sprintf("CompiledFunction[%p]", cf)
}
