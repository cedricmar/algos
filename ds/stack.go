package ds

type Stacker interface {
	Push(e ...int)
	Pop() int
	Length() int
}

type Stack []int

func (s *Stack) Push(e ...int) {
	*s = append(*s, e...)
}

// Pop returns -1 if the stack is empty
func (s *Stack) Pop() int {
	l := s.Length() - 1
	if l < 0 {
		return -1
	}
	v := (*s)[l]
	*s = (*s)[:l]
	return v
}

func (s Stack) Length() int {
	return len(s)
}
