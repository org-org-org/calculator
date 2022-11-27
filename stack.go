package calculator

type Stack struct {
	a []interface{}
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Len() int {
	return len(s.a)
}

func (s *Stack) Empty() bool {
	return s.Len() == 0
}

func (s *Stack) Push(v ...interface{}) {
	s.a = append(s.a, v...)
}

func (s *Stack) Pop() interface{} {
	if s.Empty() {
		return nil
	}
	v := s.a[s.Len()-1]
	s.a = s.a[:s.Len()-1]
	return v
}

func (s *Stack) Top() interface{} {
	if s.Empty() {
		return nil
	}
	return s.a[s.Len()-1]
}
