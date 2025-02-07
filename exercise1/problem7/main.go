package main

func highestDigit(a int) int {
	max := 0
	m := 0
	for a > 0 {
		m = a % 10
		if m > max {
			max = m
		}
		a /= 10
	}
	return max
}
