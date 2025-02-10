package problem4

import (
	"errors"
)

type Element[T comparable] struct {
	value T
	next  *Element[T]
}

type LinkedList[T comparable] struct {
	head *Element[T]
}

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
}

func (ll *LinkedList[T]) Insert(el *Element[T], index int) error {
	if index > ll.Size() {
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
	return nil
}

func (ll *LinkedList[T]) Delete(el *Element[T]) error {
	if ll.head == nil {
		return errors.New("list is empty")
	}

	if ll.head.value == el.value {
		ll.head = ll.head.next
		return nil
	}

	current := ll.head
	for current.next != nil && current.next.value != el.value {
		current = current.next
	}

	if current.next == nil {
		return errors.New("element not found")
	}

	current.next = current.next.next
	return nil
}

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

func (ll *LinkedList[T]) List() []T {
	var elements []T
	current := ll.head
	for current != nil {
		elements = append(elements, current.value)
		current = current.next
	}
	return elements
}

func (ll *LinkedList[T]) Size() int {
	return len(ll.List())
}

func (ll *LinkedList[T]) IsEmpty() bool {
	return ll.Size() == 0
}
