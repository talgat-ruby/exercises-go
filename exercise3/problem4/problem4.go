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
	size int
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
	ll.size++
}

func (ll *LinkedList[T]) Insert(el *Element[T], pos int) error {
	if pos < 0 || pos > ll.size {
		return errors.New("index out of range")
	}
	if pos == 0 {
		el.next = ll.head
		ll.head = el
	} else {
		current := ll.head
		for i := 0; i < pos-1; i++ {
			current = current.next
		}
		el.next = current.next
		current.next = el
	}
	ll.size++
	return nil
}

func (ll *LinkedList[T]) Delete(el *Element[T]) error {
	if ll.head == nil {
		return errors.New("list is empty")
	}

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

	current.next = current.next.next
	ll.size--
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
	var result []T
	current := ll.head
	for current != nil {
		result = append(result, current.value)
		current = current.next
	}
	return result
}

func (ll *LinkedList[T]) Size() int {
	return ll.size
}

func (ll *LinkedList[T]) IsEmpty() bool {
	return ll.size == 0
}
