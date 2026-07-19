package object

import "fmt"

type Module struct {
	Name      string
	GlobalNum int
	Imports   []*ImportRef
	Exports   []*ExportRef
	Structs   map[int]*StructMeta
	Cf        *CompiledFunction
}

func (m *Module) Type() ObjectType {
	return MODULE_OBJ
}

func (m *Module) Inspect() string {
	return fmt.Sprintf("Module[%p]", m)
}
