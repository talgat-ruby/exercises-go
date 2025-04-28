package problem5

func producer(list []string, ch chan<- string) {
	for i, l := range list {
		if i == len(list)-1 {
			ch <- l
			break
		}
		ch <- l + " "
	}
	close(ch)
}

func consumer(ch <-chan string) string {
	var result string
	for n := range ch {
		result += n
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
