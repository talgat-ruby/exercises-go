package problem4

import "errors"

type Element[T comparable] struct {
	value T
	next  *Element[T]
}

type LinkedList[T comparable] struct {
	element *Element[T]
}

func (l *LinkedList[T]) Add(newElement *Element[T]) {
	if l.element == nil {
		l.element = newElement
	} else {
		currentElement := l.element
		for currentElement.next != nil {
			currentElement = currentElement.next
		}
		currentElement.next = newElement
	}
}

func (l *LinkedList[T]) Insert(newElement *Element[T], pos int) error {
	if l.element == nil {
		return nil
	} else {
		currentElement := l.element
		for i := 0; i < pos; i++ {
			if currentElement == nil || (currentElement.next == nil && i < pos-1) {
				return errors.New("index out of bounds")
			}
			currentElement = currentElement.next
		}
		newElement.next = currentElement.next
		currentElement.next = newElement
		return nil
	}
}

func (l *LinkedList[T]) Delete(element *Element[T]) error {
	if l.element == nil {
		return errors.New("list is empty")
	}
	if element.value == l.element.value {
		l.element = l.element.next
		return nil
	}
	currentElement := l.element
	for currentElement.next != nil && currentElement.next.value != element.value {
		currentElement = currentElement.next
	}
	if currentElement.next == nil {
		return errors.New("element not found")
	}
	currentElement.next = currentElement.next.next
	return nil
}

func (l *LinkedList[T]) Find(value T) (*Element[T], error) {
	currentElement := l.element
	for currentElement != nil {
		if currentElement.value == value {
			return currentElement, nil
		}
		currentElement = currentElement.next
	}
	return nil, errors.New("element not found")
}

func (l *LinkedList[T]) List() []T {
	currentElement := l.element
	var list []T
	for currentElement != nil {
		list = append(list, currentElement.value)
		currentElement = currentElement.next
	}
	return list
}

func (l *LinkedList[T]) Size() int {
	size := 0
	currentElement := l.element
	for currentElement != nil {
		size++
		currentElement = currentElement.next
	}
	return size
}

func (l *LinkedList[T]) IsEmpty() bool {
	return l.element == nil
}
