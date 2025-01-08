package main

import (
	"fmt"
)

// multiplex takes multiple channels, collects values in the order they are received,
// and returns a slice of all values.
func multiplex(channels ...<-chan int) []int {
	var result []int
	done := make(chan struct{}) // To signal completion of all channels

	// Launch a goroutine for each channel to read values
	for _, ch := range channels {
		go func(c <-chan int) {
			for val := range c {
				result = append(result, val) // Collect values in the order they are received
			}
			done <- struct{}{} // Signal that this channel is done
		}(ch)
	}

	// Wait for all channels to finish
	for range channels {
		<-done
	}

	return result
}

func main() {
	// Example usage of multiplex with 3 channels
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

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

	result := multiplex(ch1, ch2, ch3)
	fmt.Println("Result:", result) // Example output could be [1, 2, 3, 4, 5], depending on order received
}
