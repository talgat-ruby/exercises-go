package main

import (
	"strconv"
	"strings"
)

func highestDigit(n int) int {
	res := 0
	text := strconv.Itoa(n)
	arr := strings.Split(text, "")

	for _, v := range arr {
		i, err := strconv.Atoi(v)

		if err != nil {
			break
		}
		if res < i {
			res = i
		}
	}

	return res
}
