package main

import "strconv"

func binary(decimal int) string {
	num, _ := strconv.Atoi(strconv.Itoa(decimal))
	binaryStr := strconv.FormatInt(int64(num), 2)
	return binaryStr
}
