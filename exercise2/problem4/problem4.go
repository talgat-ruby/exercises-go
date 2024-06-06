package problem4

import "errors"

type LinkedList[T comparable] struct {
	value T
	next  *LinkedList[T]
}
type Element[T comparable] struct {
	value T
}

func (l *LinkedList[T]) Add(el *Element[T]) {
	newNode := &LinkedList[T]{value: el.value}

	if l == nil {
		l.value = el.value
		return
	}

	current := l
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}
func (l *LinkedList[T]) Size() int {
	c := -1
	current := l
	for current != nil {
		current = current.next
		c++
	}
	return c
}
func (l *LinkedList[T]) Insert(el *Element[T], n int) error {
	llSize := l.Size()
	if llSize < n {
		return errors.New("ll < n")
	}
	cur := l
	var c int
	for c != n-1 {
		c++
		cur = cur.next
	}
	tail := cur.next
	newNode := &LinkedList[T]{el.value, tail}
	cur.next = newNode
	return nil
}
func (l *LinkedList[T]) Delete(el *Element[T]) error {
	if l == nil {
		return errors.New("nil ll")
	}
	cur := l
	for cur.next.value != el.value {
		if cur.next == nil {
			return errors.New("nil ll")
		}
		cur = cur.next
	}
	tail := cur.next.next
	cur.next = tail
	return nil
}
func (l *LinkedList[T]) Find(el T) (LinkedList[T], error) {
	if l == nil {
		return *l, errors.New("nil ll")
	}
	cur := l.next
	for cur != nil {
		if cur.value == el {
			return *cur, nil
		}
		cur = cur.next
	}
	return *l, errors.New("this data err")
}
func (l *LinkedList[T]) List() []T {
	var arr []T
	cur := l.next
	for cur != nil {
		arr = append(arr, cur.value)
		cur = cur.next
	}
	return arr
}
func (l *LinkedList[T]) IsEmpty() bool {
	if l.next == nil {
		return true
	}
	return false
}
