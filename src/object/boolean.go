package object

import "strconv"

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType {
	return BOOLEAN_OBJ
}

func (b *Boolean) Inspect() string {
	return strconv.FormatBool(b.Value)
}

func (b *Boolean) HashKey() HashKey {
	v := 0
	if b.Value {
		v = 1
	}

	return HashKey{Type: b.Type(), Value: uint64(v)}
}

var (
	TRUE  = &Boolean{Value: true}
	FALSE = &Boolean{Value: false}
)
