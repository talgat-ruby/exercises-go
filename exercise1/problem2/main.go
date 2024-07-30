package main

import (
	"fmt"
	"strconv"
)

func binary(num int) string {
	way := 2
	var result string

	switch way {
	case 1:
		result = ConvertInt(num, 2)
	case 2:
		result = convert(num)
	}

	return result
}

func convert(num int) string {
	n := 1

	for n*2 <= num {
		n = n * 2
	}

	res := ""
	for n >= 1 {
		res = res + fmt.Sprintf("%d", num/n)
		num = num % n
		n = n / 2
	}

	return res
}

func ConvertInt(val int, toBase int) string {
	return strconv.FormatInt(int64(val), toBase)
}
