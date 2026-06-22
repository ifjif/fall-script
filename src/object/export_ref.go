package object

import "fmt"

type ExportRef struct {
	Name     int
	GlobalId int
}

func (er *ExportRef) Type() ObjectType {
	return EXPORT_REF_OBJ
}

func (er *ExportRef) Inspect() string {
	return fmt.Sprintf("export_ref[%p]", er)
}
