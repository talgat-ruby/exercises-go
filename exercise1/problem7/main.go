package main

func highestDigit(num int) int {
	res := 0
	for num > 0 {
		temp := num % 10
		if temp > res {
			res = temp
		}
		if res == 9 {
			return 9
		}
		num /= 10
	}
	return res
}
