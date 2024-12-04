package problem7

func multiplex(ch1 <-chan string, ch2 <-chan string) []string {
	words := []string{}

	for {
		select {
		case word, ok := <-ch1:
			if ok {
				words = append(words, word)
			} else {
				ch1 = nil
			}
		case word, ok := <-ch2:
			if ok {
				words = append(words, word)
			} else {
				ch2 = nil
			}
		}

		if ch1 == nil && ch2 == nil {
			break
		}
	}

	return words
}
