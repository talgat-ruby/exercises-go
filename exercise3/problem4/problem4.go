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
	tail *Element[T]
	size int
}

func (l *LinkedList[T]) Add(e *Element[T]) {
	if l.head == nil {
		l.head = e
		l.tail = e
	} else {
		l.tail.next = e
		l.tail = e
	}
	l.size++
}

func (l *LinkedList[T]) Size() int {
	return l.size
}

func (l *LinkedList[T]) Insert(e *Element[T], index int) error {
	if index < 0 || index > l.size {
		return errors.New("index out of range")
	}

	if index == 0 {
		e.next = l.head
		l.head = e
		if l.size == 0 {
			l.tail = e
		}
	} else {
		prev, err := l.findPreviousByIndex(index - 1)
		if err != nil {
			return err
		}
		e.next = prev.next
		prev.next = e
		if e.next == nil {
			l.tail = e
		}
	}
	l.size++
	return nil
}

func (l *LinkedList[T]) Delete(e *Element[T]) error {
	if l.IsEmpty() {
		return errors.New("linkedlist is empty")
	}

	if l.head.value == e.value {
		l.head = l.head.next
		if l.head == nil {
			l.tail = nil
		}
	} else {
		prev, err := l.findPrevious(e.value)
		if err != nil {
			return err
		}
		prev.next = prev.next.next
		if prev.next == nil {
			l.tail = prev
		}
	}
	l.size--
	return nil
}

func (l *LinkedList[T]) Find(value T) (*Element[T], error) {
	current := l.head
	for current != nil {
		if current.value == value {
			return current, nil
		}
		current = current.next
	}
	return nil, errors.New("value not found")
}

func (l *LinkedList[T]) List() []T {
	result := make([]T, 0, l.size)
	for current := l.head; current != nil; current = current.next {
		result = append(result, current.value)
	}
	return result
}

func (l *LinkedList[T]) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkedList[T]) findPreviousByIndex(index int) (*Element[T], error) {
	if index < 0 || index >= l.size {
		return nil, errors.New("index out of range")
	}

	current := l.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	return current, nil
}

func (l *LinkedList[T]) findPrevious(value T) (*Element[T], error) {
	if l.head == nil || l.head.value == value {
		return nil, errors.New("value not found")
	}

	current := l.head
	for current.next != nil && current.next.value != value {
		current = current.next
	}

	if current.next == nil {
		return nil, errors.New("value not found")
	}
	return current, nil
}

func (l *LinkedList[T]) Reverse() {
	var prev, next *Element[T]
	current := l.head
	l.tail = current

	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}

	l.head = prev
}
