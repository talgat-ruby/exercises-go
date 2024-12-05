package problem2

import "errors"

type Stack struct {
	slc []interface{}
}

func (s *Stack) Push(value interface{}) {
	s.slc = append(s.slc, value)
}

func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("Пумпурум, пусто")
	}
	lastValue := s.slc[len(s.slc)-1]
	s.slc = s.slc[:len(s.slc)-1]
	return lastValue, nil
}

func (s *Stack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("Пумпурум, пусто")
	}
	return s.slc[len(s.slc)-1], nil
}

func (s *Stack) Size() int {
	return len(s.slc)
}

func (s *Stack) IsEmpty() bool {
	return len(s.slc) == 0
}
