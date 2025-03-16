package problem5

import "strings"

func producer(str []string, ch chan<- string) {
	for _, s := range str {
		ch <- s
	}
	defer close(ch)
}

func consumer(ch <-chan string) string {
	var str []string

	for s := range ch {
		str = append(str, s)
	}

	return strings.Join(str, " ")
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
