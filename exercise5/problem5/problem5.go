package problem5

import "strings"

func producer(words []string, ch chan<- string) {
	for _, word := range words {
		ch <- word
	}
	close(ch)
}

func consumer(ch <-chan string) string {
	var result []string
	for word := range ch {
		result = append(result, word)
	}
	return strings.Join(result, " ")
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
