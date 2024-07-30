package main

import "regexp"

func detectWord(str string) string {
	regex := regexp.MustCompile(`[A-Z]`)
	return regex.ReplaceAllString(str, "")
}
