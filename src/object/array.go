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

func (a *Array) Len() int {
	return len(a.Elems)
}

func (a *Array) Cap() int {
	return cap(a.Elems)
}

func (a *Array) Slice(start, end, capc int) Sliceable {
	if capc == SliceOmitted {
		capc = a.Cap()
	}
	return &Array{Elems: a.Elems[start:end:capc]}
}

func (a *Array) SliceCopy(start, end, step int) Sliceable {
	length := calcNewLen(start, end, step)

	newElems := make([]Object, length)

	elems := a.Elems
	for i, idx := 0, start; i < length; i, idx = i+1, idx+step {
		newElems[i] = elems[idx]
	}

	return &Array{Elems: newElems}
}
