package instructions

import (
	"fmt"

	"zzc/fall-script/src/code"
	"zzc/fall-script/src/instructions/base"
	. "zzc/fall-script/src/instructions/comparison"
	. "zzc/fall-script/src/instructions/consts"
	. "zzc/fall-script/src/instructions/control"
	. "zzc/fall-script/src/instructions/funct"
	. "zzc/fall-script/src/instructions/load"
	. "zzc/fall-script/src/instructions/logic"
	. "zzc/fall-script/src/instructions/math"
	. "zzc/fall-script/src/instructions/reference"
	. "zzc/fall-script/src/instructions/stack"
	. "zzc/fall-script/src/instructions/store"
)

var (
	nop         = &Nop{}
	null        = &Null{}
	pop         = &Pop{}
	gt          = &Gt{}
	ge          = &Ge{}
	add         = &Add{}
	sub         = &Sub{}
	mul         = &Mul{}
	div         = &Div{}
	eq          = &Eq{}
	neq         = &Neq{}
	neg         = &Neg{}
	not         = &Not{}
	true_       = &True{}
	false_      = &False{}
	index       = &Index{}
	set_index   = &SetIndex{}
	dup         = &Dup{}
	cur_closure = &CurClosure{}
	ret         = &Return{}
	xret        = &XReturn{}
)

func NewInstruction(opcode byte) base.Instruction {
	switch code.OpCode(opcode) {
	case code.Nop:
		return nop
	case code.Null_:
		return null
	case code.Const:
		return &Const{}
	case code.Pop:
		return pop
	case code.Gt:
		return gt
	case code.Ge:
		return ge
	case code.Add:
		return add
	case code.Sub:
		return sub
	case code.Mul:
		return mul
	case code.Div:
		return div
	case code.Eq:
		return eq
	case code.Neq:
		return neq
	case code.Neg:
		return neg
	case code.Not:
		return not
	case code.True:
		return true_
	case code.False:
		return false_
	case code.Array_:
		return &Array{}
	case code.Hash_:
		return &Hash{}
	case code.Index:
		return index
	case code.Call:
		return &Call{}
	case code.Jump:
		return &Jump{}
	case code.JumpIsFalse:
		return &JumpIsFalse{}
	case code.SetGlobal:
		return &SetGlobal{}
	case code.GetGlobal:
		return &GetGlobal{}
	case code.SetLocal:
		return &SetLocal{}
	case code.GetLocal:
		return &GetLocal{}
	case code.NewBoxLocal:
		return &NewBoxLocal{}
	case code.SetBoxLocal:
		return &SetBoxLocal{}
	case code.GetBoxLocal:
		return &GetBoxLocal{}
	case code.SetFree:
		return &SetFree{}
	case code.GetFree:
		return &GetFree{}
	case code.GetFreeRaw:
		return &GetFreeRaw{}
	case code.GetBuiltin:
		return &GetBuiltin{}
	case code.SetIndex:
		return set_index
	case code.Closure_:
		return &Closure{}
	case code.CurClosure:
		return cur_closure
	case code.Dup:
		return dup
	case code.Return:
		return ret
	case code.XReturn:
		return xret
	}

	panic(fmt.Sprintf("Error: unsupported opcode 0x%03x!", opcode))
}
