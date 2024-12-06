package problem7

import "sync"

func multiplex(ch1 <-chan string, ch2 <-chan string) []string {
	var mu sync.Mutex
	var wg sync.WaitGroup

	res := []string{}
	ReadChannel := func(channel <-chan string) {
		defer wg.Done()
		for msq := range channel {
			mu.Lock()
			res = append(res, msq)
			mu.Unlock()
		}
	}
	wg.Add(2)
	go ReadChannel(ch1)
	go ReadChannel(ch2)
	wg.Wait()
	return res
}
