package problem7

import "sync"

func multiplex(channels ...<-chan string) []string {
	result := []string{}
	wg := new(sync.WaitGroup)
	wg.Add(len(channels))
	mx := new(sync.Mutex)

	for _, ch := range channels {
		go func(c <-chan string) {
			defer wg.Done()
			for v := range c {
				mx.Lock()
				result = append(result, v)
				mx.Unlock()
			}
		}(ch)
	}

	wg.Wait()

	return result
}
