package problem6

import (
	"sync"
)

var once sync.Once

func runTasks(init func()) {
	var wg sync.WaitGroup

	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()

			once.Do(init)
		}()
	}
	wg.Wait()
}
