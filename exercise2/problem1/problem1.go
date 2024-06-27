package problem1

import "errors"

type Queue struct {
	elements []any
}

func (q *Queue) Enqueue(element any) {
	if q.elements == nil {
		q.elements = make([]any, 0)
	}
	q.elements = append(q.elements, element)
}

func (q *Queue) Dequeue() (any, error) {
	if q.IsEmpty() {
		return nil, errors.New("empty")
	}

	first := q.elements[0]
	q.elements = q.elements[1:]

	return first, nil
}

func (q *Queue) Peek() (any, error) {
	if q.IsEmpty() {
		return nil, errors.New("empty")
	}

	return q.elements[0], nil
}

func (q *Queue) Size() int {
	return len(q.elements)
}

func (q *Queue) IsEmpty() bool {
	return q.elements == nil || len(q.elements) == 0
}
