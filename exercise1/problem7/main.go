package main

func highestDigit(number int) int {
	maxDigit := 0
	for number != 0 {
		digit := number % 10
		if digit > maxDigit {
			maxDigit = digit
		}
		number = number / 10
	}
	return maxDigit
}
