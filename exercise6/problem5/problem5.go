package problem5

import (
	"sync"
	"time"
)

var listCond = sync.NewCond(&sync.Mutex{})
var listFilled bool

func worker(id int, shoppingList *[]string, ch chan<- int) {
	listCond.L.Lock()
	for !listFilled {
		listCond.Wait()
	}
	ch <- id
	listCond.L.Unlock()
}

func updateShopList(shoppingList *[]string) {
	time.Sleep(10 * time.Millisecond)
	*shoppingList = append(*shoppingList, "apples", "milk", "bake soda")
	listCond.L.Lock()
	listFilled = true
	listCond.Broadcast()
	listCond.L.Unlock()
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
