package problem7

func multiplex(ch1 <-chan string, ch2 <-chan string) []string {
	var result []string

	for ch1 != nil || ch2 != nil {
		select {
		case msg, open := <-ch1:
			if open {
				result = append(result, msg)
			} else {
				ch1 = nil
			}
		case msg, open := <-ch2:
			if open {
				result = append(result, msg)
			} else {
				ch2 = nil
			}
		}
	}

	return result
}
