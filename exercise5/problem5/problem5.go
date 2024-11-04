package problem5

func send(words []string, pr func([]string, chan<- string), cons func(<-chan string) string) string {
	ch := make(chan string)
	go pr(words, ch)
	return cons(ch)
}

func producer(words []string, ch chan<- string) {
	for _, word := range words {
		ch <- word
	}
	close(ch)
}

func consumer(ch <-chan string) string {
	var result string
	for word := range ch {
		if len(result) > 0 {
			result += " "
		}
		result += word
	}
	return result
}
