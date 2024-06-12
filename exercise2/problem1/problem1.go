package problem1

import "errors"

type Queue struct {
	values []any
}

func (q *Queue) Enqueue(e any) {
	q.values = append(q.values, e)
}

func (q *Queue) Dequeue() (any, error) {
	if len(q.values) == 0 {
		return nil, errors.New("queue is empty")
	}
	var first any = q.values[0]
	q.values = q.values[1:]
	return first, nil
}

func (q *Queue) Peek() (any, error) {
	if len(q.values) == 0 {
		return nil, errors.New("queue is empty")
	}
	var first = q.values[0]
	return first, nil
}

func (q *Queue) Size() int {
	return len(q.values)
}

func (q *Queue) IsEmpty() bool {
	return len(q.values) == 0
}
