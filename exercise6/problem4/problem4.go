package problem4

import (
	"sync"
	"time"
)

func worker(id int, shoppingList *[]string, ch chan<- int, cond *sync.Cond, mu *sync.Mutex, once *sync.Once) {
	mu.Lock()
	for len(*shoppingList) == 0 {
		cond.Wait()
	}
	once.Do(func() {
		ch <- id
	})
	mu.Unlock()
}

func updateShopList(shoppingList *[]string, cond *sync.Cond, mu *sync.Mutex) {
	time.Sleep(10 * time.Millisecond)

	mu.Lock()
	*shoppingList = append(*shoppingList, "apples")
	*shoppingList = append(*shoppingList, "milk")
	*shoppingList = append(*shoppingList, "bake soda")
	cond.Broadcast()
	mu.Unlock()
}

func notifyOnShopListUpdate(shoppingList *[]string, numWorkers int) <-chan int {
	notifier := make(chan int)
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	var once sync.Once

	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			worker(workerID, shoppingList, notifier, cond, &mu, &once)
		}(i + 1)
		time.Sleep(time.Millisecond)
	}

	go func() {
		updateShopList(shoppingList, cond, &mu)
		wg.Wait()
		close(notifier)
	}()

	return notifier
}
