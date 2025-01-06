package problem1

import (
	"errors"
)

type Queue struct {
	items []any
}

func (q *Queue) Enqueue(val any) {
	q.items = append(q.items, val)
}

func (q *Queue) Dequeue() (any, error) {
	if len(q.items) == 0 {
		return nil, errors.New("queue is empty")
	}
	element := q.items[0]
	q.items = q.items[1:]
	return element, nil
}

func (q *Queue) Peek() (any, error) {
	if len(q.items) == 0 {
		return nil, errors.New("queue is empty")
	}
	return q.items[0], nil
}

func (q *Queue) Size() int {
	return len(q.items)
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}
