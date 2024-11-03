package main

func highestDigit(num int) int {

	highest := 0

	for {
		digit := num % 10
		if digit > highest {
			highest = digit
		}

		num /= 10
		if num == 0 {
			break
		}
	}

	return highest
}
