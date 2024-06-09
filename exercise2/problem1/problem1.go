package problem1

import (
	"errors"
)

type Queue struct {
	items []any
}

func (q *Queue) Enqueue(item any) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() (any, error) {
	if len(q.items) == 0 {
		err := errors.New("queue is empty. Nothing to dequeue")

		return nil, err
	}

	cur := q.items[0]
	q.items = q.items[1:]

	return cur, nil
}

func (q *Queue) Size() int {
	return len(q.items)
}

func (q *Queue) Peek() (any, error) {
	if len(q.items) == 0 {
		err := errors.New("queue is empty. Nothing to peek")

		return nil, err
	}

	return q.items[0], nil
}

func (q *Queue) IsEmpty() bool {
	return q.Size() == 0
}
