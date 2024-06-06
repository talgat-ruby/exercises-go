package problem1

import "errors"

type Queue struct {
	Arr []any
}

func (q *Queue) Enqueue(elem any) {
	q.Arr = append(q.Arr, elem)
}

func (q *Queue) Dequeue() (any, error) {
	if len(q.Arr) == 0 {
		return nil, errors.New("queue is empty")
	}
	elem := q.Arr[0]
	q.Arr = q.Arr[1:]
	return elem, nil
}
func (q *Queue) Peek() (any, error) {
	if len(q.Arr) == 0 {
		return nil, errors.New("queue is empty")
	}
	return q.Arr[0], nil
}
func (q *Queue) Size() int {
	return len(q.Arr)
}
func (q *Queue) IsEmpty() bool {
	if len(q.Arr) == 0 {
		return true
	}
	return false
}
