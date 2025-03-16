package problem6

import (
	"sync"
)

var once sync.Once

func runTasks(init func()) {
	var wg sync.WaitGroup
	count := 0
	for count < 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			once.Do(init)
		}()
		count++
	}
	wg.Wait()
}
