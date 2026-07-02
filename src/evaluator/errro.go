package evaluator

import (
	"fmt"

	"zzc/fall-script/src/object"
	. "zzc/fall-script/src/object"
)

func (e Evaluator) appendLineAndCol(err *object.ErrorObj) *ErrorObj {
	token := e.curNode.GetToken()
	msg := fmt.Sprintf("at lint %d, column %d", token.Line, token.Col)
	err.Msg = fmt.Sprintf("%s %s", err.Msg, msg)
	return err
}
