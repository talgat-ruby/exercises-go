package main

import (
	"fmt"
)

func sum(nums []int) int {
	numParts := 4 // Split into parts, adjust based on desired concurrency
	partSize := (len(nums) + numParts - 1) / numParts

	resultChan := make(chan int, numParts)

	for i := 0; i < numParts; i++ {
		start := i * partSize
		end := start + partSize
		if end > len(nums) {
			end = len(nums)
		}
		// Launch a goroutine for each part to compute its partial sum
		go func(numsPart []int) {
			sum := 0
			for _, num := range numsPart {
				sum += num
			}
			resultChan <- sum // Send partial sum to channel
		}(nums[start:end])
	}

	// Collect results from all goroutines
	totalSum := 0
	for i := 0; i < numParts; i++ {
		totalSum += <-resultChan
	}

	return totalSum
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Total sum:", sum(nums)) // Should print the total sum of the nums slice
}
