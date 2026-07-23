package reference

import (
	"zzc/fall-script/src/instructions/base"
	"zzc/fall-script/src/object"
	"zzc/fall-script/src/vm/rt"
)

type InitStruct struct {
	index int
}

func (is *InitStruct) FetchOperand(br *base.ByteReader) {
	idx := br.ReadUint16()
	is.index = int(idx)
}

// [struct][key][value]
func (is *InitStruct) Execute(frame *rt.Frame) {
	pairs := frame.PopStacks(2 * is.index)
	meta := frame.PopStack()
	structMeta, ok := meta.(*object.StructMeta)
	if !ok {
		panic("not struct meta")
	}

	fields := structMeta.Fields
	count := structMeta.FieldCount
	slots := make([]object.Object, count)
	for i := range slots {
		slots[i] = object.NULL
	}
	instance := &object.StructInstance{
		StructMeta: structMeta,
		Slots:      slots,
	}

	for i := 0; i < len(pairs); i += 2 {
		key := pairs[i]
		value := pairs[i+1]
		name, ok := key.(*object.String)
		if !ok {
			panic("key is not string")
		}

		field, ok := fields[name.Value]
		if !ok {
			panic("field not found")
		}

		// 判断是否是 embed
		if field.IsEmbed {
			vi, ok := value.(*object.StructInstance)
			if !ok {
				panic("not struct instance to field xxxx")
			}

			baseOffset := field.Index
			for i, v := range vi.Slots {
				slots[baseOffset+i] = v
			}
		} else {
			slots[field.Index] = value
		}
	}

	frame.PushStack(instance)
}
