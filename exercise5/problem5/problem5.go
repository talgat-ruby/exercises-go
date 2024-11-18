package problem5

func producer(words []string, ch chan<- string) {
	ch <- "first"
	for _, n := range words {
		ch <- n
	}
	ch <- "last"
	close(ch)
}

func consumer(ch <-chan string) string {
	var sum string
	var word string
	for n := range ch {
		if n != "last" && word != "first" && n != "first" {
			sum += word + " "
		}
		if n == "last" {
			sum += word
		}
		word = n
	}
	return sum
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
