package problem4

import "errors"

type Element[T any] struct {
	value T
	next  *Element[T]
}

type LinkedList[T any] struct {
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
		return errors.New("position out of bounds")
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

func (ll *LinkedList[T]) Delete(el *Element[T], compare func(a, b T) bool) error {
	if ll.head == nil {
		return errors.New("list is empty")
	}
	if compare(ll.head.value, el.value) {
		ll.head = ll.head.next
		ll.size--
		return nil
	}
	current := ll.head
	for current.next != nil && !compare(current.next.value, el.value) {
		current = current.next
	}
	if current.next == nil {
		return errors.New("element not found")
	}
	current.next = current.next.next
	ll.size--
	return nil
}

func (ll *LinkedList[T]) Find(value T, compare func(a, b T) bool) (*Element[T], error) {
	current := ll.head
	for current != nil {
		if compare(current.value, value) {
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
	return ll.size
}

func (ll *LinkedList[T]) IsEmpty() bool {
	return ll.size == 0
}
