package problem2

import "errors"

type Stack struct {
	stack []any
}

func (s *Stack) Push(val any) {
	s.stack = append(s.stack, val)
}

func (s *Stack) Pop() (any, error) {
	if len(s.stack) == 0 {
		return nil, errors.New("stack is empty")
	}
	a := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return a, nil
}
func (s *Stack) Peek() (any, error) {
	if len(s.stack) == 0 {
		return nil, errors.New("stack is empty")
	}
	return s.stack[len(s.stack)-1], nil
}

func (s *Stack) Size() int {
	return len(s.stack)
}

func (s *Stack) IsEmpty() bool {
	return len(s.stack) == 0
}
