package reference

import (
	"errors"

	"zzc/fall-script/src/instructions/base"
	"zzc/fall-script/src/object"
	"zzc/fall-script/src/vm/rt"
)

type CallMethod struct {
	args int
}

func (ce *CallMethod) FetchOperand(br *base.ByteReader) {
	count := br.ReadUint8()
	ce.args = int(count)
}

func (ce *CallMethod) Execute(frame *rt.Frame) {
	args := frame.PopStacks(ce.args)
	member := frame.PopStack()
	visitor := frame.PopStack()

	arg1 := visitor
	si, ok := visitor.(*object.StructInstance)
	if !ok {
		panic("not struct instance")
	}
	name, ok := member.(*object.String)
	if !ok {
		panic("only supported string name")
	}

	targetStructMeta := si.StructMeta
	mr, ok := targetStructMeta.Methods[name.Value]

	if !ok {
		panic("not exist method xx")
	}

	var fn object.Object
	if mr.TargetStructMeta != -1 {
		structMeta, err := base.GetTargetStructMeta(frame, si.StructMeta, mr.TargetStructMeta)
		if errors.Is(err, base.ErrAwaitModuleInitialization) {
			return
		}
		if err != nil {
			panic(err)
		}

		sname := structMeta.Name
		// 找 offset
		field, ok := si.StructMeta.Fields[sname]
		if !ok {
			panic("not exist field, cannot call field method")
		}
		offset := field.Index
		offsetInstance := &object.StructInstance{
			Slots:      si.Slots[offset:],
			StructMeta: structMeta,
		}
		arg1 = offsetInstance
		mr = structMeta.Methods[name.Value]
		targetStructMeta = structMeta
	}

	fn = targetStructMeta.OwnerModule.Globals[mr.Index]

	closure, ok := fn.(*object.Closure)
	if !ok {
		panic("not method xxxx")
	}

	thread := frame.Thread()
	newFrame := thread.NewFrame(closure)

	newFrame.SetLocal(0, arg1)
	for i, arg := range args {
		newFrame.SetLocal(i+1, arg)
	}
	thread.PushFrame(newFrame)
}
