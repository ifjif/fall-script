package object

type Byte struct {
	Value byte
}

func (b *Byte) Type() ObjectType {
	return BYTE_OBJ
}

func (b *Byte) Inspect() string {
	return string(b.Value)
}

func (b *Byte) HashKey() HashKey {
	return HashKey{Type: b.Type(), Value: uint64(b.Value)}
}
