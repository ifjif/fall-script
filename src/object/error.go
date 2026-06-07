package object

import "fmt"

type ErrorObj struct {
	Msg string
}

func (e *ErrorObj) Type() ObjectType {
	return ERROR_OBJ
}

func (e *ErrorObj) Inspect() string {
	return e.Msg
}

func NewError(msg string, a ...any) *ErrorObj {
	return &ErrorObj{Msg: fmt.Sprintf(msg, a...)}
}
