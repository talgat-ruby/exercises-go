package main

import "fmt"

// Queue represents a simple queue structure with basic operations
type Queue[T any] struct {
	items []T
}

// Enqueue adds an element to the back of the Queue
func (q *Queue[T]) Enqueue(value T) {
	q.items = append(q.items, value)
}

// Dequeue removes and returns the element at the front of the Queue
// If the Queue is empty, it returns the zero value of the element type and false
func (q *Queue[T]) Dequeue() (T, bool) {
	if q.IsEmpty() {
		var zero T
		return zero, false
	}
	value := q.items[0]
	q.items = q.items[1:]
	return value, true
}

// Peek returns the element at the front of the Queue without removing it
// If the Queue is empty, it returns the zero value of the element type and false
func (q *Queue[T]) Peek() (T, bool) {
	if q.IsEmpty() {
		var zero T
		return zero, false
	}
	return q.items[0], true
}

// Size returns the number of elements in the Queue
func (q *Queue[T]) Size() int {
	return len(q.items)
}

// IsEmpty checks if the Queue is empty
func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

// Test the Queue data structure
func main() {
	queue := &Queue[int]{}

	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)

	fmt.Println("Size:", queue.Size())       // Output: Size: 3
	fmt.Println("Peek:", queue.Peek())       // Output: Peek: 10
	fmt.Println("Dequeue:", queue.Dequeue()) // Output: Dequeue: 10
	fmt.Println("Size after dequeue:", queue.Size()) // Output: Size after dequeue: 2
	fmt.Println("IsEmpty:", queue.IsEmpty()) // Output: IsEmpty: false
	queue.Dequeue()
	queue.Dequeue()
	fmt.Println("IsEmpty after all dequeues:", queue.IsEmpty()) // Output: IsEmpty after all dequeues: true
}
