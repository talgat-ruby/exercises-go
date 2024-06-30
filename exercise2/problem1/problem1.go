package problem1

import "errors"

type Queue struct {
	items []any
	size  int
}

// push into the queue
func (q *Queue) Enqueue(item any) {
	q.items = append(q.items, item)
	q.size++
}

// pop from the queue
func (q *Queue) Dequeue() (any, error) {
	if q.size == 0 {
		return nil, errors.New("queue is empty")
	}
	poppedItem := q.items[0]
	q.items = q.items[1:]
	q.size--
	return poppedItem, nil
}

// peek into the queue
func (q *Queue) Peek() (any, error) {
	if q.size == 0 {
		return nil, errors.New("queue is empty")
	}
	return q.items[0], nil
}

// size of the queue
func (q *Queue) Size() int {
	return q.size
}

// is empty?
func (q *Queue) IsEmpty() bool {
	return q.size == 0
}
