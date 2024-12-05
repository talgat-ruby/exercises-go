package problem7

func multiplex(ch1 <-chan string, ch2 <-chan string) []string {
	ot := []string{}
	str := make(chan string)
	go func() {
		for {
			select {
			case val, ok := <-ch1:
				if ok {
					str <- val
				} else {
					ch1 = nil
				}
			case val, ok := <-ch2:
				if ok {
					str <- val
				} else {
					ch2 = nil
				}
			}

			if ch1 == nil && ch2 == nil {
				break
			}
		}
		close(str)
	}()

	for val := range str {
		ot = append(ot, val)
	}

	return ot
}
