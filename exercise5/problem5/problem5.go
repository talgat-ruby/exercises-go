package problem5

import "fmt"

func producer(words []string, ch chan<- string) {
	defer close(ch)
	for _, word := range words {
		ch <- word
	}
}

func consumer(ch <-chan string) string {
	var result string
	for word := range ch {
		if result == "" {
			result += word
		} else {
			result += fmt.Sprintf(" %s", word)
		}
	}
	return result
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
