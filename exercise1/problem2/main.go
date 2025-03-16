package main

import "strconv"

func binary(num int) string {
	return strconv.FormatInt(int64(num), 2)
}
