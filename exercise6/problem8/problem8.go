package problem8

func multiplex(chs []<-chan string) []string {
	var result []string
	done := make(chan struct{})
	resultChan := make(chan string)
	remaining := len(chs)

	for _, ch := range chs {
		go func(c <-chan string) {
			for val := range c {
				resultChan <- val
			}
			remaining--
			if remaining == 0 {
				close(resultChan)
			}
		}(ch)
	}

	go func() {
		for val := range resultChan {
			result = append(result, val)
		}
		close(done)
	}()

	<-done
	return result
}
