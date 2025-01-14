package problem2

import "errors"

type Stack struct {
	vals []any
}

func (s *Stack) Push(val any) {
	s.vals = append(s.vals, val)
}

func (s *Stack) Pop() (any, error) {
	if len(s.vals) == 0 {
		return nil, errors.New("stack is empty")
	}
	topElement := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	return topElement, nil
}

func (s *Stack) Peek() (any, error) {
	if len(s.vals) == 0 {
		return nil, errors.New("stack is empty")
	}
	return s.vals[len(s.vals)-1], nil
}

func (s *Stack) Size() int {
	return len(s.vals)
}

func (s *Stack) IsEmpty() bool {
	return len(s.vals) == 0
}
