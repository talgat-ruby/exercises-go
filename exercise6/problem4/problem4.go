package problem4

import (
	"time"
)

func worker(id int, shoppingList *[]string, ch chan<- int) {
	// Ждем, пока список не будет заполнен
	for len(*shoppingList) == 0 {
		time.Sleep(time.Millisecond)
	}

	select {
	case ch <- id:

	default:

	}
}

func updateShopList(shoppingList *[]string) {
	time.Sleep(10 * time.Millisecond)
	*shoppingList = append(*shoppingList, "apples")
	*shoppingList = append(*shoppingList, "milk")
	*shoppingList = append(*shoppingList, "bake soda")
}

func notifyOnShopListUpdate(shoppingList *[]string, numWorkers int) <-chan int {
	notifier := make(chan int, 1)

	for i := range numWorkers {
		go worker(i+1, shoppingList, notifier)
		time.Sleep(time.Millisecond)
	}

	go updateShopList(shoppingList)

	return notifier
}
