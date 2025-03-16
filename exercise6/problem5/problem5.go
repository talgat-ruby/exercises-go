HEAD

Eldar
6786b69556adaf3ccffc518c7655a26812d55e64
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

// NotifyAllWorkers notifies all workers waiting on the condition
func NotifyAllWorkers(cond *sync.Cond) {
	cond.L.Lock()
	cond.Broadcast() // Notify all waiting workers
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
	fmt.Println("Shopping list is ready. Notifying all workers...")

	// Notify all workers
	NotifyAllWorkers(cond)

	wg.Wait() // Wait for all workers to complete
HEAD


package problem5

import (
	"time"
)

func worker(id int, shoppingList *[]string, ch chan<- int) {
	// TODO wait for shopping list to be completed
	ch <- id
}

func updateShopList(shoppingList *[]string) {
	time.Sleep(10 * time.Millisecond)

	*shoppingList = append(*shoppingList, "apples")
	*shoppingList = append(*shoppingList, "milk")
	*shoppingList = append(*shoppingList, "bake soda")
}

func notifyOnShopListUpdate(shoppingList *[]string, numWorkers int) <-chan int {
	notifier := make(chan int)

	for i := range numWorkers {
		go worker(i+1, shoppingList, notifier)
		time.Sleep(time.Millisecond) // order matters
	}

	go updateShopList(shoppingList)

	return notifier
 main
6786b69556adaf3ccffc518c7655a26812d55e64
}
