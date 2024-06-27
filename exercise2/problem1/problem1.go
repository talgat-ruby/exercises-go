package problem1

import "errors"

type Queue struct {
	elements []any
}

func (q *Queue) Enqueue(val any) {
	q.elements = append(q.elements, val)
}

func (q *Queue) Size() int {
	return len(q.elements)
}

func (q *Queue) Dequeue() (any, error) {
	if len(q.elements) == 0 {
		return nil, errors.New("queue is empty")
	}

	val := q.elements[0]
	q.elements = q.elements[1:]

	return val, nil
}

func (q *Queue) IsEmpty() bool {
	return len(q.elements) == 0
}

func (q *Queue) Peek() (any, error) {
	if len(q.elements) == 0 {
		return nil, errors.New("queue is empty")
	}

	return q.elements[0], nil
}
