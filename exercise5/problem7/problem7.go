package problem7

import "sync"

func multiplex(ch1 <-chan string, ch2 <-chan string) []string {
	var result []string
	var mu sync.Mutex
	var wg sync.WaitGroup

	read := func(ch <-chan string) {
		defer wg.Done()
		for val := range ch {
			mu.Lock()
			result = append(result, val)
			mu.Unlock()
		}

	}
	wg.Add(2)
	go read(ch1)
	go read(ch2)
	wg.Wait()
	return result
}
