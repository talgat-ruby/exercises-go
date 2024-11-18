package problem4

import (
	"errors"
)

type Element[T comparable] struct {
	value T
	next  *Element[T]
	prev  *Element[T]
}

type LinkedList[T comparable] struct {
	head *Element[T]
}

func (list *LinkedList[T]) Add(el *Element[T]) {
	if list.head == nil {
		list.head = el
		return
	}

	current := list.head
	for current.next != nil {
		current = current.next
	}
	current.next = el
}

func (list *LinkedList[T]) Insert(el *Element[T], index int) error {
	if index < 0 || index > list.Size() {
		return errors.New("index out of range")
	}

	idx := 1
	current := list.head

	for current.next != nil {
		idx++
		if idx == index {
			temp := current.next
			current.next = el
			current.next.next = temp
			return nil
		}
		current = current.next
	}
	return nil
}

func (list *LinkedList[T]) Delete(el *Element[T]) error {
	if list.head.value == el.value {
		list.head = list.head.next
		return nil
	}

	if list.head == nil {
		return errors.New("list is empty")
	}

	current := list.head
	for current.next != nil {
		if current.next.value == el.value {
			current.next = current.next.next
			return nil
		}
		current = current.next
	}
	return errors.New("not found")
}

func (list *LinkedList[T]) Find(value any) (*Element[T], error) {
	if list.head.value == value {
		return list.head, nil
	}

	current := list.head
	for current.next != nil {
		if current.next.value == value {
			return current.next, nil
		}
		current = current.next
	}
	return &Element[T]{}, errors.New("not found")
}

func (list *LinkedList[T]) List() []T {
	var result []T
	if list.head == nil {
		return result
	}
	current := list.head
	result = append(result, current.value)
	for current.next != nil {
		current = current.next
		result = append(result, current.value)
	}
	return result
}

func (list *LinkedList[T]) Size() int {
	return len(list.List())
}

func (list *LinkedList[T]) IsEmpty() bool {
	return list.Size() == 0
}
