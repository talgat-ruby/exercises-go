package problem6

import (
	"sync"
)

func runTasks(init func()) {
	var wg sync.WaitGroup
	var once sync.Once // Ensures `init` is executed only once

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// Ensure `init` is called only once
			once.Do(init)
		}()
	}
	wg.Wait()
}
