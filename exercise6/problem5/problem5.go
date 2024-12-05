package problem5

import (
	"sync"
	"time"
)

func worker(id int, shoppingList *[]string, ch chan<- int, c *sync.Cond) {
	defer c.L.Unlock()
	c.L.Lock()
	for len(*shoppingList) == 0 {
		c.Wait()
	}
	ch <- id
}

func updateShopList(shoppingList *[]string, c *sync.Cond) {
	defer c.L.Unlock()
	c.L.Lock()
	time.Sleep(10 * time.Millisecond)

	*shoppingList = append(*shoppingList, "apples")
	*shoppingList = append(*shoppingList, "milk")
	*shoppingList = append(*shoppingList, "bake soda")

	c.Broadcast()
}

func notifyOnShopListUpdate(shoppingList *[]string, numWorkers int) <-chan int {
	notifier := make(chan int)
	c := sync.NewCond(&sync.Mutex{})

	for i := range numWorkers {
		go worker(i+1, shoppingList, notifier, c)
		time.Sleep(time.Millisecond) // order matters
	}

	go updateShopList(shoppingList, c)

	return notifier
}
