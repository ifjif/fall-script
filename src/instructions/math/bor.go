package math

import (
	"zzc/fall-script/src/instructions/base"
	"zzc/fall-script/src/object"
	"zzc/fall-script/src/utils"
	"zzc/fall-script/src/vm/rt"
)

type Bor struct {
	base.NoOperandInstruction
}

func (b *Bor) Execute(frame *rt.Frame) {
	right := frame.PopStack()
	left := frame.PopStack()

	result := utils.CalcInfix("|", left, right)

	if err, ok := result.(*object.ErrorObj); ok {
		panic(err.Msg)
	}
	frame.PushStack(result)
}
