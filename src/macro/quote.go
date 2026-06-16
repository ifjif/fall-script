package macro

import (
	"zzc/fall-script/src/ast"
	"zzc/fall-script/src/evaluator"
	"zzc/fall-script/src/object"
)

func quote(eval *evaluator.Evaluator, node ast.Node) object.Object {
	quoted := evalUnquote(eval, node)
	return &object.Quote{Node: quoted}
}
