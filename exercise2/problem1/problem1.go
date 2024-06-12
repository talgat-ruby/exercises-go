package problem1

import "errors"

type Queue struct {
	elements []any
}

func (q *Queue) Enqueue(newElement any) {
	q.elements = append(q.elements, newElement)
}

func (q *Queue) Dequeue() (any, error) {
	if len(q.elements) == 0 {
		return nil, errors.New("queue is empty")
	}
	firstElement := q.elements[0]
	q.elements = q.elements[1:]
	return firstElement, nil
}

func (q *Queue) Peek() (any, error) {
	if len(q.elements) == 0 {
		return nil, errors.New("queue is empty")
	}
	return q.elements[0], nil
}

func (q *Queue) Size() any {
	return len(q.elements)
}

func (q *Queue) IsEmpty() any {
	return len(q.elements) == 0
}
