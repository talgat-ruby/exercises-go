package problem2

import "errors"

type Stack struct {
	slc []any
}

func (s *Stack) Push(x any) {
	s.slc = append(s.slc, x)
}

func (s *Stack) Pop() (any, error) {
	if len(s.slc) == 0 {
		return nil, errors.New("stack is empty")
	}
	val := s.slc[len(s.slc)-1]
	s.slc = s.slc[:len(s.slc)-1]
	return val, nil
}

func (s *Stack) Peek() (any, error) {
	if len(s.slc) == 0 {
		return nil, errors.New("stack is empty")
	}
	return s.slc[len(s.slc)-1], nil
}
func (s *Stack) Size() int {
	return len(s.slc)
}
func (s *Stack) IsEmpty() bool {
	return len(s.slc) == 0
}
