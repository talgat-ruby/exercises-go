HEAD

 Eldar
6786b69556adaf3ccffc518c7655a26812d55e64
package main

import (
	"fmt"
)

// multiplexer receives a slice of channels, collects values in the order they are received,
// and returns a slice of all values.
func multiplexer(channels []<-chan int) []int {
	var result []int
	done := make(chan struct{}) // Signal channel to close after processing all values

	// Launch a goroutine for each channel to listen and add values to result
	for _, ch := range channels {
		go func(c <-chan int) {
			for val := range c {
				result = append(result, val) // Append values to result as they are received
			}
			done <- struct{}{} // Signal that this channel is done
		}(ch)
	}

	// Wait for all channels to signal completion
	for range channels {
		<-done
	}

	return result
}

func main() {
	// Create example channels
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	// Start sending values on each channel
	go func() {
		ch1 <- 1
		ch1 <- 2
		close(ch1)
	}()

	go func() {
		ch2 <- 3
		close(ch2)
	}()

	go func() {
		ch3 <- 4
		ch3 <- 5
		close(ch3)
	}()

	// Pass the channels to multiplexer
	result := multiplexer([]<-chan int{ch1, ch2, ch3})
	fmt.Println("Result:", result) // Output: [1 2 3 4 5] (order may vary based on arrival)
}
HEAD


package problem8

func multiplex(chs []<-chan string) []string {}
 main
6786b69556adaf3ccffc518c7655a26812d55e64
