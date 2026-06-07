package object

type Ret struct {
	Value Object
}

func (rv *Ret) Type() ObjectType {
	return RETURN_OBJ
}

func (rv *Ret) Inspect() string {
	return rv.Value.Inspect()
}
