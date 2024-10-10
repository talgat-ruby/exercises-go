package problem11

import "fmt"

func removeDups[T comparable](items []T) []T {
	seen := make(map[T]bool)
	var result []T

	for _, item := range items {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}

	return result
}

func main() {
	fmt.Println(removeDups([]int{1, 0, 1, 0}))
	fmt.Println(removeDups([]bool{true, false, false, true}))
	fmt.Println(removeDups([]string{"John", "Taylor", "John"}))
}
