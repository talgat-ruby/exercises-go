package problem2

import "errors"

type Stack struct {
	elements []any
}

func (s *Stack) Push(newElement any) {
	s.elements = append(s.elements, newElement)
}

func (s *Stack) Pop() (any, error) {
	if len(s.elements) == 0 {
		return nil, errors.New("stack is empty")
	}
	lastElement := s.elements[len(s.elements)-1]
	s.elements = s.elements[0 : len(s.elements)-1]
	return lastElement, nil
}

func (s *Stack) Peek() (any, error) {
	if len(s.elements) == 0 {
		return nil, errors.New("stack is empty")
	}
	return s.elements[len(s.elements)-1], nil
}

func (s *Stack) Size() any {
	return len(s.elements)
}

func (s *Stack) IsEmpty() any {
	return len(s.elements) == 0
}
