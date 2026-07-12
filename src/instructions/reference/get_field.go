package reference

import (
	"errors"

	"zzc/fall-script/src/instructions/base"
	"zzc/fall-script/src/object"
	"zzc/fall-script/src/vm/rt"
)

type GetField struct {
	base.NoOperandInstruction
}

func (gf *GetField) Execute(frame *rt.Frame) {
	member := frame.PopStack()
	visitor := frame.PopStack()

	si, ok := visitor.(*object.StructInstance)
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

	// 判断是否是 embed
	if fi.IsEmbed {
		baseOffset := fi.Index

		structMeta, err := base.GetTargetStructMeta(frame, si.StructMeta, fi.TargetStructMeta)
		if errors.Is(err, base.ErrAwaitModuleInitialization) {
			return
		}
		if err != nil {
			panic(err)
		}

		nsi := &object.StructInstance{
			Slots:      si.Slots[baseOffset:],
			StructMeta: structMeta,
		}
		frame.PushStack(nsi)
		return
	}

	result := si.Slots[fi.Index]

	frame.PushStack(result)
}
