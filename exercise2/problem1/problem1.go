package problem1

import "errors"

type Queue struct {
	queue []any
}

func (q *Queue) Enqueue(value any) {
	q.queue = append(q.queue, value)
}

func (q *Queue) Dequeue() (any, error) {
	if len(q.queue) == 0 {
		return nil, errors.New("queue is empty")
	}

	item := q.queue[0]
	q.queue = q.queue[1:]

	return item, nil
}

func (q *Queue) Peek() (any, error) {
	if len(q.queue) == 0 {
		return nil, errors.New("queue is empty")
	}

	return q.queue[0], nil
}

func (q *Queue) Size() int {
	return len(q.queue)
}

func (q *Queue) IsEmpty() bool {
	if len(q.queue) == 0 {
		return true
	}

	return false
}
