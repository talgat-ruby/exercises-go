package main

import (
	"fmt"
	"math"
)

// Main function
func main() {

	// Finding largest number
	// among the given numbers
	// Using Max() function
	nvalue_1 := math.Max(34, 67)
	nvalue_2 := math.Max(56.7, 90.8)

	// Adding maximum numbers
	res := nvalue_1 + nvalue_2
	fmt.Printf("%.2f + %.2f = %.2f",
		nvalue_1, nvalue_2, res)

}
