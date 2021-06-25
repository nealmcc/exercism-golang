package brackets

type stack struct {
	data []bracket
}

func (s *stack) len() int {
	return len(s.data)
}

func (s *stack) push(b bracket) {
	s.data = append(s.data, b)
}

func (s *stack) pop() (bracket, bool) {
	if len(s.data) == 0 {
		return 0, false
	}
	top := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return top, true
}
