package problem1

import "errors"

type Queue struct {
	array []any
	size  int
}

func (q *Queue) Enqueue(val any) {
	q.array = append(q.array, val)
	q.size++
}
func (q *Queue) Dequeue() (any, error) {
	if q.size <= 0 {
		return nil, errors.New("queue is empty")
	}
	val := q.array[0]
	q.array = q.array[1:]
	q.size--
	return val, nil
}

func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) Peek() (any, error) {
	if q.size == 0 {
		return nil, errors.New("queue is empty")
	}
	return q.array[0], nil
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}
