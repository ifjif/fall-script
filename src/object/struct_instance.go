package object

import "fmt"

type StructInstance struct {
	Slots      []Object
	StructMeta *StructMeta
}

func (si *StructInstance) Type() ObjectType {
	return STRUCT_INSTANCE_OBJ
}

func (si *StructInstance) Inspect() string {
	return fmt.Sprintf("struct_instance[%p]", si)
}
