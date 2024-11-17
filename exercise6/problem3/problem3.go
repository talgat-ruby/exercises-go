 Eldar
package main

import (
	"fmt"
	"sync/atomic"
)

// counter represents a thread-safe counter using atomic operations
type counter struct {
	value int64
}

// Increment atomically increases the counter by 1
func (c *counter) Increment() {
	atomic.AddInt64(&c.value, 1)
}

// Decrement atomically decreases the counter by 1
func (c *counter) Decrement() {
	atomic.AddInt64(&c.value, -1)
}

// Value atomically loads and returns the current counter value
func (c *counter) Value() int64 {
	return atomic.LoadInt64(&c.value)
}

func main() {
	c := &counter{}

	// Example usage
	c.Increment()
	fmt.Println("Counter after increment:", c.Value()) // Should print 1

	c.Decrement()
	fmt.Println("Counter after decrement:", c.Value()) // Should print 0

package problem3

type counter struct {
	val int64
}

func newCounter() *counter {
	return &counter{
		val: 0,
	}
 main
}
