package problem1

import "errors"

type Queue struct {
	elements []interface{}
}

func (q *Queue) Enqueue(element interface{}) {
	q.elements = append(q.elements, element)
}

func (q Queue) IsEmpty() bool {
	return len(q.elements) == 0
}

func (q *Queue) Dequeue() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("Queue is empty")
	}
	element := q.elements[0]
	q.elements = q.elements[1:]
	return element, nil
}

func (q *Queue) Peek() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("Queue is empty")
	}
	return q.elements[0], nil
}

func (q *Queue) Size() int {
	return len(q.elements)
}
