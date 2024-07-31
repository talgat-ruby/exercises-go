package main

func main() {
}

func highestDigit(num int) int {

	highest := 0

	for num > 0 {
		digit := num % 10
		if digit > highest {
			highest = digit
		}
		num = num / 10
	}

	return highest
}
