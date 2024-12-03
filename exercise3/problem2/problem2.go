package problem2

import "errors"

type Stack struct {
	items []interface{}
}

func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("Stack is empty")
	}
	lastIndex := s.Size() - 1
	item := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return item, nil
}

func (s *Stack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("Stack is empty")
	}
	return s.items[s.Size()-1], nil
}

func (s *Stack) Size() int {
	return len(s.items)
}

func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}
