package problem1

import "errors"

type Node struct {
	value any
	next  *Node
}

type Queue struct {
	head *Node
	tail *Node
	size int
}

func (q *Queue) Enqueue(value any) {
	newNode := &Node{value: value}
	if q.tail != nil {
		q.tail.next = newNode
	}
	q.tail = newNode
	if q.head == nil {
		q.head = newNode
	}
	q.size++
}
func (q *Queue) Dequeue() (any, error) {
	if q.head == nil {
		return nil, errors.New("dequeue from empty queue")
	}
	value := q.head.value
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
	q.size--
	return value, nil
}
func (q *Queue) Peek() (any, error) {
	if q.head == nil {
		return nil, errors.New("peek from empty queue")
	}
	return q.head.value, nil
}

func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}
