package problem4

import (
	"errors"
	"strconv"
)

type Element[T int | string | bool] struct {
	value T
	next  *Element[T]
}

type LinkedList[T int | string | bool] struct {
	head *Element[T]
}

func (list *LinkedList[T]) Add(el *Element[T]) {
	if list.head == nil {
		list.head = el
	} else {
		current := list.head

		for current.next != nil {
			current = current.next
		}

		current.next = el
	}
}

func (list *LinkedList[T]) Size() int {
	size := 0
	current := list.head

	for current != nil {
		size++
		current = current.next
	}

	return size
}

func (list *LinkedList[T]) Insert(el *Element[T], position int) error {
	current := list.head
	currentPosition := 0

	for currentPosition != position {
		if current.next != nil {
			current = current.next
		}

		next := current.next
		current.next = el
		el.next = next

		if next == nil {
			list.head = el
		}

		currentPosition++

		if currentPosition == position {
			err := errors.New("Cannot insert element at position " + strconv.Itoa(position))

			return err
		}
	}

	return nil
}
