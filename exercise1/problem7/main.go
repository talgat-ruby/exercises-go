package main

func highestDigit(n int) int {
	highest := 0
	for n > 0 {
		digit := n % 10
		if digit > highest {
			highest = digit
		}
		n /= 10
	}
	return highest
}
