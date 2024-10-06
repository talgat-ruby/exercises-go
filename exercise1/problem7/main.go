package main

func highestDigit(num int) int {
	var max int = 0
	for num > 1 {
		cur := num % 10
		num = num / 10
		if cur > max {
			max = cur
		}
	}
	return max
}
