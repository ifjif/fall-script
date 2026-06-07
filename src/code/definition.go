package code

type Definition struct {
	Name          string
	OperandWidths []int
}

var definitions = map[OpCode]Definition{
	NOP:   {"nop", []int{}},
	CONST: {"const", []int{2}},
}
