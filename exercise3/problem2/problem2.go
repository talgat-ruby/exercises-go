package problem2

import "errors"

type Stack struct {
	elements []any
}

func (s *Stack) Push(value any) {
	s.elements = append(s.elements, value)
}

func (s *Stack) Pop() (any, error) {
	if s.IsEmpty() {
		return nil, errors.New("Stack is empty")
	}

	index := len(s.elements) - 1
	val := s.elements[index]
	s.elements = s.elements[:index]
	return val, nil
}

func (s *Stack) Peek() (any, error) {
	if s.IsEmpty() {
		return nil, errors.New("Stack is empty")
	}
	return s.elements[len(s.elements)-1], nil
}

func (s *Stack) Size() int {
	return len(s.elements)
}

func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}
