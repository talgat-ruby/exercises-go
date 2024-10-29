package problem4

import (
	"errors"
)

type Element[T comparable] struct {
	value T
	next  *Element[T]
}

type LinkedList[T comparable] struct {
	Head *Element[T]
	size int
}

func (l *LinkedList[T]) Add(n *Element[T]) {
	if l.Head == nil {
		l.Head = n
	} else {
		current := l.Head
		for current.next != nil {
			current = current.next
		}
		current.next = n
	}
	l.size++
}

func (l *LinkedList[T]) Size() int {
	return l.size
}

func (l *LinkedList[T]) IsEmpty() bool {
	return l.Head == nil
}

func (l *LinkedList[T]) Insert(n *Element[T], pos int) error {
	if pos < 0 || pos > l.size {
		return errors.New("Error")
	}

	if pos == 0 {
		n.next = l.Head
		l.Head = n
	} else {
		current := l.Head
		for i := 0; i < pos-1; i++ {
			current = current.next
		}
		n.next = current.next
		current.next = n
	}

	l.size++
	return nil
}

func (l *LinkedList[T]) Delete(n *Element[T]) error {
	if l.Head == nil {
		return errors.New("Empty")
	}

	if l.Head.value == n.value {
		l.Head = l.Head.next
		l.size--
		return nil
	}

	previous := l.Head
	current := l.Head.next

	for current != nil {

		if current.value == n.value {
			previous.next = current.next
			l.size--
			return nil
		}

		previous = current
		current = current.next
	}
	return errors.New("Error")
}

func (l *LinkedList[T]) Find(n any) (*Element[T], error) {

	if l.Head == nil {
		return nil, errors.New("Empty")
	}

	current := l.Head

	for current != nil {

		if current.value == n {

			return current, nil
		}

		current = current.next
	}

	return nil, errors.New("not find")
}

func (l *LinkedList[T]) List() []T {
	elements := []T{}

	current := l.Head

	for current != nil {
		elements = append(elements, current.value)
		current = current.next
	}

	return elements

}
