package problem2

import "errors"

type Stack struct {
	items []any
}

func (s *Stack) Push(item any) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() (any, error) {
	last, err := s.Peek()
	if err == nil {
		s.items = s.items[:s.Size()-1]
	}
	return last, err
}

func (s *Stack) Peek() (any, error) {
	if s.IsEmpty() {
		return nil, errors.New("Stack empty")
	}
	return s.items[len(s.items)-1], nil
}

func (s *Stack) Size() int {
	return len(s.items)
}

func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}
