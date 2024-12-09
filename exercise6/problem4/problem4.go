package problem4

import (
	"sync"
	"time"
)

func worker(id int, ready *bool, mu *sync.Mutex, ch chan<- int) {
	mu.Lock()
	defer mu.Unlock()

	for !*ready {
		mu.Unlock()
		time.Sleep(time.Millisecond)
		mu.Lock()
	}

	if id == 1 {
		ch <- id
	}
}

func updateShopList(shoppingList *[]string, ready *bool, mu *sync.Mutex) {
	time.Sleep(10 * time.Millisecond)

	mu.Lock()
	*shoppingList = append(*shoppingList, "apples", "milk", "bake soda")
	*ready = true
	mu.Unlock()
}

func notifyOnShopListUpdate(shoppingList *[]string, numWorkers int) <-chan int {
	notifier := make(chan int)
	ready := false
	var mu sync.Mutex

	for i := 0; i < numWorkers; i++ {
		go worker(i+1, &ready, &mu, notifier)
		time.Sleep(time.Millisecond)
	}

	go updateShopList(shoppingList, &ready, &mu)

	return notifier
}
