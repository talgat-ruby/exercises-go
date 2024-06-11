package problem1

import "fmt"

type Node struct {
	Data any
	Tail *Node
}

type Queue struct {
	head *Node
}

func (q *Queue) Size() int {
	head := q.head
	size := 0
	for head != nil {
		head = head.Tail
		size++
	}

	return size
}

func (q *Queue) IsEmpty() bool {
	return q.head == nil
}

func (q *Queue) Enqueue(v any) {
	node := q.head
	if node == nil {
		q.head = &Node{Data: v}
		return
	}

	for node.Tail != nil {
		node = node.Tail
	}
	node.Tail = &Node{Data: v}
}

func (q *Queue) Dequeue() (any, error) {
	if q.head == nil {
		return nil, fmt.Errorf("empty queue")
	}

	v := q.head.Data
	q.head = q.head.Tail
	return v, nil
}

func (q *Queue) Peek() (any, error) {
	if q.head == nil {
		return nil, fmt.Errorf("empty queue")
	}

	v := q.head.Data
	return v, nil
}
