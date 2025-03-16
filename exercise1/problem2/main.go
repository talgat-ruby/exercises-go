package main

import "strconv"

func binary(n int) string {
	return strconv.FormatInt(int64(n), 2)
}
