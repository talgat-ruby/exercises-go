package problem8

func multiplex(chs []<-chan string) []string {
	resultChan := make(chan string)
	results := []string{}
	ok := make(chan struct{})

	for _, ch := range chs {
		go func(ch <-chan string) {
			for val := range ch {
				resultChan <- val
			}
			ok <- struct{}{}
		}(ch)
	}

	go func() {
		for range chs {
			<-ok
		}
		close(resultChan)
	}()

	for val := range resultChan {
		results = append(results, val)
	}

	return results
}
