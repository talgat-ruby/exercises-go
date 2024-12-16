package problem1

import "errors"

type Queue struct {
	elements []any
	size     any
}

func (q *Queue) Enqueue(elem any) {
	if q.Size() == q.size {
		return
	}
	q.elements = append(q.elements, elem)
}

func (q *Queue) Dequeue() (any, error) {
	if q.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	element := q.elements[0]
	if q.Size() == 1 {
		q.elements = nil
		return element, nil
	}
	q.elements = q.elements[1:]
	return element, nil
}

func (q *Queue) Size() int {
	return len(q.elements)
}

func (q *Queue) IsEmpty() bool {
	return len(q.elements) == 0
}

func (q *Queue) Peek() (any, error) {
	if q.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	return q.elements[0], nil
}
