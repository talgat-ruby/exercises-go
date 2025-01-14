package problem4

import (
	"sync"
	"time"
)

func worker(id int, shoppingList *[]string, c *sync.Cond, ch chan<- int) {
	c.L.Lock()
	defer c.L.Unlock()
	for len(*shoppingList) == 0 {
		c.Wait()
	}
	ch <- id
}
func updateShopList(shoppingList *[]string, c *sync.Cond) {
	time.Sleep(10 * time.Millisecond)
	c.L.Lock()
	defer c.L.Unlock()
	*shoppingList = append(*shoppingList, "apples")
	*shoppingList = append(*shoppingList, "milk")
	*shoppingList = append(*shoppingList, "bake soda")
	c.Signal()
}

func notifyOnShopListUpdate(shoppingList *[]string, numWorkers int) <-chan int {
	notifier := make(chan int)
	m := sync.Mutex{}
	newCond := sync.NewCond(&m)
	for i := range numWorkers {
		go worker(i+1, shoppingList, newCond, notifier)
		time.Sleep(time.Millisecond) // order matters
	}
	go updateShopList(shoppingList, newCond)
	return notifier
}
