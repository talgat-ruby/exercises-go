package problem5

import (
	"strings"
)

func send(words []string) string {
	ch := make(chan string)

	for _, word := range words {
		ch <- word
	}
	close(ch)

	var messages []string
	for word := range ch {
		messages = append(messages, word)
	}

	return strings.Join(messages, " ")
}
