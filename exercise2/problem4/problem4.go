package problem4

import "fmt"

type Element[V comparable] struct {
	value V
	next  *Element[V]
}

type LinkedList[V comparable] struct {
	head *Element[V]
}

func (ll *LinkedList[V]) Add(v *Element[V]) {
	head := ll.head
	if head == nil {
		ll.head = v
		return
	}

	for head.next != nil {
		head = head.next
	}
	head.next = v
}

func (ll *LinkedList[V]) Insert(v *Element[V], pos int) error {
	pos--
	head := ll.head
	if pos == 0 {
		v.next = head
		ll.head = v
		return nil
	}
	prev := ll.head
	head = head.next
	for i := 1; head != nil; i++ {
		if i == pos {
			prev.next = v
			v.next = head
			return nil
		}
		head = head.next
	}
	return fmt.Errorf("position outsid the range of the linked list")
}

func (ll *LinkedList[V]) Delete(v *Element[V]) error {
	head := ll.head
	if ll.IsEmpty() {
		return fmt.Errorf("linked list is empty")
	}

	if head.value == v.value {
		ll.head = head.next
		return nil
	}
	prev := ll.head
	head = head.next
	for head != nil {
		if head.value == v.value {
			prev.next = head.next
			return nil
		}
		prev = head
		head = head.next
	}
	return fmt.Errorf("element not found")
}

func (ll *LinkedList[V]) Find(v V) (*Element[V], error) {
	head := ll.head
	if ll.IsEmpty() {
		return nil, fmt.Errorf("linked list is empty")
	}

	for head != nil {
		if head.value == v {
			return head, nil
		}
		head = head.next
	}

	return nil, fmt.Errorf("element not found")
}

func (ll *LinkedList[V]) List() []V {
	size := ll.Size()
	slice := make([]V, size)
	head := ll.head
	for i := 0; i < size; i++ {
		slice[i] = head.value
		head = head.next
	}
	return slice
}

func (ll *LinkedList[V]) Size() int {
	size := 0
	head := ll.head
	for head != nil {
		size++
		head = head.next
	}
	return size
}

func (ll *LinkedList[V]) IsEmpty() bool {
	return ll.head == nil
}
