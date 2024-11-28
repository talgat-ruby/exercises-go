package problem1

import "errors"

type Queue struct {
	items []any
}

func (q *Queue) Enqueue(item any) {
	q.items = append(q.items, item)
}
func (q *Queue) Dequeue() (any, error) {
	if q.IsEmpty() {
		return nil, errors.New("Queue is empty")
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, nil
}
func (q *Queue) Peek() (any, error) {
	if q.IsEmpty() {
		return nil, errors.New("Queue is empty")
	}
	return q.items[0], nil
}
func (q *Queue) Size() int {
	return len(q.items)
}
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}
