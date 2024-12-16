package problem5

import (
	"sync"
	"time"
)

func worker(id int, shoppingList *[]string, ch chan<- int, done *sync.Cond) {
	// TODO wait for shopping list to be completed

	done.L.Lock()
	done.Wait()
	done.L.Unlock()
	ch <- id
}

func updateShopList(shoppingList *[]string, done *sync.Cond) {
	time.Sleep(10 * time.Millisecond)

	*shoppingList = append(*shoppingList, "apples")
	*shoppingList = append(*shoppingList, "milk")
	*shoppingList = append(*shoppingList, "bake soda")

	done.Broadcast()
}

func notifyOnShopListUpdate(shoppingList *[]string, numWorkers int) <-chan int {
	notifier := make(chan int)

	var mu sync.Mutex
	done := sync.NewCond(&mu)

	for i := range numWorkers {
		go worker(i+1, shoppingList, notifier, done)
		time.Sleep(time.Millisecond) // order matters
	}

	go updateShopList(shoppingList, done)

	return notifier
}
