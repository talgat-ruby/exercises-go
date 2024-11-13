package main

import "fmt"

// removeDups removes duplicate items from a slice while maintaining order
func removeDups[T comparable](items []T) []T {
	seen := make(map[T]bool)
	result := []T{}

	for _, item := range items {
		// If item is not in the map, add it to the result and mark it as seen
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}

	return result
}

func main() {
	fmt.Println(removeDups([]int{1, 0, 1, 0}))             // Output: [1, 0]
	fmt.Println(removeDups([]bool{true, false, false, true})) // Output: [true, false]
	fmt.Println(removeDups([]string{"John", "Taylor", "John"})) // Output: ["John", "Taylor"]
}
