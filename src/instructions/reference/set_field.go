package reference

import (
	"zzc/fall-script/src/instructions/base"
	"zzc/fall-script/src/object"
	"zzc/fall-script/src/vm/rt"
)

type SetField struct {
	base.NoOperandInstruction
}

func (sf *SetField) Execute(frame *rt.Frame) {
	value := frame.PopStack()
	member := frame.PopStack()
	owner := frame.PopStack()

	si, ok := owner.(*object.StructInstance)
	if !ok {
		panic("not struct instance")
	}
	name, ok := member.(*object.String)
	if !ok {
		panic("only supported string name")
	}

	fi, ok := si.StructMeta.Fields[name.Value]

	if !ok {
		panic("not exist field xx")
	}

	si.Slots[fi.Index] = value
	frame.PushStack(value)
}
