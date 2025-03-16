package main

import (
	"strconv"
)

func binary(n int) string {
	if n == 0 {
		return "0"
	}
	return strconv.FormatInt(int64(n), 2)
}
