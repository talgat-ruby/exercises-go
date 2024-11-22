package problem5

func producer(words []string, ch chan<- string) {
	for _, word := range words {
		ch <- word
	}
	close(ch)
}

func consumer(ch <-chan string) string {
	var res string
	for word := range ch {
		if res != "" {
			res += " " // Add a space between words
		}
		res += word
	}
	return res
}

func send(
	words []string,
	pr func([]string, chan<- string),
	cons func(<-chan string) string,
) string {
	ch := make(chan string)

	go pr(words, ch)

	return cons(ch)
}
