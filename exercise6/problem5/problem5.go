package problem5

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, shoppingList *[]string, ch chan<- int, mu *sync.Mutex, wg *sync.WaitGroup, ready *bool) {
	defer wg.Done()
	for {
		mu.Lock()
		if *ready {
			fmt.Println(321)
			ch <- id

			mu.Unlock()
			return
		}
		fmt.Println(654)
		mu.Unlock()

		time.Sleep(10 * time.Millisecond)
	}

	// TODO wait for shopping list to be completed
}

func updateShopList(shoppingList *[]string, mu *sync.Mutex, ready *bool) {
	time.Sleep(10 * time.Millisecond)
	mu.Lock()
	fmt.Println(987)
	*shoppingList = append(*shoppingList, "apples")
	*shoppingList = append(*shoppingList, "milk")
	*shoppingList = append(*shoppingList, "bake soda")
	*ready = true
	mu.Unlock()
}

func notifyOnShopListUpdate(shoppingList *[]string, numWorkers int) <-chan int {
	mu := &sync.Mutex{}
	var wg sync.WaitGroup
	notifier := make(chan int)
	ready := false
	for i := range numWorkers {
		wg.Add(1)
		go worker(i, shoppingList, notifier, mu, &wg, &ready)
		time.Sleep(time.Millisecond) // order matters
	}

	go func() {
		updateShopList(shoppingList, mu, &ready)
		wg.Wait()
		close(notifier)
	}()
	defer close(notifier)
	return notifier
}
