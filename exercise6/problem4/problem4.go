package problem4

import (
	"sync"
	"time"
)

func worker(id int, cond *sync.Cond, shoppingList *[]string, notified *int, done *bool, wg *sync.WaitGroup) {
	defer wg.Done()

	cond.L.Lock()
	for !*done {
		cond.Wait()
	}
	if len(*shoppingList) > 0 && *notified == 0 {
		*notified = id
	}
	cond.L.Unlock()
}

func updateShopList(shoppingList *[]string, cond *sync.Cond, done *bool) {
	time.Sleep(10 * time.Millisecond)
	*shoppingList = append(*shoppingList, "apples", "milk", "baking soda")
	cond.L.Lock()
	*done = true
	cond.Broadcast()
	cond.L.Unlock()
}

func notifyOnShopListUpdate(shoppingList *[]string, numWorkers int) <-chan int {
	notifier := make(chan int, 1)
	var notified int
	done := false
	mutex := sync.Mutex{}
	cond := sync.NewCond(&mutex)
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i+1, cond, shoppingList, &notified, &done, &wg)
		time.Sleep(time.Millisecond)
	}

	go func() {
		updateShopList(shoppingList, cond, &done)
		wg.Wait()
		if notified != 0 {
			notifier <- notified
		}
		close(notifier)
	}()

	return notifier
}
