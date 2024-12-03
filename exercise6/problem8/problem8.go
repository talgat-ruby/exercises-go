package problem8

func multiplex(chs []<-chan string) []string {
	resultChan := make(chan string)
	results := []string{}
	done := make(chan struct{})

	for _, ch := range chs {
		go func(ch <-chan string) {
			for val := range ch {
				resultChan <- val
			}
			done <- struct{}{}
		}(ch)
	}

	go func() {
		for range chs {
			<-done
		}
		close(resultChan)
	}()

	for val := range resultChan {
		results = append(results, val)
	}

	return results
}
