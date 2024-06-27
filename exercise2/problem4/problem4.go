package problem4

import "errors"

type Element[T comparable] struct {
	value T
	next  *Element[T]
}
type LinkedList[T comparable] struct {
	element *Element[T]
}

func (ll *LinkedList[T]) Add(element *Element[T]) {
	newNode := &Element[T]{value: element.value}

	if ll.element == nil {
		ll.element = newNode
	} else {
		current := ll.element

		for current.next != nil {
			current = current.next
		}

		current.next = newNode
	}
}

func (ll *LinkedList[T]) Size() int {
	size := 0
	current := ll.element

	for current != nil {
		size++
		current = current.next
	}

	return size
}

func (ll *LinkedList[T]) Insert(element *Element[T], index int) error {
	if index == 0 {
		element.next = ll.element
		ll.element = element

		return nil
	}

	current := ll.element

	for i := 0; i < index; i++ {
		if current == nil || current.next == nil && i < index-1 {
			return errors.New("index out of bounds")
		}

		current = current.next
	}

	element.next = current.next
	current.next = element

	return nil
}

func (ll *LinkedList[T]) Delete(element *Element[T]) error {
	if ll.element == nil {
		return errors.New("element is nil")
	}

	if element.value == ll.element.value {
		ll.element = ll.element.next
		return nil
	}

	current := ll.element

	for current.next != nil && current.next.value != element.value {
		current = current.next
	}

	if current.next == nil {
		return errors.New("element not found")
	}

	current.next = current.next.next

	return nil
}

func (ll *LinkedList[T]) Find(value T) (*Element[T], error) {
	current := ll.element

	for current != nil {
		if current.value == value {
			return current, nil
		}

		current = current.next
	}

	return nil, errors.New("element not found")
}

func (ll *LinkedList[T]) List() []any {
	current := ll.element
	result := make([]any, 0)

	for current != nil {
		result = append(result, current.value)
		current = current.next
	}

	return result
}

func (ll *LinkedList[T]) IsEmpty() bool {
	return ll.Size() == 0
}
