package code

import (
	"encoding/binary"
	"fmt"
)

type Definition struct {
	Name          string
	OperandWidths []int
}

var definitions = map[OpCode]*Definition{
	Nop:         {"nop", []int{}},
	Null_:       {"null", []int{}},
	Const:       {"const", []int{2}},
	Pop:         {"pop", []int{}},
	Gt:          {"gt", []int{}},
	Ge:          {"ge", []int{}},
	Add:         {"add", []int{}},
	Sub:         {"sub", []int{}},
	Mul:         {"mul", []int{}},
	Div:         {"div", []int{}},
	Eq:          {"eq", []int{}},
	Neq:         {"neq", []int{}},
	Neg:         {"neg", []int{}},
	Not:         {"not", []int{}},
	True:        {"true", []int{}},
	False:       {"false", []int{}},
	Array_:      {"array", []int{2}},
	Hash_:       {"hash", []int{2}},
	Index:       {"index", []int{}},
	Slice:       {"slice", []int{}},
	Call:        {"call", []int{2}},
	Jump:        {"jump", []int{2}},
	JumpIsFalse: {"jump_is_false", []int{2}},
	SetGlobal:   {"set_global", []int{2}},
	GetGlobal:   {"get_global", []int{2}},
	SetLocal:    {"set_local", []int{1}},
	GetLocal:    {"get_local", []int{1}},
	NewBoxLocal: {"new_box_local", []int{1}},
	SetBoxLocal: {"set_box_local", []int{1}},
	GetBoxLocal: {"get_box_local", []int{1}},
	SetFree:     {"set_free", []int{1}},
	GetFree:     {"get_free", []int{1}},
	GetFreeRaw:  {"get_free_raw", []int{1}},
	GetBuiltin:  {"get_builtin", []int{1}},
	SetIndex:    {"set_index", []int{}},
	Closure_:    {"closure", []int{2, 1}},
	CurClosure:  {"cur_closure", []int{}},
	Dup:         {"dup", []int{}},
	Return:      {"return", []int{}},
	XReturn:     {"xreturn", []int{}},
}

func LookupDef(op byte) (*Definition, error) {
	def, ok := definitions[OpCode(op)]

	if !ok {
		return nil, fmt.Errorf("opcode %d undefined", op)
	}

	return def, nil
}

func (def *Definition) ReadOperand(ins Instructions) ([]int, int) {
	operands := make([]int, len(def.OperandWidths))

	offset := 0
	for i, width := range def.OperandWidths {
		switch width {
		case 1:
			operands[i] = int(ReadUint8(ins[offset:]))
		case 2:
			operands[i] = int(ReadUint16(ins[offset:]))
		}
		offset += width
	}

	return operands, offset
}

func ReadUint8(ins Instructions) uint8 {
	return ins[0]
}

func ReadUint16(ins Instructions) uint16 {
	return binary.BigEndian.Uint16(ins)
}
