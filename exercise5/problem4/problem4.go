package main

import "fmt"

// sum calculates the sum of all elements in the nums slice
func sum(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Total sum:", sum(nums)) // Should print the total sum of the nums slice
}
