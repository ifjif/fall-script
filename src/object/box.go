package object

import "fmt"

type Box struct {
	Value Object
}

func (b *Box) Type() ObjectType {
	return BOX_OBJ
}

func (b *Box) Inspect() string {
	return fmt.Sprintf("box[%p]", b)
}
