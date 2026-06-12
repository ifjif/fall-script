package compiler

import (
	"zzc/fall-script/src/code"
	"zzc/fall-script/src/object"
)

type Scope struct {
	Constants     []object.Object
	Instructions  code.Instructions
	LastInst      code.EmittedInstruct
	PrevInst      code.EmittedInstruct
	StackDepth    int
	MaxStackDepth int
}

func NewScope() *Scope {
	return &Scope{
		Constants:    make([]object.Object, 0),
		Instructions: code.Instructions{},
	}
}
