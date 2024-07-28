package main

import "strings"

func countVowels(s string) int {
	sum := 0
	s = strings.ToLower(s)
	for _, j := range s {
		switch j {
		case 'a', 'e', 'i', 'o', 'u':
			sum++
		}
	}
	return sum
}
