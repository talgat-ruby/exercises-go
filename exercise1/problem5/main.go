package main

import "regexp"

func potatoes(word string) int {
	res := 0

	wordRegexp := regexp.MustCompile("potato")
	all := wordRegexp.FindAllString(word, -1)

	for range all {
		res += 1
	}

	return res
}
