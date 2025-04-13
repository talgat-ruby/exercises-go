package problem7

func multiplex(ch1 <-chan string, ch2 <-chan string) []string {
	var result []string
	done := make(chan chan struct{})

	go func() {
		for {
			select {
			case val, ok := <-ch1:
				if ok {
					result = append(result, val)
				} else {
					ch1 = nil
				}
			case val, ok := <-ch2:
				if ok {
					result = append(result, val)
				} else {
					ch2 = nil
				}
			}
			if ch1 == nil && ch2 == nil {
				break
			}
		}
		close(done)
	}()
	<-done
	return result
}
