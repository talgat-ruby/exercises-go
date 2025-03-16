package main

import (
	"fmt"
	"strings"
)

// producer sends each word in the list to the channel
func producer(words []string, ch chan<- string) {
	for _, word := range words {
		ch <- word // send each word to the channel
	}
	close(ch) // close the channel when done
}

// consumer reads words from the channel and combines them into a single message
func consumer(ch <-chan string) string {
	var message []string
	for word := range ch {
		message = append(message, word) // collect words from the channel
	}
	return strings.Join(message, " ") // join the words with a space
}

// send combines the functionality of producer and consumer to create the final message
func send(words []string, producerFunc func([]string, chan<- string), consumerFunc func(<-chan string) string) string {
	ch := make(chan string)
	go producerFunc(words, ch)       // start the producer goroutine
	return consumerFunc(ch)          // consume from the channel and return the combined message
}

func main() {
	list := []string{
		"Hello",
		"dear",
		"friend!",
		"Learn",
		"from",
		"yesterday.",
		"Save",
		"our",
		"soles.",
	}

	message := send(list, producer, consumer)
	fmt.Println(message) // Output: "Hello dear friend! Learn from yesterday. Save our soles."
}
