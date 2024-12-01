package problem2

import "errors"

type Stack struct {
	Items []any
}

func (s *Stack) Push(value any) {
	s.Items = append(s.Items, value)
}

func (s *Stack) Pop() (any, error) {
	if s.IsEmpty() {
		return 0, errors.New("error")
	}

	index := len(s.Items) - 1
	element := s.Items[index]
	s.Items = s.Items[:index]
	return element, nil
}

func (s *Stack) Peek() (any, error) {
	if s.IsEmpty() {
		return 0, errors.New("error")
	}
	return s.Items[len(s.Items)-1], nil
}

func (s *Stack) Size() int {
	return len(s.Items)
}

func (s *Stack) IsEmpty() bool {
	return len(s.Items) == 0
}
