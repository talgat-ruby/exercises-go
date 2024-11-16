package problem7

func multiplex(ch1 <-chan string, ch2 <-chan string) []string {
	var result []string

	for ch1 != nil || ch2 != nil {
		select {
		case val, ok := <-ch1:
			if ok {
				result = append(result, val)
			} else {
				ch1 = nil // Канал ch1 закрыт
			}
		case val, ok := <-ch2:
			if ok {
				result = append(result, val)
			} else {
				ch2 = nil // Канал ch2 закрыт
			}
		}
	}

	return result
}
