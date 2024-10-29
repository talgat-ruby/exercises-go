package problem5

func producer(words []string, ch chan<- string) {
	for _, i := range words {
		ch <- i
	}
	close(ch)
}

func consumer(ch <-chan string) string {
	var words string

	for i := range ch {
		words += i + " "
	}
	return words[:len(words)-1]
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
