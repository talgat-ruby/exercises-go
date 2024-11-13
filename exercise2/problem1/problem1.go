package main

import "fmt"

// isChangeEnough checks if the given change array can cover the total amount due
func isChangeEnough(change []int, totalDue float64) bool {
	// Define the values for each coin type
	quarterValue := 0.25
	dimeValue := 0.10
	nickelValue := 0.05
	pennyValue := 0.01

	// Calculate the total amount from the change
	totalChange := float64(change[0])*quarterValue +
		float64(change[1])*dimeValue +
		float64(change[2])*nickelValue +
		float64(change[3])*pennyValue

	// Check if the total change is enough to cover the total due
	return totalChange >= totalDue
}

func main() {
	fmt.Println(isChangeEnough([]int{2, 100, 0, 0}, 14.11))   // Output: false
	fmt.Println(isChangeEnough([]int{0, 0, 20, 5}, 0.75))     // Output: true
	fmt.Println(isChangeEnough([]int{30, 40, 20, 5}, 12.55))  // Output: true
	fmt.Println(isChangeEnough([]int{10, 0, 0, 50}, 3.85))    // Output: false
	fmt.Println(isChangeEnough([]int{1, 0, 5, 219}, 19.99))   // Output: false
}
