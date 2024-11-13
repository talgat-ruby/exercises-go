package main

import (
	"fmt"
	"time"
)

// withTimeout waits for a message on the input channel until the specified ttl (time-to-live).
// If a message is received before the timeout, it forwards the message to the output channel.
// Otherwise, it sends "timeout" to the output channel after the ttl duration.
func withTimeout(input <-chan string, ttl time.Duration) <-chan string {
	output := make(chan string)

	go func() {
		defer close(output)
		select {
		case msg := <-input:
			// If a message is received from input before the timeout
			output <- msg
		case <-time.After(ttl):
			// If the timeout duration passes before receiving a message
			output <- "timeout"
		}
	}()

	return output
}

func main() {
	input := make(chan string)

	// Test with message received before timeout
	go func() {
		time.Sleep(1 * time.Second)
		input <- "hello"
	}()

	output := withTimeout(input, 2*time.Second)
	fmt.Println(<-output) // Should print "hello" since it arrives before the 2-second ttl

	// Test with message arriving after timeout
	go func() {
		time.Sleep(3 * time.Second)
		input <- "world"
	}()

	output = withTimeout(input, 2*time.Second)
	fmt.Println(<-output) // Should print "timeout" since 2-second ttl expires before "world" arrives
}
