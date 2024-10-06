package problem4

import "errors"

type Element[T comparable] struct {
	value T
	next  *Element[T]
}

func (e *Element[T]) Equal(b *Element[T]) bool {
	return e.value == b.value
}

type LinkedList[T comparable] struct {
	head *Element[T]
}

func (l *LinkedList[T]) Add(value *Element[T]) {
	if l.head == nil {
		l.head = value
		return
	}

	head := l.head
	for head.next != nil {
		head = head.next
	}
	head.next = value
}

func (l *LinkedList[T]) Insert(value *Element[T], index int) error {
	if l.Size() <= index {
		return errors.New("index is greater than size of linked list")
	}
	head := l.head
	if index == 0 {
		temp := head
		value.next = temp
		head = value
		return nil
	}
	for range index - 1 {
		head = head.next
	}

	value.next = head.next
	head.next = value
	return nil
}

func (l *LinkedList[T]) Delete(e *Element[T]) error {
	if l.head.value == e.value {
		l.head = l.head.next
		return nil
	}
	head := l.head
	for head != nil {
		if head.next.value == e.value {
			head.next = head.next.next
			return nil
		}
		head = head.next
	}

	return errors.New("don't find an element")

}

func (l *LinkedList[T]) Find(value T) (*Element[T], error) {
	head := l.head
	for head != nil {
		if head.value == value {
			return head, nil
		}
		head = head.next
	}
	return nil, errors.New("don't find such a value")
}
func (l *LinkedList[T]) List() []T {
	values := make([]T, 0, l.Size())
	temp := l.head
	for temp != nil {
		values = append(values, temp.value)
		temp = temp.next
	}
	return values
}

func (l *LinkedList[T]) Size() int {
	cnt := 0
	temp := l.head
	for temp != nil {
		temp = temp.next
		cnt += 1
	}
	return cnt
}
func (l *LinkedList[T]) IsEmpty() bool {
	return l.Size() == 0
}
