package problem5

import "strings"

func producer(words []string, ch chan<- string) {
	for _, i := range words {
		ch <- i
	}
	close(ch)
}

func consumer(ch <-chan string) string {
	words := ""

	for i := range ch {
		words += i + " "
	}
	return strings.TrimSpace(words)
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
