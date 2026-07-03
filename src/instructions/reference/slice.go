package reference

import (
	"zzc/fall-script/src/instructions/base"
	"zzc/fall-script/src/object"
	"zzc/fall-script/src/utils"
	"zzc/fall-script/src/vm/rt"
)

type Slice struct {
	base.NoOperandInstruction
}

func (s *Slice) Execute(frame *rt.Frame) {
	info := frame.PopStacks(4)
	container := frame.PopStack()

	result := utils.CalcSlice(container, info)
	if err, ok := result.(*object.ErrorObj); ok {
		panic(err.Msg)
	}

	frame.PushStack(result)
}
