package heap

type Prototype struct {
	Bytecode  []byte
	ConstPool *ConstPool
}

func NewPrototype(bytecode []byte, constPool *ConstPool) *Prototype {
	return &Prototype{
		Bytecode:  bytecode,
		ConstPool: constPool,
	}
}
