package problem6

import (
	"sync"
)

func runTasks(init func()) {
	var wg sync.WaitGroup
	onceInit := sync.OnceFunc(init)

	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			onceInit()
		}()
	}
	wg.Wait()
}
