package problem1

import "errors"

type Queue struct {
	Elements []any
}

func (q *Queue) Enqueue(element any) {
	q.Elements = append(q.Elements, element)
}

func (q *Queue) Dequeue() (any, error) {
	if len(q.Elements) == 0 {
		return 0, errors.New("Error")
	}

	element := q.Elements[0]
	q.Elements = q.Elements[1:]
	return element, nil
}

func (q *Queue) Peek() (any, error) {
	if len(q.Elements) == 0 {
		return 0, errors.New("Error")
	}

	element := q.Elements[0]
	return element, nil
}

func (q *Queue) Size() any {
	return len(q.Elements)
}

func (q *Queue) IsEmpty() bool {
	return len(q.Elements) == 0
}
