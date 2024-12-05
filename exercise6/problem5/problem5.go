package problem5

import (
	"sync"
	"time"
)

var shopListCond = sync.NewCond(&sync.Mutex{})
var listFull bool

func worker(id int, shoppingList *[]string, ch chan<- int) {
	shopListCond.L.Lock()
	defer shopListCond.L.Unlock()
	for !listFull {
		shopListCond.Wait()
	}
	ch <- id
}

func updateShopList(shoppingList *[]string) {
	time.Sleep(10 * time.Millisecond)

	*shoppingList = append(*shoppingList, "apples")
	*shoppingList = append(*shoppingList, "milk")
	*shoppingList = append(*shoppingList, "bake soda")
	shopListCond.L.Lock()
	listFull = true
	shopListCond.Broadcast()
	shopListCond.L.Unlock()
}

func notifyOnShopListUpdate(shoppingList *[]string, numWorkers int) <-chan int {
	notifier := make(chan int)

	for i := range numWorkers {
		go worker(i+1, shoppingList, notifier)
		time.Sleep(time.Millisecond) // order matters
	}

	go updateShopList(shoppingList)

	return notifier
}
