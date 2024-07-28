package main

import (
	"strconv"
)

func binary(num int) string {
	if num == 0 {
		return "0"
	}

	var res string

	primes := []int{}
	sum := num

	for sum > 0 {
		primes = append(primes, sum%2)
		sum /= 2
	}

	for i := len(primes) - 1; i >= 0; i-- {
		res += strconv.Itoa(primes[i])
	}

	return res
}
