package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	value int
	mu    sync.Mutex
}

func (c *Counter) Increment() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

// incrementConcurrently increments the counter concurrently across n goroutines
func incrementConcurrently(counter *Counter, n int) {
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			counter.Increment()
		}(
