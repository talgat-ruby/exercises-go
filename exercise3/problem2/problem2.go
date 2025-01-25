package problem2

import "errors"

type Stack struct {
	data []any
}

func (s *Stack) Push(elem any) {
	s.data = append(s.data, elem)
}

func (s *Stack) Size() int {
	return len(s.data)
}
func (s *Stack) Pop() (any, error) {
	if s.data == nil {
		return nil, errors.New("err")
	}
	last := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return last, nil
}
func (s *Stack) Peek() (any, error) {
	if s.data == nil {
		return nil, errors.New("err")
	}
	return s.data[len(s.data)-1], nil
}
func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}
