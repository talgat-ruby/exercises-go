package problem8

import "sync"

func multiplex(chs []<-chan string) []string {
	var wg sync.WaitGroup
	results := make([]string, 0)

	// Create a channel to collect results
	resultCh := make(chan string)

	// Start a goroutine for each input channel
	for _, ch := range chs {
		wg.Add(1)
		go func(ch <-chan string) {
			defer wg.Done()
			for msg := range ch {
				resultCh <- msg
			}
		}(ch)
	}

	// Close the result channel once all workers finish
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	// Collect results from the result channel
	for msg := range resultCh {
		results = append(results, msg)
	}

	return results
}
