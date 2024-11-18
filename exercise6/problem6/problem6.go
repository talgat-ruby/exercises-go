package problem6

import (
	"sync"
)

func runTasks(init func()) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var isInit bool

	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()

			mu.Lock()
			if !isInit {
				init()
				isInit = true
			}
			mu.Unlock()
		}()
	}
	wg.Wait()
}
