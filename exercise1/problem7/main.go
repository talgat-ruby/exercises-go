package main

func highestDigit(n int) int {
	res := 0

	if n == 0 {
		return res
	}

	if n < 0 {
		n *= -1
	}

	a := n
	for a > 0 {
		temp := a % 10
		if temp > res {
			res = temp
		}
		a = a / 10
	}
	return res

}
