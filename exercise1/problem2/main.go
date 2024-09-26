package main

import "strconv"

func binary(n int) string {
	if n == 0 {
		return "0"
	}

	binaryStr := ""
	for n > 0 {
		remainder := n % 2
		binaryStr = strconv.Itoa(remainder) + binaryStr
		n = n / 2
	}

	return binaryStr
}
