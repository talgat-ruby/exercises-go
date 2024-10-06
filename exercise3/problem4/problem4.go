package problem4

import "fmt"

type Element[T comparable] struct {
	value T
	next  *Element[T]
}

type LinkedList[T comparable] struct {
	head *Element[T]
	tail *Element[T]
	size int
}

func (ll *LinkedList[T]) Add(el *Element[T]) {
	if ll.head == nil {
		ll.head = el
		ll.tail = el
	} else {
		ll.tail.next = el
		ll.tail = el
	}
	ll.size++
}

func (ll *LinkedList[T]) Insert(el *Element[T], index int) error {
	if index < 0 || index > ll.size {
		return fmt.Errorf("index out of range")
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

func (ll *LinkedList[T]) Delete(el *Element[T]) error {
	if ll.head == nil {
		return fmt.Errorf("list is empty")
	}
	if ll.head.value == el.value {
		ll.head = ll.head.next
		ll.size--
		return nil
	}
	current := ll.head
	for current.next != nil {
		if current.next.value == el.value {
			current.next = current.next.next
			ll.size--
			return nil
		}
		current = current.next
	}
	return fmt.Errorf("element not found")
}

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

func (ll *LinkedList[T]) List() []T {
	result := make([]T, ll.size)
	current := ll.head
	for i := 0; i < ll.size; i++ {
		result[i] = current.value
		current = current.next
	}
	return result
}

func (ll *LinkedList[T]) IsEmpty() bool {
	return ll.size == 0
}

func (ll *LinkedList[T]) Size() int {
	return ll.size
}
