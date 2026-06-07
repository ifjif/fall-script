package rt

type Stack struct {
	top *Frame
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) PushFrame(frame *Frame) {
	frame.Lower = s.top
	s.top = frame
}
