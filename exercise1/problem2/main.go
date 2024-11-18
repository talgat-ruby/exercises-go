package main

import "strconv"

func binary(n int) string {
	//n = int64(n)
	return strconv.FormatInt(int64(n), 2)
}
