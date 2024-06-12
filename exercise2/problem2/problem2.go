package problem2

import "errors"

type Stack struct {
	data []any
}

func (s *Stack) Push(v any) {
	s.data = append(s.data, v)
}

func (s *Stack) Pop() (any, error) {
	if len(s.data) == 0 {
		return nil, errors.New("Stack is empty")
	}
	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v, nil
}

func (s *Stack) Size() int {
	return len(s.data)
}

func (s *Stack) Peek() (any, error) {
	if len(s.data) == 0 {
		return nil, errors.New("Stack is empty")
	}
	return s.data[len(s.data)-1], nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}
