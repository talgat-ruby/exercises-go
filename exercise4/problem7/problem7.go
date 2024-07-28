package problem7

func multiplex(ch1 <-chan string, ch2 <-chan string) []string {
	var result []string
	done := make(chan bool)

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
				done <- true
				return
			}
		}
	}()

	<-done
	return result
}
