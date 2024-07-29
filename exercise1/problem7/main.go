package main

func highestDigit(num int) int {
	maxDigit := 0

	// Обрабатываем отрицательные числа
	if num < 0 {
		num = -num
	}

	for num > 0 {
		digit := num % 10
		if digit > maxDigit {
			maxDigit = digit
		}
		num /= 10
	}

	return maxDigit
}
