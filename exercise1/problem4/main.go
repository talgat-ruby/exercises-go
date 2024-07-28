package main

import (
	"regexp"
)

func detectWord(word string) string {
	var res string

	wordRegexp := regexp.MustCompile("[a-z]")
	all := wordRegexp.FindAllString(word, -1)

	for _, val := range all {
		res += val
	}

	return res
}
