package problem7

func multiplex(ch1 <-chan string, ch2 <-chan string) []string {
	var result []string
	for ch1 != nil || ch2 != nil {
		select {
		case msg, ok := <-ch1:
			if ok {
				result = append(result, msg)
			} else {
				ch1 = nil // Mark channel as closed
			}
		case msg, ok := <-ch2:
			if ok {
				result = append(result, msg)
			} else {
				ch2 = nil // Mark channel as closed
			}
		}
	}
	return result
}
