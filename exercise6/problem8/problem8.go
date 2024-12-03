package problem8

import "sync"

func multiplex(chs []<-chan string) []string {
	result := []string{}
	wg := new(sync.WaitGroup)
	wg.Add(len(chs))
	mx := new(sync.Mutex)

	for _, ch := range chs {
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
