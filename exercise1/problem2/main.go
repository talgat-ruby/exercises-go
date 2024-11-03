package main

import "strconv"

func binary(num int) string {
	output := ""

	for {
		digit := strconv.Itoa(num % 2)
		num /= 2

		output = digit + output

		if num == 0 {
			break
		}
	}

	return output
}
