package problem5

func producer(words []string, channel chan<- string) {
	for _, word := range words {
		channel <- word
	}
	close(channel)
}

func consumer(channel <-chan string) string {
	out := ""
	for word := range channel {
		out += word + " "
	}
	return out[:len(out)-1]
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
