package problem2

import "errors"

type Stack struct {
	items []any
	size  int
}

func (s *Stack) Push(item any) {
	s.items = append(s.items, item)
	s.size++
}

func (s *Stack) Pop() (any, error) {
	if s.size == 0 {
		return nil, errors.New("stack is empty")
	}
	poppedItem := s.items[s.size-1]
	s.items = s.items[:s.size-1]
	s.size--
	return poppedItem, nil
}

func (s *Stack) Peek() (any, error) {
	if s.size == 0 {
		return nil, errors.New("stack is empty")
	}
	return s.items[s.size-1], nil
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}
