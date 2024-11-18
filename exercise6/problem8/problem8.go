package problem8

import "sync"

func multiplex(chs []<-chan string) []string {
	var wg sync.WaitGroup
	results := make([]string, 0)
	resultCh := make(chan string)

	for _, ch := range chs {
		wg.Add(1)
		go func(ch <-chan string) {
			defer wg.Done()
			for msg := range ch {
				resultCh <- msg
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	for msg := range resultCh {
		results = append(results, msg)
	}

	return results
}
