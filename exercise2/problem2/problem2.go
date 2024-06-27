package problem2

import "errors"

type Stack struct {
	elements []any
}

func (q *Stack) Push(element any) {
	if q.elements == nil {
		q.elements = make([]any, 0)
	}
	q.elements = append([]any{element}, q.elements...)
}

func (q *Stack) Pop() (any, error) {
	if q.IsEmpty() {
		return nil, errors.New("empty")
	}

	first := q.elements[0]
	q.elements = q.elements[1:]

	return first, nil
}

func (q *Stack) Peek() (any, error) {
	if q.IsEmpty() {
		return nil, errors.New("empty")
	}

	return q.elements[0], nil
}

func (q *Stack) Size() int {
	return len(q.elements)
}

func (q *Stack) IsEmpty() bool {
	return q.elements == nil || len(q.elements) == 0
}
