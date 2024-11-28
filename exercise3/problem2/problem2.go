package problem2

import "errors"

type Stack struct {
	data []any
}

func (s *Stack) Push(value any) {
	s.data = append(s.data, value)
}
func (s *Stack) Pop() (any, error) {
	if len(s.data) == 0 {
		return nil, errors.New("stack is empty")
	}
	top := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return top, nil
}
func (s *Stack) Peek() (any, error) {
	if len(s.data) == 0 {
		return nil, errors.New("stack is empty")
	}
	return s.data[len(s.data)-1], nil
}
func (s *Stack) Size() int {
	return len(s.data)
}
func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}
