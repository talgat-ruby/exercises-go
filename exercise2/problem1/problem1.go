package problem1

import "errors"

type Queue struct {
	data []any
}

func (q *Queue) Enqueue(val any) {
	q.data = append(q.data, val)
}

func (q *Queue) Dequeue() (any, error) {
	if len(q.data) == 0 {
		return nil, errors.New("Queue is empty")
	}
	val := q.data[0]
	q.data = q.data[1:]
	return val, nil
}

func (q *Queue) Size() int {
	return len(q.data)
}

func (q *Queue) Peek() (any, error) {
	if len(q.data) == 0 {
		return nil, errors.New("Queue is empty")
	}
	return q.data[0], nil
}

func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}
