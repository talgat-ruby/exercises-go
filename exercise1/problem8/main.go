package main

import "regexp"

func countVowels(str string) int {
	regex := regexp.MustCompile(`[aeiouAEIOU]`)
	res := regex.FindAllString(str, -1)
	return len(res)
}
