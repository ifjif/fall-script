package object

import "fmt"

type StructMeta struct {
	Name       string
	FieldCount int
	Fields     map[string]*FieldInfo
	Methods    map[string]*MethodRef
}
type MethodRef struct {
	TargetStructMeta int
	Name             string
	Index            int
}

func (sm *StructMeta) Type() ObjectType {
	return STRUCT_META_OBJ
}

func (sm *StructMeta) Inspect() string {
	return fmt.Sprintf("struct_meta[%p]", sm)
}
