package problem8

import "sync"

func multiplex(chs []<-chan string) []string {
	result := []string{}
	wg := new(sync.WaitGroup)
	resultCh := make(chan string)

	for _, ch := range chs {
		wg.Add(1)
		go func(c <-chan string) {
			defer wg.Done()
			for v := range c {
				resultCh <- v
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	for v := range resultCh {
		result = append(result, v)
	}

	return result
}
