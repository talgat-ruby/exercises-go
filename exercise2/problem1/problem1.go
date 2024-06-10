package problem1

import "errors"

type Queue struct {
	length int
	array  []any
}

func (q *Queue) Enqueue(val any) {
	q.array = append(q.array[:], val)
	q.length++
}

func (q *Queue) Dequeue() (any, error) {
	if q.length > 0 {
		q.length--
		val := q.array[0]
		q.array = q.array[1 : q.length+1]
		return val, nil
	}
	return nil, errors.New("queue is empty")
}

func (q *Queue) Peek() (any, error) {
	if q.length > 0 {
		return q.array[0], nil
	}
	return nil, errors.New("queue is empty")
}

func (q *Queue) Size() int {
	return q.length
}

func (q *Queue) IsEmpty() bool {
	return q.length == 0
}
