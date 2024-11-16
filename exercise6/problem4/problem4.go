package problem4

import (
	"sync"
	"time"
)

// Worker function
func worker(id int, shoppingList *[]string, mutex *sync.Mutex, notifier *int, isListFilled *bool, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		mutex.Lock()
		// Wait until the shopping list is filled
		if *isListFilled {
			// Notify that this worker is first to be notified
			if *notifier == 0 {
				*notifier = id
			}
			mutex.Unlock()
			return
		}
		mutex.Unlock()
		time.Sleep(1 * time.Millisecond) // Simulate waiting
	}
}

// Updates the shopping list
func updateShopList(shoppingList *[]string, mutex *sync.Mutex, isListFilled *bool) {
	time.Sleep(10 * time.Millisecond)

	mutex.Lock()
	*shoppingList = append(*shoppingList, "apples", "milk", "bake soda")
	*isListFilled = true
	mutex.Unlock()
}

// Notify workers when the shopping list is filled
func notifyOnShopListUpdate(shoppingList *[]string, numWorkers int) <-chan int {
	notifier := make(chan int)
	var wg sync.WaitGroup
	var mutex sync.Mutex
	var isListFilled bool
	var firstNotifier int

	// Launch all workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, shoppingList, &mutex, &firstNotifier, &isListFilled, &wg)
		time.Sleep(time.Millisecond) // Maintain worker order
	}

	// Fill the shopping list and mark it as ready
	go func() {
		updateShopList(shoppingList, &mutex, &isListFilled)
		wg.Wait()
		notifier <- firstNotifier
		close(notifier)
	}()

	return notifier
}
