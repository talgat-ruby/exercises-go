package problem1

import "errors"

type Queue struct {
	array []any
}

func (queue *Queue) Enqueue(element any) {
	queue.array = append(queue.array, element)
}

func (queue *Queue) Dequeue() (any, error) {
	if queue.IsEmpty() {
		var zeroValue any
		return zeroValue, errors.New("queue is empty")
	}
	val := queue.array[0]
	queue.array = queue.array[1:]
	return val, nil
}

func (queue *Queue) Peek() (any, error) {
	if queue.IsEmpty() {
		var zeroValue any
		return zeroValue, errors.New("queue is empty")
	}
	return queue.array[0], nil
}

func (queue *Queue) Size() int {
	return len(queue.array)
}

func (queue *Queue) IsEmpty() bool {
	return queue.Size() == 0
}
