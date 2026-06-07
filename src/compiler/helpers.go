package compiler

import (
	"zzc/fall-script/src/code"
	"zzc/fall-script/src/object"
)

func (c *Compiler) emit(op code.OpCode, operands ...int) {
	inst := code.Make(op, operands...)

	c.Instructions = append(c.Instructions, inst...)
}

func (c *Compiler) addConstant(o object.Object) int {
	c.Constants = append(c.Constants, o)
	return len(c.Constants) - 1
}
