package rt

type Stack struct {
	top     *Frame
	maxSize int
	size    int
}

func NewStack(maxSize int) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}

func (s *Stack) PushFrame(f *Frame) {
	s.isOutOfStack()
	f.SetLower(s.top)
	s.top = f
	s.size++
}

func (s *Stack) PopFrame() *Frame {
	if s.top == nil {
		return nil
	}

	cur := s.top
	s.top = s.top.lower
	if cur.lower != nil {
		cur.lower = nil
	}
	s.size--

	return cur
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) isOutOfStack() {
	if s.size >= s.maxSize {
		panic("out of stack")
	}
}
