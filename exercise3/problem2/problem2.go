package problem2

import "errors"

type Stack struct {
	values []any
}

func (s *Stack) Push(value any) {
	s.values = append(s.values, value)
}

func (s *Stack) Pop() (any, error) {
	if s.IsEmpty() {
		return nil, errors.New("Stack is empty")
	}
	valueToRemove := s.values[s.Size()-1]
	s.values = s.values[:s.Size()-1]
	return valueToRemove, nil
}

func (s *Stack) Peek() (any, error) {
	if s.IsEmpty() {
		return nil, errors.New("Stack is empty")
	}
	return s.values[s.Size()-1], nil
}

func (s *Stack) Size() int {
	return len(s.values)
}

func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}
