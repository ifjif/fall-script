package evaluator

import (
	. "zzc/fall-script/src/object"
	"zzc/fall-script/src/utils"
)

func (e *Evaluator) applyFunction(fn Object, args []Object) Object {
	switch fn := fn.(type) {
	case *Function:
		curEnv := e.env
		env := extendFunctionEnv(fn, args)
		e.env = env
		result := unwrapReturnValue(e.eval(fn.Body))
		e.env = curEnv
		return result
	case *Builtin:
		return fn.Fn(args...)
	}

	return e.appendLineAndCol(utils.NotAFunctionErr(fn))
}

func extendFunctionEnv(fn *Function, args []Object) *Environment {
	env := NewEnclosedEnvironment(fn.Env)
	params := fn.Params

	for i, param := range params {
		env.Set(param.Value, args[i])
	}

	return env
}
