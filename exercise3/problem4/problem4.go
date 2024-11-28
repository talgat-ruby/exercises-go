package main

import "fmt"

// Node represents a single element in the linked list
type Node[T any] struct {
	value T
	next  *Node[T]
}

// LinkedList represents a linked list structure with various operations
type LinkedList[T any] struct {
	head *Node[T]
	size int
}

// Add appends a new element to the end of the linked list
func (ll *LinkedList[T]) Add(value T) {
	newNode := &Node[T]{value: value}
	if ll.head == nil {
		ll.head = newNode
	} else {
		current := ll.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
	ll.size++
}

// Insert inserts a new element at a specific position in the linked list
func (ll *LinkedList[T]) Insert(value T, position int) bool {
	if position < 0 || position > ll.size {
		return false
	}
	newNode := &Node[T]{value: value}
	if position == 0 {
		newNode.next = ll.head
		ll.head = newNode
	} else {
		current := ll.head
		for i := 0; i < position-1; i++ {
			current = current.next
		}
		newNode.next = current.next
		current.next = newNode
	}
	ll.size++
	return true
}

// Delete removes the first occurrence of a specific element from the linked list
func (ll *LinkedList[T]) Delete(value T) bool {
	if ll.head == nil {
		return false
	}
	if ll.head.value == value {
		ll.head = ll.head.next
		ll.size--
		return true
	}
	current := ll.head
	for current.next != nil && current.next.value != value {
		current = current.next
	}
	if current.next == nil {
		return false
	}
	current.next = current.next.next
	ll.size--
	return true
}

// Find searches for an element and returns its pointer if found, or nil otherwise
func (ll *LinkedList[T]) Find(value T) *Node[T] {
	current := ll.head
	for current != nil {
		if current.value == value {
			return current
		}
		current = current.next
	}
	return nil
}

// List returns a slice containing all elements in the linked list
func (ll *LinkedList[T]) List() []T {
	var elements []T
	current := ll.head
	for current != nil {
		elements = append(elements, current.value)
		current = current.next
	}
	return elements
}

// Size returns the number of elements in the linked list
func (ll *LinkedList[T]) Size() int {
	return ll.size
}

// IsEmpty checks if the linked list is empty
func (ll *LinkedList[T]) IsEmpty() bool {
	return ll.size == 0
}

// Test the LinkedList data structure
func main() {
	ll := &LinkedList[int]{}
	ll.Add(10)
	ll.Add(20)
	ll.Add(30)

	fmt.Println("List:", ll.List())          // Output: [10, 20, 30]
	fmt.Println("Size:", ll.Size())          // Output: 3
	fmt.Println("IsEmpty:", ll.IsEmpty())    // Output: false

	ll.Insert(15, 1)                         // Insert 15 at position 1
	fmt.Println("After Insert:", ll.List())  // Output: [10, 15, 20, 30]

	ll.Delete(20)                            // Delete value 20
	fmt.Println("After Delete:", ll.List())  // Output: [10, 15, 30]

	fmt.Println("Find 15:", ll.Find(15))     // Output: &{15 <next node>}
	fmt.Println("Find 99:", ll.Find(99))     // Output: <nil>
}
