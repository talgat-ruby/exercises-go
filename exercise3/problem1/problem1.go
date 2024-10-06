package problem1

import "slices"

type Queue struct {
	vals []any
}

func (q *Queue) Enqueue(val any) {
	q.vals = append(q.vals, val)
}

func (q *Queue) Dequeue() ([]any, error) {
	valLen := len(q.vals)
	q.vals = slices.Delete(q.vals, valLen-1, valLen)
	return q.vals, nil
}

func (q *Queue) Peek() any {
	valLen := len(q.vals) - 1
	return q.vals[valLen]
}

func (q *Queue) Size() int {
	return len(q.vals)
}

func (q *Queue) IsEmpty() bool {
	if len(q.vals) > 0 {
		return false
	}
	return true
}
