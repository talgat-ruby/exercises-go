package problem4

import "errors"

type Element[T comparable] struct {
	value T
	next  *Element[T]
}

func NewElement[T comparable](data T) *Element[T] {
	return &Element[T]{value: data}
}

type LinkedList[T comparable] struct {
	head *Element[T]
	size int
}

func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (l *LinkedList[T]) Size() int {
	return l.size
}

func (l *LinkedList[T]) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkedList[T]) List() []T {
	if l.head == nil {
		return nil
	}
	output := make([]T, 0, l.size)
	current := l.head
	for current != nil {
		output = append(output, current.value)
		current = current.next
	}
	return output
}

func (l *LinkedList[T]) Add(el *Element[T]) {
	if el == nil {
		return
	}

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
func (l *LinkedList[T]) Delete(el *Element[T]) bool {
	if el == nil || l.head == nil {
		return false
	}

	if l.head.value == el.value {
		l.head = l.head.next
		l.size--
		return true
	}

	cur := l.head
	for cur.next != nil {
		if cur.next.value == el.value {
			cur.next = cur.next.next
			l.size--
			return true
		}
		cur = cur.next
	}
	return false
}

func (l *LinkedList[T]) Find(val T) (*Element[T], error) {
	if l.head == nil {
		return nil, errors.New("list is empty")
	}

	current := l.head
	for current != nil {
		if current.value == val {
			return current, nil
		}
		current = current.next
	}
	return nil, errors.New("element not found")
}

func (l *LinkedList[T]) Insert(el *Element[T], pos int) error {
	if pos < 0 || pos > l.size {
		return errors.New("position is out of range")
	}

	if pos == 0 {
		el.next = l.head
		l.head = el
		l.size++
		return nil
	}

	if pos == l.size {
		l.Add(el)
		return nil
	}

	current := l.head
	count := 0
	for current != nil {
		if count == pos-1 {
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
