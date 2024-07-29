package main

import "strconv"

func binary(a int) string {
	if a == 0 {
		return "0"
	}
	res := ""

	for a != 0 {
		b := a % 2
		c := strconv.Itoa(b)
		res = c + res
		a = a / 2
	}
	return res
}
