package problem6

import (
	"sync"
)

func runTasks(init func()) {
	var wg sync.WaitGroup

	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()

			//TODO: modify so that load function gets called only once.
			init()
		}()
	}
	wg.Wait()
}
