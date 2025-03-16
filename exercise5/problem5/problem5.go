package problem5

func producer(words []string, ch chan<- string) {
	for _, word := range words {
		ch <- word
	}

	ch <- "done"
}

func consumer(ch <-chan string) string {
	var res string
	for n := range ch {
		if n == "done" {
			break
		}
		if res != "" {
			res += " "
		}
		res += n
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
