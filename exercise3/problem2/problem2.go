package problem2

import (
	"errors"
)

type Stack struct {
	elements []interface{}
}

func (s *Stack) Push(element interface{}) {
	s.elements = append(s.elements, element)
}

func (s *Stack) Pop() (interface{}, error) {
	if len(s.elements) == 0 {
		return nil, errors.New("Stack is empty")
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element, nil
}

func (s *Stack) Peek() (interface{}, error) {
	if len(s.elements) == 0 {
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
