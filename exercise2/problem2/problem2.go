package problem2

import "errors"

type Stack struct {
	length int
	pos    int
	array  []any
}

func (s *Stack) Push(val any) {
	if s.pos < s.length {
		s.array[s.pos] = val
	} else {
		s.array = append(s.array[:s.pos], val)
		s.length++
	}
	s.pos++
}

func (s *Stack) Pop() (any, error) {
	if s.pos > 0 {
		s.pos--
		return s.array[s.pos], nil
	}
	return nil, errors.New("stack is empty")
}

func (s *Stack) Peek() (any, error) {
	if s.pos > 0 {
		return s.array[s.pos-1], nil
	}
	return nil, errors.New("stack is empty")
}

func (s *Stack) Size() int {
	return s.pos
}

func (s *Stack) IsEmpty() bool {
	return s.pos == 0
}
