package evaluator

import (
	"fmt"

	. "zzc/fall-script/src/object"
)

func (e Evaluator) unknownBinaryOperationErr(left Object, right Object, op string) *ErrorObj {
	token := e.curNode.GetToken()
	msg := binaryOperatorMsg(left, right, op)
	return NewError("Error(unknown operation): %s at line %d, column %d", msg, token.Line, token.Col)
}

func (e Evaluator) divideByZeroErr(left Object, right Object, op string) *ErrorObj {
	token := e.curNode.GetToken()
	msg := binaryOperatorMsg(left, right, op)
	return NewError("Error(divide by zero): %s at line %d, column %d", msg, token.Line, token.Col)
}

func (e Evaluator) unknownUnaryOperationErr(op string, right Object) *ErrorObj {
	token := e.curNode.GetToken()
	msg := unaryOperatorMsg(op, right)
	return NewError("Error(unknown operation): %s at line %d, column %d", msg, token.Line, token.Col)
}

func (e Evaluator) unusableAsHashKeyErr(key Object) *ErrorObj {
	token := e.curNode.GetToken()
	msg := objectMsg(key)
	return NewError("Error(unusable as hash key): %s at line %d, column %d", msg, token.Line, token.Col)
}

func (e *Evaluator) indexOutOfBoundErr(left, index Object) *ErrorObj {
	token := e.curNode.GetToken()
	leftMsg := objectMsg(left)
	indexMsg := objectMsg(index)
	return NewError("Error(index out of bound): %s[%s] at line %d, column %d", leftMsg, indexMsg, token.Line, token.Col)
}

func (e *Evaluator) unsupportedIndexOperationErr(left, index Object) *ErrorObj {
	token := e.curNode.GetToken()
	leftMsg := objectMsg(left)
	indexMsg := objectMsg(index)
	return NewError("Error(unsupported index operation): %s[%s] at line %d, column %d", leftMsg, indexMsg, token.Line, token.Col)
}

func (e *Evaluator) identifierNotFoundErr(name string) *ErrorObj {
	token := e.curNode.GetToken()
	return NewError("Error(identifier not found): %q at line %d, column %d", name, token.Line, token.Col)
}

func (e *Evaluator) redeclaredErr(name string) *ErrorObj {
	token := e.curNode.GetToken()
	return NewError("Error(redeclared): %q at lint %d, column %d", name, token.Line, token.Col)
}

func (e *Evaluator) notAFunctionErr(o Object) *ErrorObj {
	token := e.curNode.GetToken()
	msg := o.Type()
	return NewError("Error(not a function): %s at line %d, column %d", msg, token.Line, token.Col)
}

func objectMsg(o Object) string {
	otype := o.Type()
	ovalue := o.Inspect()
	return fmt.Sprintf("%s(%s)", otype, ovalue)
}

func unaryOperatorMsg(op string, right Object) string {
	rightType := right.Type()
	rightValue := right.Inspect()
	return fmt.Sprintf("%s%s(%s)", op, rightType, rightValue)
}

func binaryOperatorMsg(left Object, right Object, op string) string {
	leftType := left.Type()
	leftValue := left.Inspect()
	rightType := right.Type()
	rightValue := right.Inspect()
	return fmt.Sprintf("%s(%s) %s %s(%s)", leftType, leftValue, op, rightType, rightValue)
}
