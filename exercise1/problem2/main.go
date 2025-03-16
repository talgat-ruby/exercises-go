package main

import "fmt"

func binary(decimal int) string {
	if decimal == 0 {
		return "0"
	}

	binaryStr := ""
	for decimal > 0 {
		remainder := decimal % 2
		binaryStr = fmt.Sprintf("%d", remainder) + binaryStr
		decimal = decimal / 2

	}
	return binaryStr
}
