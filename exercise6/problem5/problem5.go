package problem5

import (
	"sync"
	"time"
)

var cond = sync.NewCond(&sync.Mutex{})
var isFilled bool

func worker(id int, _ *[]string, ch chan<- int) {
	cond.L.Lock()
	for !isFilled {
		cond.Wait()
	}
	ch <- id
	cond.L.Unlock()
}

func updateShopList(shoppingList *[]string) {
	time.Sleep(10 * time.Millisecond)
	cond.L.Lock()
	defer cond.L.Unlock()

	*shoppingList = append(*shoppingList, "apples")
	*shoppingList = append(*shoppingList, "milk")
	*shoppingList = append(*shoppingList, "bake soda")
	isFilled = true
	cond.Broadcast()
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
