package math

import (
	"zzc/fall-script/src/instructions/base"
	"zzc/fall-script/src/object"
	"zzc/fall-script/src/utils"
	"zzc/fall-script/src/vm/rt"
)

type Sub struct {
	base.NoOperandInstruction
}

func (s *Sub) Execute(frame *rt.Frame) {
	right := frame.PopStack()
	left := frame.PopStack()

	result := utils.CalcInfix("-", left, right)

	if err, ok := result.(*object.ErrorObj); ok {
		panic(err.Msg)
	}

	frame.PushStack(result)
}
