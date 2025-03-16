package main

import "fmt"

// Stack represents a simple stack structure with basic operations
type Stack[T any] struct {
	items []T
}

// Push adds an element to the top of the Stack
func (s *Stack[T]) Push(value T) {
	s.items = append(s.items, value)
}

// Pop removes and returns the element at the top of the Stack
// If the Stack is empty, it returns the zero value of the element type and false
func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}
	// Remove and return the last element
	value := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return value, true
}

// Peek returns the element at the top of the Stack without removing it
// If the Stack is empty, it returns the zero value of the element type and false
func (s *Stack[T]) Peek() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}
	return s.items[len(s.items)-1], true
}

// Size returns the number of elements in the Stack
func (s *Stack[T]) Size() int {
	return len(s.items)
}

// IsEmpty checks if the Stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Test the Stack data structure
func main() {
	stack := &Stack[int]{}

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	fmt.Println("Size:", stack.Size())       // Output: Size: 3
	fmt.Println("Peek:", stack.Peek())       // Output: Peek: 30
	fmt.Println("Pop:", stack.Pop())         // Output: Pop: 30
	fmt.Println("Size after pop:", stack.Size()) // Output: Size after pop: 2
	fmt.Println("IsEmpty:", stack.IsEmpty()) // Output: IsEmpty: false
	stack.Pop()
	stack.Pop()
	fmt.Println("IsEmpty after all pops:", stack.IsEmpty()) // Output: IsEmpty after all pops: true
}
