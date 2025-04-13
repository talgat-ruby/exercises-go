package problem1

import "errors"

type Queue struct {
	arr []any
}

func (q *Queue) Enqueue(element any) {
	q.arr = append(q.arr, element)
}

func (q *Queue) Dequeue() (any, error) {
	if q.IsEmpty() {
		var nullValue any
		return nullValue, errors.New("queue is empty")
	}
	val := q.arr[0]
	q.arr = q.arr[1:]
	return val, nil
}

func (q *Queue) Peek() (any, error) {
	if q.IsEmpty() {
		var nullValue any
		return nullValue, errors.New("queue is empty")
	}
	return q.arr[0], nil
}

func (q *Queue) Size() int {
	return len(q.arr)
}

func (q *Queue) IsEmpty() bool {
	return q.Size() == 0
}
