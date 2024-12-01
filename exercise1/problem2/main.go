package main

import (
	"strconv"
)

func Binary(num int) string {
	if num == 0 || num == 1 {
		return strconv.Itoa(num)
	}
	return Binary(num/2) + strconv.Itoa(num%2)
}
