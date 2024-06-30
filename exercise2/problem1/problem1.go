package problem1

import "errors"

type Node struct {
	value interface{}
	next  *Node
}

type Queue struct {
	head *Node
	tail *Node
	size int
}

func (q *Queue) Enqueue(value interface{}) {
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

func (q *Queue) Dequeue() (interface{}, error) {
	if q.head == nil {
		return nil, errors.New("queue is empty")
	}
	value := q.head.value
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
	q.size--
	return value, nil
}

func (q *Queue) Peek() (interface{}, error) {
	if q.head == nil {
		return nil, errors.New("queue is empty")
	}
	return q.head.value, nil
}

func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}
