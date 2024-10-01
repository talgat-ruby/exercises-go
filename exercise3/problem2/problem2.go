package problem2

import "errors"

type Stack struct {
	Elements []any
}

func (s *Stack) Push(element any) {
	s.Elements = append(s.Elements, element)
}

func (s *Stack) Pop() (any, error) {
	if len(s.Elements) == 0 {
		return 0, errors.New("Error")
	}
	element := s.Elements[len(s.Elements)-1]
	s.Elements = s.Elements[:len(s.Elements)-1]
	return element, nil
}

func (s *Stack)Peek() (any, error) {
	if len(s.Elements) == 0 {
		return 0, errors.New("Error")
	}

	return s.Elements[len(s.Elements)-1], nil
}

func (s *Stack)Size() int {
	return len(s.Elements)
}

func (s *Stack)IsEmpty() bool {
	return len(s.Elements) == 0
}

