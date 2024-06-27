package problem1

import "errors"

type Queue struct{
	linkedList []any
}

func (q *Queue) Enqueue(value any) {
	q.linkedList = append(q.linkedList, value)
}

func(q *Queue) Dequeue() (any, error) {
	if q.IsEmpty() {
		return nil, errors.New("linked list is empty")
	}

	del, newLL := q.linkedList[0], q.linkedList[1:]
	q.linkedList = newLL

	return del, nil

}

func(q *Queue) Size() int {
	size := len(q.linkedList)

	return size
}

func(q *Queue) Peek() (any, error) {
	if q.IsEmpty() {
		return nil, errors.New("linked list is empty")
	}

	del := q.linkedList[0]

	return del, nil
}

func(q *Queue) IsEmpty() bool{
	return q.Size() == 0
}