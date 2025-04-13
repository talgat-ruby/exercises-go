package main

func highestDigit(num int) int {

	res := 0

	for num > 0 {

		n := num % 10

		if n > res {

			res = n
		}
		num = num / 10
	}
	return res
}
