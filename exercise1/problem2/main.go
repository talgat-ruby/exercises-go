package main

import "fmt"

func binary(x int) string {
	var result string
	if x == 0 {
		return "0"
	}
	for x > 0 {
		result = fmt.Sprintf("%v", x%2) + result
		x = x / 2
	}
	return result
}
