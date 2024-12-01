package problem4

import (
	"errors"
)

type LinkedList[T comparable] struct {
	head *Element[T]
	size int
}

type Element[T comparable] struct {
	value T
	next  *Element[T]
}

// Add adds a new element at the end of the list.
func (ll *LinkedList[T]) Add(el *Element[T]) {
	if ll.head == nil {
		ll.head = el
	} else {
		current := ll.head
		for current.next != nil {
			current = current.next
		}
		current.next = el
	}
	ll.size++
}

// Insert inserts a new element at the specified index.
func (ll *LinkedList[T]) Insert(el *Element[T], index int) error {
	if index < 0 || index > ll.size {
		return errors.New("index out of bounds")
	}
	if index == 0 {
		el.next = ll.head
		ll.head = el
	} else {
		current := ll.head
		for i := 0; i < index-1; i++ {
			current = current.next
		}
		el.next = current.next
		current.next = el
	}
	ll.size++
	return nil
}

// Delete removes an element from the list.
func (ll *LinkedList[T]) Delete(el *Element[T]) error {
	if ll.head == nil {
		return errors.New("list is empty")
	}

	// If the element to delete is the head
	if ll.head.value == el.value {
		ll.head = ll.head.next
		ll.size--
		return nil
	}

	current := ll.head
	for current.next != nil && current.next.value != el.value {
		current = current.next
	}

	if current.next == nil {
		return errors.New("element not found")
	}

	// Remove the element
	current.next = current.next.next
	ll.size--
	return nil
}

// Find returns the element with the given value.
func (ll *LinkedList[T]) Find(value T) (*Element[T], error) {
	current := ll.head
	for current != nil {
		if current.value == value {
			return current, nil
		}
		current = current.next
	}
	return nil, errors.New("element not found")
}

// List returns a slice containing all the elements in the list.
func (ll *LinkedList[T]) List() []T {
	result := make([]T, 0, ll.size)
	current := ll.head
	for current != nil {
		result = append(result, current.value)
		current = current.next
	}
	return result
}

// Size returns the number of elements in the list.
func (ll *LinkedList[T]) Size() int {
	return ll.size
}

// IsEmpty returns true if the list is empty.
func (ll *LinkedList[T]) IsEmpty() bool {
	return ll.size == 0
}
