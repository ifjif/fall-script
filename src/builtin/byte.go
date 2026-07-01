package builtin

import "zzc/fall-script/src/object"

func Byte(args ...object.Object) object.Object {
	length := len(args)

	if length != 1 {
		panic("Error: byte() expected 1 argument")
	}

	o := args[0]

	switch o.Type() {
	case object.INTEGER_OBJ:
		v := o.(*object.Integer)
		return &object.Byte{Value: byte(v.Value)}
	case object.STRING_OBJ:
		v := o.(*object.String)
		if len(v.Value) != 1 {
			panic("Error: Expected string length is 1")
		}
		return &object.Byte{Value: v.Value[0]}
	}

	panic("Error: byte() argument must be integer or single string")
}
