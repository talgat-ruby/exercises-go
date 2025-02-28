package problem1

import "errors"

type Queue struct {
	values []any
}

func (q *Queue) Enqueue(value any) {
	q.values = append(q.values, value)
}

func (q *Queue) Dequeue() (any, error) {
	if q.IsEmpty() {
		return nil, errors.New("Queue is empty")
	}
	valueToRemove := q.values[0]
	q.values = q.values[1:]
	return valueToRemove, nil
}

func (q *Queue) Peek() (any, error) {
	if q.IsEmpty() {
		return nil, errors.New("Queue is empty")
	}
	return q.values[0], nil
}

func (q *Queue) Size() int {
	return len(q.values)
}

func (q *Queue) IsEmpty() bool {
	return q.Size() == 0
}
