package problem1

import "errors"

type Queue struct {
	slc []interface{}
}

func (q *Queue) Enqueue(value interface{}) {
	q.slc = append(q.slc, value)
}

func (q *Queue) Dequeue() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("Пумпурум, пусто")
	}
	firstValue := q.slc[0]
	q.slc = q.slc[1:]
	return firstValue, nil
}

func (q *Queue) Peek() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("Пумпурум, пусто")
	}
	return q.slc[0], nil
}

func (q *Queue) Size() int {
	return len(q.slc)
}

func (q *Queue) IsEmpty() bool {
	return len(q.slc) == 0
}
