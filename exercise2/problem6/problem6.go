package main

import "fmt"

// sumOfTwo checks if there exists a pair of numbers from slices `a` and `b` that adds up to `v`
func sumOfTwo(a, b []int, v int) bool {
	// Create a map to store the required complement values from array `a`
	complements := make(map[int]bool)

	// Fill the map with the complements for each element in `a`
	for _, num := range a {
		complements[v - num] = true
	}

	// Check if any element in `b` exists in the map of complements
	for _, num := range b {
		if complements[num] {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(sumOfTwo([]int{1, 2}, []int{4, 5, 6}, 5)) // Output: true
	fmt.Println(sumOfTwo([]int{1, 2}, []int{4, 5, 6}, 8)) // Output: true
	fmt.Println(sumOfTwo([]int{1, 2}, []int{4, 5, 6}, 3)) // Output: false
	fmt.Println(sumOfTwo([]int{1, 2}, []int{4, 5, 6}, 9)) // Output: false
}
