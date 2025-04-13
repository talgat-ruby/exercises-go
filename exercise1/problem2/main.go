package main

import "fmt"

func binary(num int) string {
	if num == 0 {
		return "0"
	}
	var binaryStr string
	for num > 0 {
		remainder := num % 2

		binaryStr = fmt.Sprintf("%d", remainder) + binaryStr

		num /= 2
	}

	return binaryStr
}
