package problem4

import (
	"sync"
	"time"
)

func worker(id int, shoppingList *[]string, ch chan<- int, cond *sync.Cond) {
	defer cond.L.Unlock()
	cond.L.Lock()

	if len(*shoppingList) == 0 {
		cond.Wait()
	}

	ch <- id
}

func updateShopList(shoppingList *[]string, cond *sync.Cond) {
	time.Sleep(10 * time.Millisecond)

	defer cond.L.Unlock()
	cond.L.Lock()

	*shoppingList = append(*shoppingList, "apples")
	*shoppingList = append(*shoppingList, "milk")
	*shoppingList = append(*shoppingList, "bake soda")

	cond.Signal()
}

func notifyOnShopListUpdate(shoppingList *[]string, numWorkers int) <-chan int {
	notifier := make(chan int)
	cond := &sync.Cond{L: &sync.Mutex{}}

	for i := range numWorkers {
		go worker(i+1, shoppingList, notifier, cond)
		time.Sleep(time.Millisecond)
	}

	go updateShopList(shoppingList, cond)

	return notifier
}
