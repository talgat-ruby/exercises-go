package main

import "strconv"

func binary(x int) string {
	if x == 0 {
		 return "0"
	}
	var ans string
	for x > 0 {
		 ans =strconv.Itoa(x%2) + ans
		 x /= 2
	}
	return ans
}
