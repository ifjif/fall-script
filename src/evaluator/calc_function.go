package evaluator

import . "zzc/fall-script/src/object"

func (e *Evaluator) applyFunction(fn *Function, args []Object) Object {
	curEnv := e.env
	env := extendFunctionEnv(fn, args)
	e.env = env
	result := unwrapReturnValue(e.eval(fn.Body))
	e.env = curEnv
	return result
}

func extendFunctionEnv(fn *Function, args []Object) *Environment {
	env := NewEnclosedEnvironment(fn.Env)
	params := fn.Params

	for i, param := range params {
		env.Set(param.Value, args[i])
	}

	return env
}
