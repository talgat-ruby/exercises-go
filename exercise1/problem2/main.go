package main

import "strconv"

func binary(num int) string {
	res := ""
	if num == 0 {
		return "0"
	}
	for num > 0 {
		res += strconv.Itoa(num % 2)
		num /= 2
	}
	temp_res := ""
	for i := len(res) - 1; i >= 0; i-- {
		temp_res += string(res[i])
	}
	return temp_res
}
