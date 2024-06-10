package problem1

import "errors"

type Queue struct {
	vals []any
}

func (q *Queue) Enqueue(val any) {
	q.vals = append(q.vals, val)
}

func (q *Queue) Size() int {
	return len(q.vals)
}

func (q *Queue) IsEmpty() bool {
	return len(q.vals) == 0
}

func (q *Queue) Dequeue() (any, error) {
	if len(q.vals) == 0 {
		return nil, errors.New("queue is empty")
	}
	val := q.vals[0]
	q.vals = q.vals[1:]
	return val, nil
}

func (q *Queue) Peek() (any, error) {
	if len(q.vals) == 0 {
		return nil, errors.New("queue is empty")
	}
	return q.vals[0], nil
}
