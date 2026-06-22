package object

import "fmt"

type ImportRef struct {
	From     int
	Imported int
	Local    int
}

func (ir *ImportRef) Type() ObjectType {
	return IMPORT_REF_OBJ
}

func (ir *ImportRef) Inspect() string {
	return fmt.Sprintf("import_ref[%p]", ir)
}
