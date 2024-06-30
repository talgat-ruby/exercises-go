package problem2

import (
	"fmt"
)

type Node struct {
	Data any
	Tail *Node
}

type Stack struct {
	head *Node
}

func (s *Stack) Push(v any) {
	currentHead := s.head
	s.head = &Node{
		Data: v,
		Tail: currentHead,
	}
}

func (s *Stack) Pop() (any, error) {
	if s.head == nil {
		return nil, fmt.Errorf("stack is empty")
	}

	v := s.head.Data
	s.head = s.head.Tail
	return v, nil
}

func (s *Stack) Peek() (any, error) {
	if s.head == nil {
		return nil, fmt.Errorf("stack is empty")
	}

	v := s.head.Data
	return v, nil
}

func (s *Stack) Size() int {
	size := 0
	node := s.head
	for node != nil {
		node = node.Tail
		size++
	}
	return size
}

func (s *Stack) IsEmpty() bool {
	return s.head == nil
}
