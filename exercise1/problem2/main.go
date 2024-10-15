package main

import "strconv"

func binary(num int) string {
	newNum := int64(num)
	desiredBasedNumber := strconv.FormatInt(newNum, 2)
	return desiredBasedNumber
}
