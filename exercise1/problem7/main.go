package main

func highestDigit(a int) int {
	res := 0

	for a != 0 {
		b := a%10
		if b > res {
			res = b
		}
		a = a/10
	}
	return res
}
