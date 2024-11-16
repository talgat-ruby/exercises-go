package problem5

import (
	"sync"
	"time"
)

// Worker function
func worker(id int, shoppingList *[]string, wg *sync.WaitGroup, mutex *sync.Mutex, ch chan<- int) {
	defer wg.Done()

	for {
		mutex.Lock()
		// Check if the shopping list is filled
		if len(*shoppingList) > 0 {
			ch <- id
			mutex.Unlock()
			return
		}
		mutex.Unlock()
		time.Sleep(time.Millisecond) // Allow other workers to check
	}
}

// Updates the shopping list
func updateShopList(shoppingList *[]string, mutex *sync.Mutex) {
	time.Sleep(10 * time.Millisecond)

	mutex.Lock()
	*shoppingList = append(*shoppingList, "apples", "milk", "bake soda")
	mutex.Unlock()
}

// Notify workers once the shopping list is filled
func notifyOnShopListUpdate(shoppingList *[]string, numWorkers int) <-chan int {
	notifier := make(chan int)
	var wg sync.WaitGroup
	var mutex sync.Mutex

	// Launch all workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i+1, shoppingList, &wg, &mutex, notifier)
	}

	// Fill the shopping list
	go func() {
		updateShopList(shoppingList, &mutex)
		wg.Wait() // Wait for all workers to finish
		close(notifier)
	}()

	return notifier
}
