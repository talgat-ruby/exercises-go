package main

import (
	"fmt"
	"slices"
)

type Queue struct {
	vals []any
}

func (q *Queue) Enqueue(val any) {
	q.vals = append(q.vals, val)
}

func (q *Queue) Dequeue() (any, error) {
	valLen := len(q.vals)
	queueElement := q.vals[valLen-1]
	q.vals = slices.Delete(q.vals, valLen-1, valLen)
	return queueElement, nil
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

func main() {
	queue := Queue{}
	queue.Enqueue(1)
	queue.Enqueue(5)
	queue.Enqueue(9)
	fmt.Println(queue)
	queue.Dequeue()
	fmt.Println(queue)
	e := queue.Peek()
	fmt.Println(e)
	size := queue.Size()
	fmt.Println(size)
	fmt.Println(queue.IsEmpty())

	table := []struct {
		vals []any
	}{
		{[]any{1, 2, 3}},
		{[]any{"1", "2", "3", "4"}},
		{[]any{true, false}},
	}

	for _, r := range table {
		for _, v := range r.vals {
			queue.Enqueue(v)
		}
		for i := range r.vals {
			d, _ := queue.Dequeue()
			if d != r.vals[i] {
				fmt.Println("Error")
			}
		}
	}
}
