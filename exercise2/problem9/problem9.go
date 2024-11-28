package main

import "fmt"

// factory takes an integer multiplier and returns a function that multiplies each of its arguments by that multiplier
func factory(multiplier int) func(...int) []int {
	// Return a function that takes a variable number of ints and returns a slice of ints
	return func(nums ...int) []int {
		result := make([]int, len(nums))
		for i, num := range nums {
			result[i] = num * multiplier
		}
		return result
	}
}

func main() {
	first := factory(15)
	fmt.Println(first(2, 3, 4)) // Output: [30, 45, 60]

	second := factory(2)
	fmt.Println(second(1, 2, 3, 4)) // Output: [2, 4, 6, 8]
}
