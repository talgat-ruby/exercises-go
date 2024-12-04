package problem5

import (
	"sync"
	"time"
)

func worker(id int, shoppingList *[]string, cond *sync.Cond, ch chan<- int) {
	cond.L.Lock()
	defer cond.L.Unlock()
	for len(*shoppingList) == 0 {
		cond.Wait()
	}

	ch <- id
}
func updateShopList(shoppingList *[]string, cond *sync.Cond) {
	time.Sleep(10 * time.Millisecond)
	cond.L.Lock()
	defer cond.L.Unlock()
	*shoppingList = append(*shoppingList, "apples")
	*shoppingList = append(*shoppingList, "milk")
	*shoppingList = append(*shoppingList, "bake soda")

	cond.Broadcast()
}

func notifyOnShopListUpdate(shoppingList *[]string, numWorkers int) <-chan int {
	notifier := make(chan int)
	m := sync.Mutex{}
	cond := sync.NewCond(&m)
	for i := range numWorkers {
		go worker(i+1, shoppingList, cond, notifier)
		time.Sleep(time.Millisecond) // order matters
	}

	go updateShopList(shoppingList, cond)

	return notifier
}
