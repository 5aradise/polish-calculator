package poca

type stack[E any] struct {
	data []E
}

func (s *stack[E]) push(v E) {
	s.data = append(s.data, v)
}

func (s *stack[E]) pop() (E, bool) {
	var v E
	if len(s.data) < 1 {
		return v, false
	}

	v = s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v, true
}
