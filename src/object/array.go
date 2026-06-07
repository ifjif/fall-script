package object

import (
	"bytes"
	"strings"
)

type Array struct {
	Elems []Object
}

func (a *Array) Type() ObjectType {
	return ARRAY_OBJ
}

func (a *Array) Inspect() string {
	var buf bytes.Buffer

	buf.WriteString("[")

	elems := make([]string, len(a.Elems))
	for i, elem := range a.Elems {
		elems[i] = elem.Inspect()
	}
	buf.WriteString(strings.Join(elems, ", "))

	buf.WriteString("]")

	return buf.String()
}
