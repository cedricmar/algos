package ds

type Stack[T any] []T

func (s *Stack[T]) Push(e ...T) {
	*s = append(*s, e...)
}

// Pop returns false if the stack is empty
func (s *Stack[T]) Pop() (T, bool) {
	l := s.Length() - 1
	if l < 0 {
		var v T
		return v, false
	}
	v := (*s)[l]
	*s = (*s)[:l]
	return v, true
}

func (s Stack[T]) Length() int {
	return len(s)
}
