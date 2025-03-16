package problem2

import "errors"

type Stack struct {
	elements []any
}

func (s *Stack) Push(el any) {
	s.elements = append(s.elements, el)
}

func (s *Stack) Pop() (any, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	elements := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return elements, nil
}

func (s *Stack) Peek() (any, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	return s.elements[len(s.elements)-1], nil
}

func (s *Stack) Size() int {
	return len(s.elements)
}

func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}
