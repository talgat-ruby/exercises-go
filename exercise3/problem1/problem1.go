package problem1

import "errors"

type Queue struct {
	slc []any
}

func (q *Queue) Enqueue(x any) {
	q.slc = append(q.slc, x)
}

func (q *Queue) Dequeue() (any, error) {
	if len(q.slc) == 0 {
		return Queue{}, errors.New("queue is empty")
	}
	val := q.slc[0]
	q.slc = q.slc[1:]
	return val, nil
}

func (q *Queue) Peek() (any, error) {
	if len(q.slc) == 0 {
		return Queue{}, errors.New("queue is empty")
	}
	return q.slc[0], nil
}

func (q *Queue) Size() int {
	return len(q.slc)
}

func (q *Queue) IsEmpty() bool {
	return 0 == q.Size()
}
