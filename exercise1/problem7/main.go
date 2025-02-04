package main

func highestDigit(num int) int {
	highest := 0
	for num > 0 {
		rem := num % 10
		if rem > highest {
			highest = rem
		}
		num = num / 10
	}
	return highest
}
