package semantic

import (
	"fmt"

	"zzc/fall-script/src/object"
	"zzc/fall-script/src/token"
)

func appendPosInfo(token token.Token, err *object.ErrorObj) string {
	msg := fmt.Sprintf("%s at line %d, column %d", err.Msg, token.Line, token.Col)
	return msg
}
