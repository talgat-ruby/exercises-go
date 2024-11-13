package main

import (
	"fmt"
	"sync"
	"time"
)

// Worker represents a worker with an ID
type Worker struct {
	id int
}

// WaitForShoppingList simulates a worker waiting for the shopping list
func (w *Worker) WaitForShoppingList(cond *sync.Cond, wg *sync.WaitGroup) {
	defer wg.Done()
	cond.L.Lock()
	cond.Wait() // Worker waits for the shopping list to be ready
	fmt.Printf("Worker %d notified and starting shopping\n", w.id)
	cond.L.Unlock()
}

// NotifyFirstWorker notifies the first worker waiting on the condition
func NotifyFirstWorker(cond *sync.Cond) {
	cond.L.Lock()
	cond.Signal() // Notify only the first worker waiting
	cond.L.Unlock()
}

func main() {
	var wg sync.WaitGroup
	mu := &sync.Mutex{}
	cond := sync.NewCond(mu)

	// Create and start several workers waiting for the shopping list
	workers := []Worker{{1}, {2}, {3}, {4}}
	for _, worker := range workers {
		wg.Add(1)
		go worker.WaitForShoppingList(cond, &wg)
	}

	// Simulate filling the shopping list after a delay
	time.Sleep(2 * time.Second)
	fmt.Println("Shopping list is ready. Notifying the first worker...")

	// Notify only the first worker
	NotifyFirstWorker(cond)

	wg.Wait() // Wait for all workers to complete
}
