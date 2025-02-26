package main

import "strconv"

func binary(n int) (binaryStr string) {
	if n == 0 {
		return "0"
	}

	for n > 0 {
		binaryStr = strconv.Itoa(n%2) + binaryStr
		n /= 2
	}

	return
}
