package main

import "regexp"

func potatoes(str string) int {
	regex := regexp.MustCompile(`potato`)
	res := regex.FindAllString(str, -1)
	return len(res)
}
