package problem5

func producer(list []string, ch chan<- string) {
	for _, v := range list {
		ch <- v
	}
	close(ch)
}

func consumer(ch <-chan string) string {
	var text string
	for v := range ch {
		if text != "" {
			text += " "
		}
		text = text + v
	}
	return text
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
