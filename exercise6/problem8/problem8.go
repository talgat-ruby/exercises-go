package problem8

func multiplex(chs []<-chan string) []string {
	var result []string

	done := make(chan struct{})

	for _, ch := range chs {
		go func(ch <-chan string) {
			for val := range ch {
				result = append(result, val)
			}
			done <- struct{}{}
		}(ch)
	}

	for i := 0; i < len(chs); i++ {
		<-done
	}
	return result
}
