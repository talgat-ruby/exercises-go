package problem2

import "fmt"

type Stack struct {
	values []any
}

func (s *Stack) Push(any2 any) {
	dS := make([]any, 0)
	dS = append(dS, any2)
	s.values = append(dS, s.values...)
}

func (s *Stack) Pop() (any, error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf("empty")
	}
	v := s.values[0]
	s.values = s.values[1:]
	return v, nil
}

func (s *Stack) Peek() (any, error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf("empty")
	}
	return s.values[0], nil
}

func (s *Stack) Size() int {
	return len(s.values)
}

func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}
