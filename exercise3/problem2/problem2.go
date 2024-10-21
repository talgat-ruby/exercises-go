package problem2

import (
	"errors"
)

type Stack struct {
	items []any
}

func (s *Stack) Push(val any) {
	s.items = append(s.items, val)
}

func (s *Stack) Pop() (any, error) {
	if len(s.items) == 0 {
		return nil, errors.New("stack is empty")
	}
	element := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return element, nil
}

func (s *Stack) Peek() (any, error) {
	if len(s.items) == 0 {
		return nil, errors.New("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

func (s *Stack) Size() int {
	return len(s.items)
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}
