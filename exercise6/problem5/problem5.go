// problem5.go
package problem5

import (
	"sync"
	"time"
)

func worker(id int, shoppingList *[]string, ch chan<- int, wg *sync.WaitGroup, mu *sync.Mutex, isReady *bool, cond *sync.Cond) {
	defer wg.Done()

	mu.Lock()
	for !*isReady {
		cond.Wait()
	}
	mu.Unlock()

	// Send the worker ID after releasing the lock
	ch <- id
}

func updateShopList(shoppingList *[]string, mu *sync.Mutex, isReady *bool, cond *sync.Cond) {
	time.Sleep(10 * time.Millisecond)

	mu.Lock()
	*shoppingList = append(*shoppingList, "apples", "milk", "bake soda")
	*isReady = true
	cond.Broadcast()
	mu.Unlock()
}

func notifyOnShopListUpdate(shoppingList *[]string, numWorkers int) <-chan int {
	notifier := make(chan int, numWorkers)
	var wg sync.WaitGroup
	var mu sync.Mutex
	isReady := false
	cond := sync.NewCond(&mu)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i+1, shoppingList, notifier, &wg, &mu, &isReady, cond)
		time.Sleep(time.Millisecond) // Order matters
	}

	go updateShopList(shoppingList, &mu, &isReady, cond)

	go func() {
		wg.Wait()
		close(notifier)
	}()

	return notifier
}
