package problem5

import (
	"time"
)

func worker(id int, shoppingList *[]string, ch chan<- int) {
	// Wait for the shopping list to be completed
	ch <- id
}

func updateShopList(shoppingList *[]string) {
	time.Sleep(10 * time.Millisecond)

	*shoppingList = append(*shoppingList, "apples")
	*shoppingList = append(*shoppingList, "milk")
	*shoppingList = append(*shoppingList, "bake soda")
}

func notifyOnShopListUpdate(shoppingList *[]string, numWorkers int) <-chan int {
	notifier := make(chan int)

	// Iterate using numWorkers as the loop limit
	for i := 0; i < numWorkers; i++ {
		go worker(i+1, shoppingList, notifier)
		time.Sleep(time.Millisecond) // order matters
	}

	go updateShopList(shoppingList)

	return notifier
}
