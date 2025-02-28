package problem4

import (
	"errors"
	"fmt"
)

type LinkedList[T comparable] struct {
	head *Element[T]
}

type Element[T comparable] struct {
	value T
	next  *Element[T]
}

func (list *LinkedList[T]) Add(element *Element[T]) {
	newElement := element
	if list.head == nil {
		list.head = newElement
	} else {
		currentList := list.head
		for currentList.next != nil {
			currentList = currentList.next
		}
		currentList.next = newElement
	}
}

func (list *LinkedList[T]) Insert(element *Element[T], index int) error {
	if index < 0 || index > list.Size() {
		return errors.New("индекс превышен")
	}

	i := 1
	if index == i {
		temporary := list.head
		list.head = element
		list.head.next = temporary
		return nil
	}

	currentList := list.head
	for currentList.next != nil {
		i++
		fmt.Println(currentList.value)
		if i == index {
			temp := currentList.next
			currentList.next = element
			currentList.next.next = temp
			return nil
		}
		currentList = currentList.next
	}
	return nil
}

func (list *LinkedList[T]) Delete(element *Element[T]) error {
	if list.head.value == element.value {
		list.head = list.head.next
	}
	currentList := list.head
	if currentList == nil {
		return errors.New("пумпумпум пусто")
	}
	for currentList.next != nil {
		if currentList.next.value == element.value {
			currentList.next = currentList.next.next
		}
		currentList = currentList.next
	}
	return errors.New("не найден")
}

func (list *LinkedList[T]) Find(value any) (Element[T], error) {
	if list.head.value == value {
		return *list.head, nil
	}
	currentList := list.head
	for currentList.next != nil {
		if currentList.next.value == value {
			return *currentList.next, nil
		}
		currentList = currentList.next
	}
	return Element[T]{}, errors.New("не найден")
}

func (list *LinkedList[T]) List() []T {
	var result []T
	if list.head == nil {
		return result
	}
	currentList := list.head
	result = append(result, currentList.value)
	for currentList.next != nil {
		currentList = currentList.next
		result = append(result, currentList.value)
	}
	return result
}

func (list *LinkedList[T]) Size() int {
	return len(list.List())
}

func (list *LinkedList[T]) IsEmpty() bool {
	return list.Size() == 0
}
