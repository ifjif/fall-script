package code

import "encoding/binary"

type OpCode byte

const (
	NOP OpCode = iota
	CONST
	LOAD
)

func Make(op OpCode, operands ...int) []byte {
	def, ok := definitions[op]

	if !ok {
		return []byte{}
	}

	length := 1

	for _, width := range def.OperandWidths {
		length += width
	}

	inst := make([]byte, length)
	inst[0] = byte(op)
	offset := 1

	for i, operand := range operands {
		width := def.OperandWidths[i]

		switch width {
		case 1:
			inst[offset] = byte(operand)
		case 2:
			binary.BigEndian.PutUint16(inst[offset:], uint16(operand))
		}

		offset += width
	}

	return inst
}
