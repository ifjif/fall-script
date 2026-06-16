package binarychunck

/*
*
	Constants     []object.Object
	Instructions  code.Instructions
	StackDepth    int
	MaxStackDepth int

	signature
	cons number
	type tag bnum data

	integer
	string
	cf {
	Constants     []object.Object
	Instructions  code.Instructions
	StackDepth    int
	MaxStackDepth int
	}
*
*/

const (
	SIGNATURE = "fallscript"
	MAJOR     = 0
	MINOR     = 1
	PATCH     = 0
)

type TypeTag byte

const (
	_   byte = iota
	I64      = iota
	STR      = iota
	CF       = iota
)
