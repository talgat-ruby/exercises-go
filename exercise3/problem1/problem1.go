package problem1

import (
	"errors"
)

type Queue struct {
	elements []any
}

func (q *Queue) Enqueue(element any) {
	q.elements = append(q.elements, element)
}

func (q *Queue) Dequeue() (any, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	element := q.elements[0]
	q.elements = q.elements[1:]
	return element, nil
}

func (q *Queue) Peek() (any, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	return q.elements[0], nil
}

func (q *Queue) Size() int {
	return len(q.elements)
}

func (q *Queue) IsEmpty() bool {
	return len(q.elements) == 0
}
