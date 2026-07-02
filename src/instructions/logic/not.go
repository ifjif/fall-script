package logic

import (
	"zzc/fall-script/src/instructions/base"
	"zzc/fall-script/src/object"
	"zzc/fall-script/src/utils"
	"zzc/fall-script/src/vm/rt"
)

type Not struct {
	base.NoOperandInstruction
}

func (n *Not) Execute(frame *rt.Frame) {
	right := frame.PopStack()

	result := utils.CalcPrefix("!", right)

	if err, ok := result.(*object.ErrorObj); ok {
		panic(err.Msg)
	}

	frame.PushStack(result)
}
