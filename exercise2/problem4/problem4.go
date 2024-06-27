package problem4

import "errors"

type Element[E comparable] struct {
	prev  *Element[E]
	value E
	next  *Element[E]
}

type LinkedList[E comparable] struct {
	first *Element[E]
	last  *Element[E]
}

func (l *LinkedList[E]) Add(e *Element[E]) {
	e.prev = l.last
	e.next = nil
	l.last = e

	if e.prev == nil {
		l.first = e
	} else {
		e.prev.next = e
	}
}

func (l *LinkedList[E]) Size() int {
	size := 0
	e := l.first
	for e != nil {
		size++
		e = e.next
	}

	return size
}

func (l *LinkedList[E]) Insert(e *Element[E], pos int) error {
	size := l.Size()
	if size < pos {
		return errors.New("there is no bathroom")
	}

	if pos == 0 || pos == size {
		l.Add(e)
		return nil
	}

	index := 1
	prev := l.first
	for index < pos {
		index++
		prev = prev.next
	}

	next := prev.next
	prev.next = e
	next.prev = e
	e.prev = prev
	e.next = next
	return nil
}

func (l *LinkedList[E]) Delete(e *Element[E]) error {
	current := l.first
	for current != nil && current.value != e.value {
		current = current.next
	}

	if current != nil {
		if current.prev == nil {
			l.RemoveHead()
		} else if current.next == nil {
			l.RemoveTail()
		} else {
			prev := current.prev
			next := current.next
			prev.next = next
			next.prev = prev
		}
	}

	return nil
}

func (l *LinkedList[E]) RemoveHead() {
	if l.first.next == nil {
		l.first = nil
		l.last = nil
	} else {
		l.first = l.first.next
		l.first.prev = nil
	}
}

func (l *LinkedList[E]) RemoveTail() {
	if l.last.prev == nil {
		l.first = nil
		l.last = nil
	} else {
		l.last = l.last.prev
		l.last.next = nil
	}
}

// Find and return an element from the linked list.
func (l *LinkedList[E]) Find(e E) (*Element[E], error) {
	current := l.first
	for current != nil {
		if current.value == e {
			return current, nil
		}
		current = current.next
	}

	return nil, nil
}

// Get the slice out of the linked list.
func (l *LinkedList[E]) List() []E {
	if l.first == nil {
		return make([]E, 0)
	}

	ans := make([]E, 1)
	ans[0] = l.first.value

	current := l.first
	for current.next != nil {
		current = current.next
		ans = append(ans, current.value)
	}

	return ans
}

// Check if the linked list is empty.
func (l *LinkedList[E]) IsEmpty() bool {
	return l.first == nil
}
