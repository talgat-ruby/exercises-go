package problem1

import "errors"

type Queue struct {
	items []any
}

func (q *Queue) Enqueue(item any) {
	q.items = append(q.items, item)
}
func (q *Queue) Dequeue() (any, error) {
	if len(q.items) == 0 {
		return nil, errors.New("Empty queue")
	}
	first := q.items[0]
	q.items = q.items[1:]
	return first, nil
}
func (q *Queue) Peek() (any, error) {
	if len(q.items) == 0 {
		return nil, errors.New("Empty queue")
	}
	return q.items[0], nil
}
func (q *Queue) Size() int {
	return len(q.items)
}
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}
