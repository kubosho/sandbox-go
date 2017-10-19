package stack

type Stack struct {
	data   []string
	length int
}

func (s *Stack) Pop() string {
	if s.length == 0 {
		return ""
	}
	s.length--
	return s.data[s.length]
}

func (s *Stack) Push(str string) {
	s.data = append(s.data, str)
	s.length++
}
