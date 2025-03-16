package problem4

import (
	"errors"
)

type LinkedList[T int | bool | string] struct {
	head *Element[T]
}

type Element[T int | bool | string] struct {
	next  *Element[T]
	value T
}

func (l *LinkedList[T]) Add(item *Element[T]) {
	if l.head == nil {
		l.head = item
	} else {
		cur := l.head
		for cur.next != nil {
			cur = cur.next
		}

		cur.next = item
	}
}

func (l *LinkedList[T]) Insert(item *Element[T], position int) error {
	if position > l.Size()+1 {
		return errors.New("incorrect position")
	}
	if position == 0 {
		item.next = l.head
		l.head = item
	} else {
		cur := l.head
		for i := 0; i < position-1; i++ {
			cur = cur.next
		}
		item.next = cur.next
		cur.next = item
	}
	return nil
}

func (l *LinkedList[T]) Delete(item *Element[T]) error {
	if l.IsEmpty() {
		return errors.New("empty list")
	}

	if l.head.value == item.value {
		l.head = l.head.next
		return nil
	}

	cur := l.head.next
	prev := l.head
	for cur.next != nil {
		if cur.value == item.value {
			prev.next = cur.next
			return nil
		}
		prev = cur
		cur = cur.next
	}

	if cur.next == nil {
		return errors.New("not found")
	}

	return nil
}

func (l *LinkedList[T]) Find(value T) (*Element[T], error) {
	cur := l.head
	for cur.next != nil {
		if cur.value == value {
			return cur, nil
		}
		cur = cur.next
	}

	return nil, errors.New("not found")
}

func (l *LinkedList[T]) List() []T {
	sl := make([]T, 0)
	cur := l.head
	for cur != nil {
		sl = append(sl, cur.value)
		cur = cur.next
	}
	return sl
}

func (l *LinkedList[T]) IsEmpty() bool {
	return l.Size() == 0
}

func (l *LinkedList[T]) Size() int {

	return len(l.List())
}
