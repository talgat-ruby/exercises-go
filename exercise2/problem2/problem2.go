package problem2

import "errors"

type Node struct {
	value any
	next  *Node
}
type Stack struct {
	top  *Node
	size int
}

func (s *Stack) Push(value any) {
	newNode := &Node{value: value, next: s.top}
	s.top = newNode
	s.size++
}

func (s *Stack) Pop() (any, error) {
	if s.size == 0 {
		return nil, errors.New("stack is empty")
	}
	value := s.top.value
	s.top = s.top.next
	s.size--
	return value, nil
}

func (s *Stack) Peek() (any, error) {
	if s.size == 0 {
		return nil, errors.New("stack is empty")
	}
	return s.top.value, nil
}
func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}
