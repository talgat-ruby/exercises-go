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

func (ll *LinkedList[T]) Add(element *Element[T]) {
	if ll.head == nil {
		ll.head = element
	} else {
		current := ll.head
		for current.next != nil {
			current = current.next
		}
		current.next = element
	}
	ll.size++
}

func (ll *LinkedList[T]) Insert(element *Element[T], position int) error {
	if position < 0 || position > ll.size {
		return errors.New("position out of range")
	}

	if position == 0 {
		element.next = ll.head
		ll.head = element
	} else {
		current := ll.head
		for i := 0; i < position-1; i++ {
			current = current.next
		}
		element.next = current.next
		current.next = element
	}
	ll.size++
	return nil
}

func (ll *LinkedList[T]) Delete(element *Element[T]) error {
	if ll.head == nil {
		return errors.New("list is empty")
	}

	if ll.head == element {
		ll.head = ll.head.next
	} else {
		current := ll.head
		for current.next != nil && current.next != nil {
			current = current.next
		}
		if current.next == nil {
			return errors.New("element not found")
		}
		current.next = current.next.next
	}
	ll.size--
	return nil

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

// Size gets the size of the linked list.
func (ll *LinkedList[T]) Size() int {
	return ll.size
}

// IsEmpty checks if the linked list is empty.
func (ll *LinkedList[T]) IsEmpty() bool {
	return ll.size == 0
}
