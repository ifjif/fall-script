package object

import "fmt"

type StructMeta struct {
	Name        string
	FieldCount  int // 总字段数
	Fields      map[string]*FieldInfo
	Methods     map[string]*MethodRef
	OwnerModule *CompiledModule // 它所属的模块，访问它的组合字段都从这个module找
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
