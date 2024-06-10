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
		return nil, errors.New("queue is empty")
	}
	val := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	return val, nil
}

func (s *Stack) Top() any {
	if len(s.vals) == 0 {
		return nil
	}
	return s.vals[len(s.vals)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.vals) == 0
}

func (s *Stack) Peek() (any, error) {
	if len(s.vals) == 0 {
		return nil, errors.New("stack is empty")
	}
	return s.vals[len(s.vals)-1], nil
}

func (s *Stack) EmptyStack() any {
	return nil
}

func (s *Stack) Size() int {
	return len(s.vals)
}
