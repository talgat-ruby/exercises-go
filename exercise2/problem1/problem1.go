package problem1

import "fmt"

type Queue struct {
	values []any
}

func (q *Queue) Enqueue(any2 any) {
	q.values = append(q.values, any2)
}

func (q *Queue) Dequeue() (any, error) {
	if q.IsEmpty() {
		return nil, fmt.Errorf("empty")
	}
	v := q.values[0]
	q.values = q.values[1:]
	return v, nil
}

func (q *Queue) Peek() (any, error) {
	if q.IsEmpty() {
		return nil, fmt.Errorf("empty")
	}
	return q.values[0], nil
}

func (q *Queue) Size() int {
	return len(q.values)
}

func (q *Queue) IsEmpty() bool {
	return q.Size() == 0
}
