package reference

import (
	"zzc/fall-script/src/instructions/base"
	"zzc/fall-script/src/object"
	"zzc/fall-script/src/utils"
	"zzc/fall-script/src/vm/rt"
)

type Index struct {
	base.NoOperandInstruction
}

func (i *Index) Execute(frame *rt.Frame) {
	index := frame.PopStack()
	left := frame.PopStack()

	result := utils.CalcIndex(left, index)

	if err, ok := result.(*object.ErrorObj); ok {
		panic(err.Msg)
	}

	frame.PushStack(result)
}
