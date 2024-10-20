package problem2

import "errors"

type Stack struct {
	arr []any
}

func (s *Stack) Push(a any) {
	s.arr = append(s.arr, a)
}

func (s *Stack) Pop() (any, error) {
	last, err := s.Peek()
	if err == nil {
		s.arr = s.arr[:s.Size()-1]
	}
	return last, err
}

func (s *Stack) Peek() (any, error) {
	if s.IsEmpty() {
		return nil, errors.New("Stack empty")
	}
	return s.arr[len(s.arr)-1], nil
}

func (s *Stack) Size() int {
	return len(s.arr)
}

func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}
