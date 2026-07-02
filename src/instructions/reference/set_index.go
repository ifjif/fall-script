package reference

import (
	"zzc/fall-script/src/instructions/base"
	"zzc/fall-script/src/object"
	"zzc/fall-script/src/utils"
	"zzc/fall-script/src/vm/rt"
)

type SetIndex struct {
	base.NoOperandInstruction
}

func (si *SetIndex) Execute(frame *rt.Frame) {
	v := frame.PopStack()
	i := frame.PopStack()
	c := frame.PopStack()

	result := utils.Assign4Index(c, i, v)

	if err, ok := result.(*object.ErrorObj); ok {
		panic(err.Msg)
	}

	frame.PushStack(result)
}
