package main

import "strconv"

func binary(num int) string {
	res := strconv.FormatInt(int64(num), 2)
	return res
}
