package problem1

import "errors"

type Queue struct {
	queue []any
}

func (q *Queue) Enqueue(val any) {
	q.queue = append(q.queue, val)
}
func (q *Queue) Dequeue() (any, error) {
	if len(q.queue) == 0 {
		return nil, errors.New("queue is empty")
	}

	a := q.queue[0]
	q.queue = q.queue[1:]
	return a, nil
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
	return len(q.queue) == 0
}
