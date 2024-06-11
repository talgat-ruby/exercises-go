package problem2

import "errors"

type Stack struct {
	stack []any
}

func (s *Stack) Push(val any) {
	s.stack = append([]any{val}, s.stack...)
}

func (s *Stack) Pop() (any, error) {
	if len(s.stack) == 0 {
		return nil, errors.New("stack is empty")
	}

	item := s.stack[0]
	s.stack = s.stack[1:]

	return item, nil
}

func (s *Stack) Peek() (any, error) {
	if len(s.stack) == 0 {
		return nil, errors.New("stack is empty")
	}

	return s.stack[0], nil
}

func (s *Stack) Size() int {
	return len(s.stack)
}

func (s *Stack) IsEmpty() bool {
	if len(s.stack) == 0 {
		return true
	}

	return false
}
