package problem2

import "errors"

type Stack struct {
	Arr []any
}

func (s *Stack) Push(elem any) {
	newArr := []any{elem}
	newArr = append(newArr, s.Arr...)
	s.Arr = newArr
}
func (s *Stack) Pop() (any, error) {
	if len(s.Arr) == 0 {
		return nil, errors.New("empty stack")
	}
	elem := s.Arr[0]
	s.Arr = s.Arr[1:]
	return elem, nil
}
func (s *Stack) Peek() (any, error) {
	if len(s.Arr) == 0 {
		return nil, errors.New("empty stack")
	}
	elem := s.Arr[0]
	return elem, nil
}
func (s *Stack) Size() int {
	return len(s.Arr)
}
func (s *Stack) IsEmpty() bool {
	if len(s.Arr) == 0 {
		return true
	}
	return false
}
