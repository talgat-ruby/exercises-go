package problem4

import (
	"errors"
	"fmt"
)

type Element[T comparable] struct {
	value T
	next  *Element[T]
}

type LinkedList[T comparable] struct {
	head *Element[T]
}

func (list *LinkedList[T]) Add(element *Element[T]) {
	newElement := element

	if list.head == nil {
		list.head = newElement
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}
		current.next = newElement
	}
}

func (list *LinkedList[T]) Insert(element *Element[T], index int) error {
	if index < 0 || index > list.Size() {
		return errors.New("index out of range")
	}

	i := 1
	if index == i {
		temp := list.head
		list.head = element
		list.head.next = temp
		return nil
	}

	current := list.head
	for current.next != nil {
		i++
		fmt.Println(current.value)
		if i == index {
			temp := current.next
			current.next = element
			current.next.next = temp
			return nil
		}
		current = current.next
	}
	return nil
}

func (list *LinkedList[T]) Delete(element *Element[T]) error {
	if list.head.value == element.value {
		list.head = list.head.next
	}
	current := list.head
	if current == nil {
		return errors.New("list is empty")
	}
	for current.next != nil {
		if current.next.value == element.value {
			current.next = current.next.next
		}
		current = current.next
	}
	return errors.New("element to delete not found")
}

func (list *LinkedList[T]) Find(value any) (Element[T], error) {
	if list.head.value == value {
		return *list.head, nil
	}
	current := list.head
	for current.next != nil {
		if current.next.value == value {
			return *current.next, nil
		}
		current = current.next
	}
	return Element[T]{}, errors.New("element not found")
}

func (list *LinkedList[T]) List() []T {
	var result []T
	if list.head == nil {
		return result
	}
	current := list.head
	result = append(result, current.value)
	for current.next != nil {
		current = current.next
		result = append(result, current.value)
	}
	return result
}

func (list *LinkedList[T]) Size() int {
	return len(list.List())
}

func (list *LinkedList[T]) IsEmpty() bool {
	return list.Size() == 0
}
