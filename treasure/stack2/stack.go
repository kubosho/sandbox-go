package stack

type Stack struct {
	data  []string
	limit int
}

func (s *Stack) Pop() string {
	if len(s.data) == 0 {
		return ""
	}

	d := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return d
}

func (s *Stack) Push(str string) {
	if len(s.data) == s.limit {
		s.data = s.data[1:]
	}
	s.data = append(s.data, str)
}
