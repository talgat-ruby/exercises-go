package problem4

import "errors"

type LinkedList[T comparable] struct {
	head *Element[T]
	tail *Element[T]
	size int
}

type Element[T comparable] struct {
	value T
	next  *Element[T]
}

func (l *LinkedList[T]) Add(el *Element[T]) {
	if l.head == nil {
		l.head = el
		l.tail = el
		l.size++
		return
	}
	l.tail.next = el
	l.tail = el
	l.size++
}

func (l *LinkedList[T]) Insert(el *Element[T], ind int) error {
	if el == nil {
		return errors.New("element is nil")
	}
	if l.size <= ind {
		return errors.New("index out of range")
	}
	l.size++
	if ind == 0 {
		el.next = l.head
		l.head = el
		return nil
	}

	cur := l.head
	for range ind - 1 {
		cur = cur.next
	}
	el.next = cur.next
	cur.next = el
	return nil
}

func (l *LinkedList[T]) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkedList[T]) Delete(el *Element[T]) error {
	if el == nil {
		return errors.New("element is nil")
	}
	cur := l.head
	if cur.value == el.value {
		l.head = l.head.next
		l.size--
		return nil
	}
	for i := 1; i < l.size; i++ {
		prev := cur
		cur = cur.next
		if prev.value == el.value {
			prev.next = cur.next
			l.size--
			return nil
		}
	}

	return errors.New("element not found")
}

func (l *LinkedList[T]) Find(val T) (*Element[T], error) {
	cur := l.head
	for i := 0; i < l.size; i++ {
		if cur.value == val {
			return cur, nil
		}
		cur = cur.next
	}
	return nil, errors.New("element not found")
}

func (l *LinkedList[T]) List() []*Element[T] {
	out := make([]*Element[T], l.size)
	cur := l.head
	for i := 0; i < l.size; i++ {
		out[i] = cur
		cur = cur.next
	}
	return out
}

func (l *LinkedList[T]) Size() int {
	return l.size
}
