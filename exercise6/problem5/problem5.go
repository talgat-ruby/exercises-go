package problem5

import (
	"sync"
	"time"
)

var mu sync.Mutex
var cond = sync.NewCond(&mu)

func worker(id int, shoppingList *[]string, ch chan<- int) {
	mu.Lock()
	for len(*shoppingList) == 0 {
		cond.Wait()
	}
	mu.Unlock()

	ch <- id
}

func updateShopList(shoppingList *[]string) {
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

	for i := 0; i < numWorkers; i++ {
		go worker(i+1, shoppingList, notifier)
		time.Sleep(time.Millisecond)
	}

	go updateShopList(shoppingList)

	return notifier
}
