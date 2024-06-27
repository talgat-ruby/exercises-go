package problem2

import "errors"

type Stack struct{
	linkedList []any
}


func (s *Stack) Push(value any) {
	s.linkedList = append(s.linkedList, value)
}

func(s *Stack) Pop() (any, error) {
	if s.IsEmpty() {
		return nil, errors.New("linked list is empty")
	}

	del, newLL := s.linkedList[s.Size() - 1], s.linkedList[:s.Size() - 1]
	s.linkedList = newLL

	return del, nil

}

func(s *Stack) Size() int {
	size := len(s.linkedList)

	return size
}

func(s *Stack) Peek() (any, error) {
	if s.IsEmpty() {
		return nil, errors.New("linked list is empty")
	}

	del := s.linkedList[s.Size() - 1]

	return del, nil
}

func(s *Stack) IsEmpty() bool{
	return s.Size() == 0
}