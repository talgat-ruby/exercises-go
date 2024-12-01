package problem5

func producer(words []string, ch chan<- string) {
	for i, w := range words {
		if i != len(words)-1 {
			ch <- w + " "
			continue
		}
		ch <- w
	}
	close(ch)
}

func consumer(ch <-chan string) string {
	msg := ""
	for {
		w, more := <-ch
		if more {
			msg += w
		} else {
			return msg
		}
	}
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
