package code

import (
	"encoding/binary"
)

type OpCode byte

const (
	Nop OpCode = iota
	Null_
	Const
	Pop
	Gt
	Ge
	Band
	Bor
	Add
	Sub
	Mul
	Div
	Eq
	Neq
	Neg
	Not
	True
	False
	Array_
	Hash_
	Index
	Slice
	Call
	Jump
	JumpIsFalse
	SetGlobal
	GetGlobal
	SetLocal
	GetLocal
	NewBoxLocal
	SetBoxLocal
	GetBoxLocal
	SetFree
	GetFree
	GetFreeRaw
	GetBuiltin
	SetIndex
	Closure_
	CurClosure
	Dup
	InitStruct
	GetField
	SetField
	CallMethod
	Return
	XReturn
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
