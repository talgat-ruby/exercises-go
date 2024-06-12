package problem4

import "fmt"

type Element[T any] struct {
	value T
	next  *Element[T]
}

type LinkedList[T any] struct {
	head   *Element[T]
	length int
}

// Add appends an element to the linked list.
func (ll *LinkedList[T]) Add(el *Element[T]) {
	el.next = ll.head
	ll.head = el
	ll.length++
}

// Size returns the number of elements in the linked list.
func (ll *LinkedList[T]) Size() int {
	return ll.length
}

// Insert adds an element at the specified position.
func (ll *LinkedList[T]) Insert(el *Element[T], pos int) error {
	if pos < 0 || pos > ll.length {
		return fmt.Errorf("position out of bounds")
	}
	if pos == 0 {
		ll.Add(el)
		return nil
	}
	current := ll.head
	for i := 1; i < pos; i++ {
		current = current.next
	}
	el.next = current.next
	current.next = el
	ll.length++
	return nil
}

// Delete removes an element from the linked list.
func (ll *LinkedList[T]) Delete(el *Element[T]) error {
	if ll.head == nil {
		return fmt.Errorf("empty list")
	}
	if ll.head.value == el.value {
		ll.head = ll.head.next
		ll.length--
		return nil
	}
	current := ll.head
	for current.next != nil && current.next.value != el.value {
		current = current.next
	}
	if current.next == nil {
		return fmt.Errorf("element not found")
	}
	current.next = current.next.next
	ll.length--
	return nil
}

// Find returns the element matching the given value.
func (ll *LinkedList[T]) Find(value T) (*Element[T], error) {
	current := ll.head
	for current != nil {
		if current.value == value {
			return current, nil
		}
		current = current.next
	}
	return nil, fmt.Errorf("element not found")
}

// List returns all the elements in the linked list as a slice.
func (ll *LinkedList[T]) List() []T {
	var elements []T
	current := ll.head
	for current != nil {
		elements = append(elements, current.value)
		current = current.next
	}
	return elements
}

// IsEmpty returns true if the linked list is empty.
func (ll *LinkedList[T]) IsEmpty() bool {
	return ll.length == 0
}
