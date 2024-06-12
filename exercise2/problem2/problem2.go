package problem2

import "errors"

type Stack struct {
	values []any
}

func (s *Stack) Push(e any) {
	s.values = append(s.values, e)
}

func (s *Stack) Pop() (any, error) {
	if len(s.values) == 0 {
		return nil, errors.New("stack is empty")
	}
	var last = s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return last, nil
}

func (s *Stack) Peek() (any, error) {
	if len(s.values) == 0 {
		return nil, errors.New("stack is empty")
	}
	var last = s.values[len(s.values)-1]
	return last, nil
}

func (s *Stack) Size() int {
	return len(s.values)
}

func (s *Stack) IsEmpty() bool {
	return len(s.values) == 0
}
