package main

import "strings"

func countVowels(text string) int {
	res := 0
	arr := strings.Split(text, "")

	for _, v := range arr {
		if v == "a" {
			res += 1
		}
		if v == "e" {
			res += 1
		}
		if v == "i" {
			res += 1
		}
		if v == "o" {
			res += 1
		}
		if v == "u" {
			res += 1
		}
	}

	return res
}
