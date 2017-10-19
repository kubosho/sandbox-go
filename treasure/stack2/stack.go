package stack

type Stack struct {
	data   []string
	length int
	limit  int
}

func (s *Stack) Pop() string {
	if s.length == 0 {
		return ""
	}
	s.length--
	return s.data[s.length]
}

func (s *Stack) Push(str string) {
	if s.length >= s.limit {
		s.data = append(s.data, str)
		s.data = s.data[1:]
		return
	}
	s.data = append(s.data, str)
	s.length++
}
