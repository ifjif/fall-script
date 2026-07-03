package utils

import (
	. "zzc/fall-script/src/object"
)

func UnknownBinaryOperationErr(left Object, right Object, op string) *ErrorObj {
	msg := binaryOperatorMsg(left, right, op)
	return NewError("Error(unknown operation): %s", msg)
}

func DivideByZeroErr(left Object, right Object, op string) *ErrorObj {
	msg := binaryOperatorMsg(left, right, op)
	return NewError("Error(divide by zero): %s", msg)
}

func UnknownUnaryOperationErr(op string, right Object) *ErrorObj {
	msg := unaryOperatorMsg(op, right)
	return NewError("Error(unknown operation): %s", msg)
}

func UnusableAsHashKeyErr(key Object) *ErrorObj {
	msg := objectMsg(key)
	return NewError("Error(unusable as hash key): %s", msg)
}

func IndexOutOfBoundErr(left, index Object) *ErrorObj {
	leftMsg := objectMsg(left)
	indexMsg := objectMsg(index)
	return NewError("Error(index out of bound): %s[%s]", leftMsg, indexMsg)
}

func UnsupportedIndexOperationErr(left, index Object) *ErrorObj {
	leftMsg := objectMsg(left)
	indexMsg := objectMsg(index)
	return NewError("Error(unsupported index operation): %s[%s]", leftMsg, indexMsg)
}

func UnsupportedIndexAssignOperationErr(left, index, value Object) *ErrorObj {
	leftMsg := objectMsg(left)
	indexMsg := objectMsg(index)
	valueMsg := objectMsg(value)
	return NewError("Error(unsupported index assign operation): %s[%s] = %s", leftMsg, indexMsg, valueMsg)
}

func UnsupportedAssignOperation(left, value Object) *ErrorObj {
	leftMsg := objectMsg(left)
	valueMsg := objectMsg(value)
	return NewError("Error(unsupported assign operation): %s = %s", leftMsg, valueMsg)
}

func IdentifierNotFoundErr(name string) *ErrorObj {
	return NewError("Error(identifier not found): %q", name)
}

func ConvertToIntegerErr(value Object) *ErrorObj {
	valueMsg := objectMsg(value)
	return NewError("Error(cannot convert to integer): %s", valueMsg)
}

func NotASliceErr(left Object) *ErrorObj {
	return NewError("Error(cannot be a slice): %s", left.Type())
}

func RedeclaredErr(name string) *ErrorObj {
	return NewError("Error(redeclared): %q", name)
}

func NotAFunctionErr(o Object) *ErrorObj {
	msg := o.Type()
	return NewError("Error(not a function): %s", msg)
}
