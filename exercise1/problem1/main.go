package main

import "log"

func addUp(n int) int {
	if n < 0 {
		log.Fatal()
	}
	if n == 0 || n == 1 {
		return 1
	}
	return n + addUp(n-1)
}
