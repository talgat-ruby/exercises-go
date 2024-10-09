package main

func highestDigit(num int) int {
	maxNum := 0
	for num > 0 {
		digit := num % 10
		maxNum = max(maxNum, digit)
		num /= 10
	}
	return maxNum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
