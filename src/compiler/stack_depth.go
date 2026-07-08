package compiler

import . "zzc/fall-script/src/code"

func (c *Compiler) getStackDelta(op OpCode, operands []int) int {
	operand := 0
	if len(operands) > 0 {
		operand = operands[0]
	}
	switch op {
	case Nop:
		return 0
	case Null_:
		return +1
	case Const:
		return +1
	case Pop:
		return -1
	case Gt:
		return -1
	case Ge:
		return -1
	case Add:
		return -1
	case Sub:
		return -1
	case Mul:
		return -1
	case Div:
		return -1
	case Eq:
		return -1
	case Neq:
		return -1
	case Neg:
		return 0
	case Not:
		return 0
	case True:
		return +1
	case False:
		return +1
	case Array_:
		return -(operand - 1)
	case Hash_:
		return -(2*operand - 1)
	case Index:
		return -1
	case Slice:
		return -4
	case Call:
		return -operand
	case Jump:
		return 0
	case JumpIsFalse:
		return -1
	case SetGlobal:
		return -1
	case GetGlobal:
		return +1
	case SetLocal:
		return -1
	case GetLocal:
		return +1
	case NewBoxLocal:
		return -1
	case SetBoxLocal:
		return -1
	case GetBoxLocal:
		return +1
	case SetFree:
		return -1
	case GetFree:
		return +1
	case GetFreeRaw:
		return +1
	case GetBuiltin:
		return +1
	case SetIndex:
		return -2
	case Closure_:
		operand = operands[1]
		return -(operand - 1)
	case CurClosure:
		return +1
	case Dup:
		return +1
	case InitStruct:
		return -(2 * operand)
	case GetField:
		return -1
	case SetField:
		return -2
	case CallMethod:
		return -(operand + 2)
	case Return:
		return 0
	case XReturn:
		return -1
	}

	return 0
}
