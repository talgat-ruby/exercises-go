package main

import "fmt"

func main() {
	fmt.Println(highestDigit(379))
}

func highestDigit(n int) int {
	var digits []int
	var highest int
	for n > 0 {
		digit := n % 10
		n = n / 10
		digits = append(digits, digit)
	}
	fmt.Println(digits)

	for _, num := range digits {
		if num > highest {
			highest = num
		}
	}
	return highest
}
