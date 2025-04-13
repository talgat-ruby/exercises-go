package problem5

import "strings"

func producer(words []string, ch chan<- string) {
	for _, word := range words {
		ch <- word
	}
	close(ch)
}

func consumer(ch <-chan string) string {
	var message []string
	for word := range ch {
		message = append(message, word)
	}
	return strings.Join(message, " ")
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
