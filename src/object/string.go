package object

import (
	"hash/fnv"
	"strings"
)

type String struct {
	Value string
}

func (s *String) Type() ObjectType {
	return STRING_OBJ
}

func (s *String) Inspect() string {
	return s.Value
}

func (s *String) HashKey() HashKey {
	h := fnv.New64()

	return HashKey{Type: s.Type(), Value: h.Sum64()}
}

func (s *String) Len() int {
	return len(s.Value)
}

func (s *String) Cap() int {
	return len(s.Value)
}

func (s *String) Slice(start, end, capc int) Sliceable {
	if capc != SliceOmitted {
		panic("Error(string slice cannot specify max capacity)")
	}

	return &String{Value: s.Value[start:end]}
}

func (s *String) SliceCopy(start, end, step int) Sliceable {
	length := calcNewLen(start, end, step)

	var sb strings.Builder
	sb.Grow(length)

	for i, idx := 0, start; i < length; i, idx = i+1, idx+step {
		sb.WriteByte(s.Value[idx])
	}

	return &String{Value: sb.String()}
}
