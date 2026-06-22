package object

import "fmt"

type Closure struct {
	Fn          *CompiledFunction
	Free        []Object
	OwnerModule *CompiledModule
}

func (c *Closure) Type() ObjectType {
	return CLOSURE_OBJ
}

func (c *Closure) Inspect() string {
	return fmt.Sprintf("Closure[%p]", c)
}

func (c *Closure) GetGlobal(idx int) Object {
	return c.OwnerModule.Globals[idx]
}

func (c *Closure) SetGlobal(idx int, o Object) {
	c.OwnerModule.Globals[idx] = o
}
