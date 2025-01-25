package problem5

import "strings"

func producer(words []string, word chan<- string) {

	for _, val := range words {
		word <- val
	}
	defer close(word)
}

func consumer(word <-chan string) string {
	var res string
	for val := range word {

		if len(val) > 0 {
			res += " "
		}
		res += val
	}
	res = strings.Trim(res, " ")
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
