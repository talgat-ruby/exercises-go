package problem5

import (
	"sync"
	"time"
)

func worker(id int, shoppingList *[]string, notifier *[]int, cond *sync.Cond, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	for len(*shoppingList) == 0 {
		cond.Wait()
	}
	*notifier = append(*notifier, id)
	mu.Unlock()
}

func updateShopList(shoppingList *[]string, cond *sync.Cond, mu *sync.Mutex) {
	time.Sleep(10 * time.Millisecond)
	mu.Lock()
	*shoppingList = append(*shoppingList, "apples", "milk", "bake soda")
	cond.Broadcast()
	mu.Unlock()
}

func notifyOnShopListUpdate(shoppingList *[]string, numWorkers int) []int {
	var notifier []int
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i+1, shoppingList, &notifier, cond, &mu, &wg)
		time.Sleep(time.Millisecond)
	}

	go func() {
		updateShopList(shoppingList, cond, &mu)
		wg.Wait()
	}()

	wg.Wait()
	return notifier
}
