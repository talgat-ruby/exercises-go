package problem5

func producer(words []string, ch chan<- string) {
	defer close(ch)
	if len(words) == 0 {
		return
	}

	for _, word := range words {
		ch <- word
	}

}

func consumer(ch <-chan string) string {
	var output string
	for w := range ch {
		output += w + " "
	}
	output = output[:len(output)-1]

	return output
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
