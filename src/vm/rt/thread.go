package rt

type Thread struct {
	Pc    int
	Stack *Stack
}

func NewThread() *Thread {
	return &Thread{
		Pc:    0,
		Stack: NewStack(),
	}
}

func (t *Thread) CurrentFrame() *Frame {
	return t.Stack.top
}
