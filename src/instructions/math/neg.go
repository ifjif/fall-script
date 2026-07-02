package math

import (
	"zzc/fall-script/src/instructions/base"
	"zzc/fall-script/src/object"
	"zzc/fall-script/src/utils"
	"zzc/fall-script/src/vm/rt"
)

type Neg struct {
	base.NoOperandInstruction
}

func (n *Neg) Execute(frame *rt.Frame) {
	right := frame.PopStack()

	result := utils.CalcPrefix("-", right)

	if err, ok := result.(*object.ErrorObj); ok {
		panic(err.Msg)
	}

	frame.PushStack(result)
}
