package main

import (
	"slices"
	"strconv"
	"strings"
)

func highestDigit(num int) int {
	str := strconv.Itoa(num)

	max, _ := strconv.Atoi(slices.Max(strings.Split(str, "")))

	return max
}
