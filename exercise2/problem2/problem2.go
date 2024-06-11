package problem2

import "errors"

type Stack struct {
	Elements []any
}

func (s *Stack) Push(val any) {
	s.Elements = append(s.Elements, val)
}

func (s *Stack) Pop() (any, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	popped := s.Elements[s.Size()-1]
	s.Elements = s.Elements[:(s.Size() - 1)]
	return popped, nil
}

func (s *Stack) Peek() (any, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	return s.Elements[s.Size()-1], nil
}

func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Stack) Size() int {
	return len(s.Elements)
}
