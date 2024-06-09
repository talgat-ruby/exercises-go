package problem2

import "errors"

type Stack struct {
	items []any
}

func (s *Stack) Push(item any) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() (any, error) {
	if len(s.items) == 0 {
		err := errors.New("stack is empty. Nothing to pop")

		return nil, err
	}

	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return item, nil
}

func (s *Stack) Size() int {
	return len(s.items)
}

func (s *Stack) Peek() (any, error) {
	if len(s.items) == 0 {
		err := errors.New("stack is empty. Nothing to peek")

		return nil, err
	}

	return s.items[len(s.items)-1], nil
}

func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}
