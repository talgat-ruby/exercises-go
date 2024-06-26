package problem4

import "errors"

type Element[T comparable] struct {
	value T
	next  *Element[T]
}

type LinkedList[T comparable] struct {
	head *Element[T]
	tail *Element[T]
	size int
}

func (ll *LinkedList[T]) Add(element *Element[T]) {
	if ll.head == nil {
		ll.head = element
		ll.tail = ll.head
	} else {
		ll.tail.next = element
		ll.tail = element
	}
	ll.size++
}

func (ll *LinkedList[T]) Insert(element *Element[T], position int) error {
	if element == nil {
		return errors.New("the element is nil")
	}

	if position < 0 || position > ll.size {
		return errors.New("the position is invalid")
	}

	current := ll.head
	for i := 0; i < ll.size; i++ {
		if i == position {
			temp := current.next
			current.next = element
			element.next = temp
			ll.size++
			return nil
		}
	}
	return nil
}

func (ll *LinkedList[T]) Delete(element *Element[T]) error {
	if element == nil {
		return errors.New("the element is nil")
	}

	if ll.size == 0 {
		return errors.New("the list is empty")
	}

	current := ll.head
	switch {
	case current.next == nil && current.value != element.value:
		return errors.New("the element was not found")
	case current.next == nil && current.value == element.value:
		ll.head = nil
		ll.tail = nil
		ll.size--
		return nil
	case current.next != nil && current.value == element.value:
		ll.head = current.next
		ll.size--
	}

	next := current.next
	for i := 1; i < ll.size; i++ {
		if next.value == element.value {
			current.next = next.next
			ll.size--
			return nil
		}
		current = current.next
		next = next.next
	}
	return errors.New("the element was not found")

}

func (ll *LinkedList[T]) List() []T {
	res := make([]T, ll.size)
	current := ll.head
	for i := 0; i < ll.size; i++ {
		res[i] = current.value
		current = current.next
	}
	return res
}

// Size gets the size of the linked list.
func (ll *LinkedList[T]) Size() int {
	return ll.size
}

// IsEmpty checks if the linked list is empty.
func (ll *LinkedList[T]) IsEmpty() bool {
	return ll.size == 0
}

func (ll *LinkedList[T]) Find(value T) (Element[T], error) {
	for current := ll.head; current != nil; current = current.next {
		if current.value == value {
			return *current, nil
		}
	}
	return Element[T]{}, errors.New("the element was not found")
}
