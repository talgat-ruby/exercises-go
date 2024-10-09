package problem2

import "errors"

type Stack struct {
	arr  []any
	size int
}

func (s *Stack) Push(val any) {
	s.arr = append(s.arr, val)
	s.size++
}

func (s *Stack) Pop() (any, error) {
	if s.size <= 0 {
		return nil, errors.New("stack is empty")
	}
	val := s.arr[s.size-1]
	s.arr = s.arr[:s.size-1]
	s.size--
	return val, nil
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) Peek() (any, error) {
	if s.size <= 0 {
		return nil, errors.New("stack is empty")
	}
	return s.arr[s.size-1], nil
}
