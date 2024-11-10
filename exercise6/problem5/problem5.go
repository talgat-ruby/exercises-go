package problem5

import (
	"sync"
	"time"
)

func worker(id int, _ *[]string, ch chan<- int, wg *sync.WaitGroup) {
	// TODO wait for shopping list to be completed
	wg.Wait()
	ch <- id
}

func updateShopList(shoppingList *[]string, wg *sync.WaitGroup) {
	time.Sleep(10 * time.Millisecond)

	*shoppingList = append(*shoppingList, "apples")
	*shoppingList = append(*shoppingList, "milk")
	*shoppingList = append(*shoppingList, "bake soda")
	defer wg.Done()
}

func notifyOnShopListUpdate(shoppingList *[]string, numWorkers int) <-chan int {
	notifier := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)

	for i := range numWorkers {
		go worker(i+1, shoppingList, notifier, &wg)
		time.Sleep(time.Millisecond) // order matters
	}

	go updateShopList(shoppingList, &wg)

	return notifier
}
