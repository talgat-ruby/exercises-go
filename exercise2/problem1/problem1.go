package problem1

import "errors"

type Queue struct {
	Elements []any
}

func (q *Queue) Enqueue(val any) {
	q.Elements = append(q.Elements, val)
}

func (q *Queue) Dequeue() (any, error) {
	if q.IsEmpty() {
		return nil, errors.New("Queue is empty")
	}
	dequeued := q.Elements[0]
	q.Elements = q.Elements[1:]
	return dequeued, nil
}

func (q *Queue) Peek() (any, error) {
	if q.IsEmpty() {
		return nil, errors.New("Queue is empty")
	}
	return q.Elements[0], nil
}

func (q *Queue) IsEmpty() bool {
	return len(q.Elements) == 0
}

func (q *Queue) Size() int {
	return len(q.Elements)
}
