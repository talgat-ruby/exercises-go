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

func (l *LinkedList[T]) Add(el *Element[T]) {
	if l.head == nil {
		l.head = el
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = el
	}
	l.size++
}

func (l *LinkedList[T]) Insert(el *Element[T], position int) error {

	if position < 0 || position > l.size {
		return errors.New("position is out of range")
	}

	if position == 0 {
		el.next = l.head
		l.head = el
		l.size++
		return nil
	}

	if position == l.size {
		l.Add(el)
		return nil
	}

	current := l.head
	count := 0
	for current != nil {
		if count == position-1 {
			el.next = current.next
			current.next = el
			l.size++
			return nil
		}
		current = current.next
		count++
	}
	return errors.New("element not found")
}

func (l *LinkedList[T]) Delete(el *Element[T]) error {
	if l.head == nil {
		return errors.New("list is empty")
	}

	if l.head.value == el.value {
		l.head = l.head.next
		l.size--
		return nil
	}

	current := l.head
	var prev *Element[T]

	for current != nil {
		if current.value == el.value {
			prev.next = current.next
			l.size--
			return nil
		}
		prev = current
		current = current.next
	}

	return errors.New("element not found")
}

func (l *LinkedList[T]) Find(el any) (*Element[T], error) {

	if l.head == nil {
		return nil, errors.New("empty")
	}

	current := l.head

	for current != nil {

		if current.value == el {

			return current, nil
		}

		current = current.next
	}

	return nil, errors.New("element not found")
}

func (l *LinkedList[T]) List() []T {
	list := make([]T, 0, l.size)
	current := l.head
	for current != nil {
		list = append(list, current.value)
		current = current.next
	}
	return list
}

func (l *LinkedList[T]) Size() int {
	return l.size
}

func (l *LinkedList[T]) IsEmpty() bool {
	return l.head == nil
}
