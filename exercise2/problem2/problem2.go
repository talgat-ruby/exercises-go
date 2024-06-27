package problem2

import "errors"

type Stack struct {
	elements []any
}

func (s *Stack) Push(x any) {
	s.elements = append(s.elements, x)
}

func (s *Stack) Size() int {
	return len(s.elements)
}

func (s *Stack) Pop() (any, error) {
	if len(s.elements) == 0 {
		return nil, errors.New("stack is empty")
	}

	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]

	return element, nil
}

func (s *Stack) Peek() (any, error) {
	if len(s.elements) == 0 {
		return nil, errors.New("stack is empty")
	}

	return s.elements[len(s.elements)-1], nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}
