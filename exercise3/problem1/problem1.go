package problem1

import (
	"errors"
)

type Queue struct {
	elment []any
}

func (q *Queue) Enqueue(elment any) {
	q.elment = append(q.elment, elment)
}
func (q *Queue) Dequeue() (any, error) {
	if q.elment == nil {
		return nil, errors.New("err")
	}
	first := q.elment[0]
	q.elment = q.elment[1:]
	return first, nil

}
func (q *Queue) Peek() (any, error) {
	if q.elment == nil {
		return nil, errors.New("err")
	}
	return q.elment[0], nil
}
func (q *Queue) Size() int {
	return len(q.elment)
}
func (q *Queue) IsEmpty() bool {
	return len(q.elment) == 0
}
