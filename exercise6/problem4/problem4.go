package problem4

import (
	"sync"
	"time"
)

func worker(id int, _ *[]string, ch chan<- int, once *sync.Once) {
	// TODO wait for shopping list to be completed
	once.Do(func() {
		ch <- id
	})

}

func updateShopList(shoppingList *[]string) {
	time.Sleep(10 * time.Millisecond)

	*shoppingList = append(*shoppingList, "apples")
	*shoppingList = append(*shoppingList, "milk")
	*shoppingList = append(*shoppingList, "bake soda")
}

func notifyOnShopListUpdate(shoppingList *[]string, numWorkers int) <-chan int {
	notifier := make(chan int)
	var once sync.Once

	for i := range numWorkers {
		go worker(i+1, shoppingList, notifier, &once)
		time.Sleep(time.Millisecond) // order matters
	}

	go updateShopList(shoppingList)

	return notifier
}
