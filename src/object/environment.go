package object

type Environment struct {
	store map[string]Object
	outer *Environment
}

func NewEnvironment() *Environment {
	store := make(map[string]Object)
	return &Environment{store: store, outer: nil}
}

func NewEnclosedEnvironment(env *Environment) *Environment {
	newEnv := NewEnvironment()
	newEnv.outer = env
	return newEnv
}

func (env *Environment) Get(name string) (Object, bool) {
	result, ok := env.store[name]
	if ok {
		return result, true
	}

	if env.outer != nil {
		return env.outer.Get(name)
	}

	return nil, false
}

func (env *Environment) ExistsCur(name string) (Object, bool) {
	result, ok := env.store[name]
	return result, ok
}

func (env *Environment) Set(name string, val Object) {
	env.store[name] = val
}

func (env *Environment) ExitsSet(name string, val Object) bool {
	_, ok := env.store[name]

	if ok {
		env.Set(name, val)
	}

	if env.outer != nil {
		env.outer.ExitsSet(name, val)
	}

	return false
}
