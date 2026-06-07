package object

import (
	"bytes"
	"strings"
)

type HashKey struct {
	Type  ObjectType
	Value uint64
}

type HashPair struct {
	Key   Object
	Value Object
}

type Hash struct {
	Pairs map[HashKey]HashPair
}

func (h *Hash) Type() ObjectType {
	return HASH_OBJ
}

func (h *Hash) Inspect() string {
	var buf bytes.Buffer

	buf.WriteString("{")
	pairs := make([]string, 0)
	for _, pair := range h.Pairs {
		key := pair.Key.Inspect()
		value := pair.Value.Inspect()

		pairs = append(pairs, key+":"+value)
	}
	buf.WriteString(strings.Join(pairs, ", "))

	buf.WriteString("}")

	return buf.String()
}
