package main

import (
	"fmt"
)

// PipeFunc represents a function that processes an integer and returns an integer
type PipeFunc func(int) int

// piper takes an initial value and a series of PipeFunc functions, passing the result of each pipe to the next
func piper(initial int, pipes ...PipeFunc) int {
	value := initial
	for _, pipe := range pipes {
		value = pipe(value)
	}
	return value
}

// Example pipes
func increment(x int) int {
	return x + 1
}

func double(x int) int {
	return x * 2
}

func square(x int) int {
	return x * x
}

func main() {
	// Example usage of piper with multiple pipes
	result := piper(2, increment, double, square) // ((2 + 1) * 2)^2 = 36
	fmt.Println("Result:", result)                // Output: 36
}
