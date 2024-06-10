package problem4

import "errors"

type Element[T comparable] struct {
	value T
	next  *Element[T]
}

type LinkedList[T comparable] struct {
	head *Element[T]
	tail *Element[T]
	size int
}

func (list *LinkedList[T]) Add(el *Element[T]) {
	if list.head == nil {
		list.head = el
		list.tail = list.head
	} else {
		list.tail.next = el
		list.tail = el
	}
	list.size++
}

func (list *LinkedList[T]) Insert(el *Element[T], pos int) error {
	if el == nil {
		return errors.New("element is nil")
	}
	if pos < 0 || pos >= list.size {
		return errors.New("invalid position")
	}
	cur := list.head
	for i := 0; i < list.size; i++ {
		if pos == i {
			tmp := cur.next
			cur.next = el
			el.next = tmp
			list.size++
			return nil
		}
		cur = cur.next
	}
	return nil
}

func (list *LinkedList[T]) Delete(el *Element[T]) error {
	if el == nil {
		return errors.New("element is nil")
	}
	if list.size == 0 {
		return errors.New("list is empty")
	}
	cur := list.head
	if cur.next == nil && cur.value != el.value {
		return errors.New("element not found in list")
	} else if cur.next == nil && cur.value == el.value {
		list.head = nil
		list.tail = nil
		list.size--
		return nil
	} else if cur.next != nil && cur.value == el.value {
		list.head = cur.next
		list.size--
	}
	next := cur.next
	for i := 1; i < list.size; i++ {
		if next.value == el.value {
			cur.next = next.next
			list.size--
			return nil
		}
		cur = cur.next
		next = next.next
	}
	return errors.New("element not found")
}

func (list *LinkedList[T]) Find(v T) (Element[T], error) {
	for cur := list.head; cur != nil; cur = cur.next {
		if cur.value == v {
			return *cur, nil
		}
	}
	return Element[T]{}, errors.New("element not found")
}

func (list *LinkedList[T]) List() []T {
	out := make([]T, list.size)
	cur := list.head
	for i := 0; i < list.size; i++ {
		out[i] = cur.value
		cur = cur.next
	}
	return out
}

func (list *LinkedList[T]) Size() int {
	return list.size
}

func (list *LinkedList[T]) IsEmpty() bool {
	return list.size == 0
}
